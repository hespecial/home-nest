package logic

import (
	"context"
	"github.com/pkg/errors"
	"home-nest/pkg/xerr"

	"github.com/jinzhu/copier"
	"home-nest/app/travel/cmd/rpc/internal/svc"
	"home-nest/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessListLogic {
	return &HomestayBusinessListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomestayBusinessListLogic) HomestayBusinessList(in *pb.HomestayBusinessListReq) (*pb.HomestayBusinessListResp, error) {
	whereBuilder := l.svcCtx.HomestayBusinessModel.SelectBuilder()
	list, err := l.svcCtx.HomestayBusinessModel.FindPageListByIdDESC(l.ctx, whereBuilder, in.LastId, in.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "HomestayBussinessList FindPageListByIdDESC db fail , err:%v", err)
	}

	var resp []*pb.HomestayBusiness
	if len(list) > 0 {
		for _, item := range list {
			var typeHomestayBusinessListInfo pb.HomestayBusiness
			_ = copier.Copy(&typeHomestayBusinessListInfo, item)

			resp = append(resp, &typeHomestayBusinessListInfo)
		}
	}

	return &pb.HomestayBusinessListResp{
		List: resp,
	}, nil
}
