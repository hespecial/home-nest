package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"home-nest/app/travel/cmd/rpc/internal/svc"
	"home-nest/app/travel/cmd/rpc/pb"
	"home-nest/app/travel/model"
	"home-nest/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayListLogic {
	return &HomestayListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomestayListLogic) HomestayList(in *pb.HomestayListReq) (*pb.HomestayListResp, error) {
	whereBuilder := l.svcCtx.HomestayActivityModel.SelectBuilder().Where(squirrel.Eq{
		"row_type":   model.HomestayActivityPreferredType,
		"row_status": model.HomestayActivityUpStatus,
	})
	homestayActivityList, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, whereBuilder, in.Page, in.PageSize, "data_id desc")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "get activity homestay id set fail rowType: %s ,err : %v", model.HomestayActivityPreferredType, err)
	}

	list, err := l.getHomeStayListByActivityList(homestayActivityList)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "homestay list mapreduce process error : %v", err)
	}

	return &pb.HomestayListResp{
		List: list,
	}, nil
}

func (l *HomestayListLogic) getHomeStayListByActivityList(homestayActivityList []*model.HomestayActivity) ([]*pb.Homestay, error) {
	list, err := mr.MapReduce[int64, *model.Homestay, []*pb.Homestay](
		func(source chan<- int64) {
			for _, homestayActivity := range homestayActivityList {
				source <- homestayActivity.DataId
			}
		},
		func(id int64, writer mr.Writer[*model.Homestay], cancel func(error)) {
			homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, id)
			if err != nil && !errors.Is(err, model.ErrNotFound) {
				logx.WithContext(l.ctx).Errorf("ActivityHomestayListLogic ActivityHomestayList 获取活动数据失败 id : %d ,err : %v", id, err)
				return
			}
			writer.Write(homestay)
		},
		func(pipe <-chan *model.Homestay, writer mr.Writer[[]*pb.Homestay], cancel func(error)) {
			var list []*pb.Homestay
			for homestay := range pipe {
				list = append(list, &pb.Homestay{
					Id:                  homestay.Id,
					Title:               homestay.Title,
					SubTitle:            homestay.SubTitle,
					Banner:              homestay.Banner,
					Info:                homestay.Info,
					PeopleNum:           homestay.PeopleNum,
					HomestayBusinessId:  homestay.HomestayBusinessId,
					UserId:              homestay.UserId,
					RowState:            homestay.RowState,
					RowType:             homestay.RowType,
					FoodInfo:            homestay.FoodInfo,
					FoodPrice:           homestay.FoodPrice,
					HomestayPrice:       homestay.HomestayPrice,
					MarketHomestayPrice: homestay.MarketHomestayPrice,
				})
			}
			writer.Write(list)
		},
	)
	return list, err
}
