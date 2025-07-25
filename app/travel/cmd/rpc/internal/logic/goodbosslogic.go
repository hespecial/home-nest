package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"home-nest/app/travel/model"
	"home-nest/app/usercenter/cmd/rpc/usercenter"
	"home-nest/pkg/xerr"

	"home-nest/app/travel/cmd/rpc/internal/svc"
	"home-nest/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodBossLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodBossLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodBossLogic {
	return &GoodBossLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoodBossLogic) GoodBoss(_ *pb.GoodBossReq) (*pb.GoodBossResp, error) {
	whereBuilder := l.svcCtx.HomestayActivityModel.SelectBuilder().Where(squirrel.Eq{
		"row_type":   model.HomestayActivityGoodBusiType,
		"row_status": model.HomestayActivityUpStatus,
	})
	// 前十房东
	homestayActivityList, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, whereBuilder, 0, 10, "data_id desc")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "get GoodBoss db err. rowType: %s ,err : %v", model.HomestayActivityGoodBusiType, err)
	}

	list, err := l.getHomeStayListByActivityList(homestayActivityList)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "good booss mapreduce process error : %v", err)
	}

	return &pb.GoodBossResp{
		List: list,
	}, nil
}

func (l *GoodBossLogic) getHomeStayListByActivityList(homestayActivityList []*model.HomestayActivity) ([]*pb.HomestayBusinessBoss, error) {
	// todo 后续补充：去重、排序

	list, err := mr.MapReduce[int64, *model.Homestay, []*pb.HomestayBusinessBoss](
		func(source chan<- int64) {
			for _, homestayActivity := range homestayActivityList {
				source <- homestayActivity.DataId
			}
		},
		func(id int64, writer mr.Writer[*model.Homestay], cancel func(error)) {
			homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, id)
			if err != nil {
				logx.WithContext(l.ctx).Errorf("GoodListLogic GoodList fail homestayId : %d ,err:%v", id, err)
				return
			}
			if homestay != nil {
				writer.Write(homestay)
			}
		},
		func(pipe <-chan *model.Homestay, writer mr.Writer[[]*pb.HomestayBusinessBoss], cancel func(error)) {
			var list []*pb.HomestayBusinessBoss
			for homestay := range pipe {
				userResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
					Id: homestay.UserId,
				})
				if err != nil {
					logx.WithContext(l.ctx).Errorf("GoodListLogic GoodList fail userId : %d ,err:%v", homestay.UserId, err)
					return
				}
				user := userResp.User

				list = append(list, &pb.HomestayBusinessBoss{
					//Id:       ,
					UserId:   user.Id,
					Nickname: user.Nickname,
					Avatar:   user.Avatar,
					Info:     user.Info,
					//Rank:     ,
				})
			}
			writer.Write(list)
		},
	)
	return list, err
}
