package result

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"home-nest/pkg/xerr"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := xerr.ServerInternalError
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err) // err类型
		var e *xerr.CodeError
		if errors.As(causeErr, &e) { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		if errcode != xerr.ServerCommonError {
			logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		}
		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
	}
}

// http 参数错误返回
func ParamErrorResult(_ *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.RequestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.RequestParamError, errMsg))
}
