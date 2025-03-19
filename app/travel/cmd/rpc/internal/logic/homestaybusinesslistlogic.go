package logic

import (
	"context"

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

func (l *HomestayBusinessListLogic) HomestayBusinessList(in *pb.HomestayListReq) (*pb.HomestayBusinessListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.HomestayBusinessListResp{}, nil
}
