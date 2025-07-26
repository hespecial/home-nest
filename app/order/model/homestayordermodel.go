package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"home-nest/pkg/globalkey"
)

var _ HomestayOrderModel = (*customHomestayOrderModel)(nil)

type (
	// HomestayOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayOrderModel.
	HomestayOrderModel interface {
		homestayOrderModel
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *HomestayOrder) error
		SelectBuilder() squirrel.SelectBuilder
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayOrder, error)
	}

	customHomestayOrderModel struct {
		*defaultHomestayOrderModel
	}
)

// NewHomestayOrderModel returns a model for the database table.
func NewHomestayOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HomestayOrderModel {
	return &customHomestayOrderModel{
		defaultHomestayOrderModel: newHomestayOrderModel(conn, c, opts...),
	}
}

func (m *defaultHomestayOrderModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultHomestayOrderModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, newData *HomestayOrder) error {

	oldVersion := newData.Version
	newData.Version += 1

	var sqlResult sql.Result
	var err error

	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	homenestOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderIdPrefix, data.Id)
	homenestOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderSnPrefix, data.Sn)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, homestayOrderRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.Sn, newData.UserId, newData.HomestayId, newData.Title, newData.SubTitle, newData.Cover, newData.Info, newData.PeopleNum, newData.RowType, newData.NeedFood, newData.FoodInfo, newData.FoodPrice, newData.HomestayPrice, newData.MarketHomestayPrice, newData.HomestayBusinessId, newData.HomestayUserId, newData.LiveStartDate, newData.LiveEndDate, newData.LivePeopleNum, newData.TradeState, newData.TradeCode, newData.Remark, newData.OrderTotalPrice, newData.FoodTotalPrice, newData.HomestayTotalPrice, newData.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.Sn, newData.UserId, newData.HomestayId, newData.Title, newData.SubTitle, newData.Cover, newData.Info, newData.PeopleNum, newData.RowType, newData.NeedFood, newData.FoodInfo, newData.FoodPrice, newData.HomestayPrice, newData.MarketHomestayPrice, newData.HomestayBusinessId, newData.HomestayUserId, newData.LiveStartDate, newData.LiveEndDate, newData.LivePeopleNum, newData.TradeState, newData.TradeCode, newData.Remark, newData.OrderTotalPrice, newData.FoodTotalPrice, newData.HomestayTotalPrice, newData.Id, oldVersion)
	}, homenestOrderHomestayOrderIdKey, homenestOrderHomestayOrderSnKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return ErrNoRowsUpdate
	}

	return nil
}

func (m *defaultHomestayOrderModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayOrder, error) {

	builder = builder.Columns(homestayOrderRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayOrder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
