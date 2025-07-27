package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"home-nest/app/mqueue/cmd/job/jobtype"
	"home-nest/app/order/model"
	"home-nest/app/travel/cmd/rpc/travel"
	"home-nest/pkg/globalkey"
	"home-nest/pkg/tool"
	"home-nest/pkg/uniqueid"
	"home-nest/pkg/xerr"
	"strings"
	"time"

	"home-nest/app/order/cmd/rpc/internal/svc"
	"home-nest/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const CloseOrderTimeMinutes = 5 //defer close order time

type CreateHomestayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateHomestayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHomestayOrderLogic {
	return &CreateHomestayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateHomestayOrderLogic) CreateHomestayOrder(in *pb.CreateHomestayOrderReq) (*pb.CreateHomestayOrderResp, error) {
	//1、Create Order
	if in.LiveEndTime <= in.LiveStartTime {
		return nil, errors.Wrapf(xerr.NewErrMsg("Stay at least one night"), "Place an order at a B&B. The end time of your stay must be greater than the start time. in : %+v", in)
	}

	resp, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &travel.HomestayDetailReq{
		Id: in.HomestayId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to query the record"), "Failed to query the record  rpc HomestayDetail fail , homestayId : %d , err : %v", in.HomestayId, err)
	}
	if resp.Homestay == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("This record does not exist"), "This record does not exist , homestayId : %d ", in.HomestayId)
	}

	var cover string //Get the cover...
	if len(resp.Homestay.Banner) > 0 {
		cover = strings.Split(resp.Homestay.Banner, ",")[0]
	}

	homestay := resp.Homestay
	order := &model.HomestayOrder{
		Sn:                  uniqueid.GenSn(uniqueid.SnPrefixHomestayOrder),
		UserId:              in.UserId,
		HomestayId:          homestay.Id,
		Title:               homestay.Title,
		SubTitle:            homestay.SubTitle,
		Cover:               cover,
		Info:                homestay.Info,
		PeopleNum:           homestay.PeopleNum,
		RowType:             homestay.RowType,
		FoodInfo:            homestay.FoodInfo,
		FoodPrice:           homestay.FoodPrice,
		HomestayPrice:       homestay.HomestayPrice,
		MarketHomestayPrice: homestay.MarketHomestayPrice,
		HomestayBusinessId:  homestay.HomestayBusinessId,
		HomestayUserId:      homestay.UserId,
		LiveStartDate:       time.Unix(in.LiveStartTime, 0),
		LiveEndDate:         time.Unix(in.LiveEndTime, 0),
		LivePeopleNum:       in.LivePeopleNum,
		TradeState:          model.HomestayOrderTradeStateWaitPay,
		TradeCode:           tool.Krand(8, tool.KcRandKindAll),
		Remark:              in.Remark,
		DeleteTime:          time.Unix(0, 0),
		DelState:            globalkey.DelStateNo,
	}

	liveDays := int64(order.LiveEndDate.Sub(order.LiveStartDate).Seconds() / 86400)

	order.HomestayTotalPrice = homestay.HomestayPrice * liveDays
	if in.IsFood {
		order.NeedFood = model.HomestayOrderNeedFoodYes
		//Calculate the total price of the meal.
		order.FoodTotalPrice = homestay.FoodPrice * in.LivePeopleNum * liveDays
	}

	order.OrderTotalPrice = order.HomestayTotalPrice + order.FoodTotalPrice //Calculate total order price.

	_, err = l.svcCtx.HomestayOrderModel.Insert(l.ctx, order)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Order Database Exception order : %+v , err: %v", order, err)
	}

	//2、Delayed closing of order tasks.
	payload, err := json.Marshal(jobtype.DeferCloseHomestayOrderPayload{Sn: order.Sn})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create defer close order task json Marshal fail err :%+v , sn : %s", err, order.Sn)
	} else {
		_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseHomestayOrder, payload), asynq.ProcessIn(CloseOrderTimeMinutes*time.Minute))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("create defer close order task insert queue fail err :%+v , sn : %s", err, order.Sn)
		}
	}

	return &pb.CreateHomestayOrderResp{
		Sn: order.Sn,
	}, nil
}
