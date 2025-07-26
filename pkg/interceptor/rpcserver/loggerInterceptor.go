package rpcserver

import (
	"context"
	"errors"
	gerr "github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"home-nest/pkg/xerr"
)

func LoggerInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := gerr.Cause(err) // err类型
		var e *xerr.CodeError
		if errors.As(causeErr, &e) { //自定义错误类型
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)

			//转成grpc err
			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
		}

	}

	return resp, err
}
