package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"home-nest/app/order/model"
	"home-nest/pkg/xerr"

	"home-nest/app/order/cmd/rpc/internal/svc"
	"home-nest/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户民宿订单
func (l *UserHomestayOrderListLogic) UserHomestayOrderList(in *pb.UserHomestayOrderListReq) (*pb.UserHomestayOrderListResp, error) {
	whereBuilder := l.svcCtx.HomestayOrderModel.SelectBuilder().Where(squirrel.Eq{"user_id": in.UserId})
	//There are supported states in the filter, otherwise return all
	if in.TraderState >= model.HomestayOrderTradeStateCancel && in.TraderState <= model.HomestayOrderTradeStateExpire {
		whereBuilder = whereBuilder.Where(squirrel.Eq{"trade_state": in.TraderState})
	}

	list, err := l.svcCtx.HomestayOrderModel.FindPageListByIdDESC(l.ctx, whereBuilder, in.LastId, in.PageSize)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Failed to get user's homestay order err : %v , in :%+v", err, in)
	}

	var resp []*pb.HomestayOrder
	if len(list) > 0 {
		for _, homestayOrder := range list {
			var pbHomestayOrder pb.HomestayOrder
			_ = copier.Copy(&pbHomestayOrder, homestayOrder)

			pbHomestayOrder.CreateTime = homestayOrder.CreateTime.Unix()
			pbHomestayOrder.LiveStartDate = homestayOrder.LiveStartDate.Unix()
			pbHomestayOrder.LiveEndDate = homestayOrder.LiveEndDate.Unix()

			resp = append(resp, &pbHomestayOrder)
		}
	}

	return &pb.UserHomestayOrderListResp{
		List: resp,
	}, nil
}
