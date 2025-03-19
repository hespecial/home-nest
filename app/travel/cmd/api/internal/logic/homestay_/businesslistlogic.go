package homestay_

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"home-nest/app/travel/cmd/api/internal/svc"
	"home-nest/app/travel/cmd/api/internal/types"
)

type BusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessListLogic {
	return &BusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BusinessListLogic) BusinessList(req *types.BusinessListReq) (resp *types.BusinessListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
