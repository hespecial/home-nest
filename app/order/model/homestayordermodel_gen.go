// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.6

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	homestayOrderFieldNames          = builder.RawFieldNames(&HomestayOrder{})
	homestayOrderRows                = strings.Join(homestayOrderFieldNames, ",")
	homestayOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(homestayOrderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	homestayOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(homestayOrderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheHomeNestOrderHomestayOrderIdPrefix = "cache:homeNestOrder:homestayOrder:id:"
	cacheHomeNestOrderHomestayOrderSnPrefix = "cache:homeNestOrder:homestayOrder:sn:"
)

type (
	homestayOrderModel interface {
		Insert(ctx context.Context, data *HomestayOrder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*HomestayOrder, error)
		FindOneBySn(ctx context.Context, sn string) (*HomestayOrder, error)
		Update(ctx context.Context, data *HomestayOrder) error
		Delete(ctx context.Context, id int64) error
	}

	defaultHomestayOrderModel struct {
		sqlc.CachedConn
		table string
	}

	HomestayOrder struct {
		Id                  int64     `db:"id"`
		CreateTime          time.Time `db:"create_time"`
		UpdateTime          time.Time `db:"update_time"`
		DeleteTime          time.Time `db:"delete_time"`
		DelState            int64     `db:"del_state"`
		Version             int64     `db:"version"`               // 版本号
		Sn                  string    `db:"sn"`                    // 订单号
		UserId              int64     `db:"user_id"`               // 下单用户id
		HomestayId          int64     `db:"homestay_id"`           // 民宿id
		Title               string    `db:"title"`                 // 标题
		SubTitle            string    `db:"sub_title"`             // 副标题
		Cover               string    `db:"cover"`                 // 封面
		Info                string    `db:"info"`                  // 介绍
		PeopleNum           int64     `db:"people_num"`            // 容纳人的数量
		RowType             int64     `db:"row_type"`              // 售卖类型0：按房间出售 1:按人次出售
		NeedFood            int64     `db:"need_food"`             // 0:不需要餐食 1:需要参数
		FoodInfo            string    `db:"food_info"`             // 餐食标准
		FoodPrice           int64     `db:"food_price"`            // 餐食价格(分)
		HomestayPrice       int64     `db:"homestay_price"`        // 民宿价格(分)
		MarketHomestayPrice int64     `db:"market_homestay_price"` // 民宿市场价格(分)
		HomestayBusinessId  int64     `db:"homestay_business_id"`  // 店铺id
		HomestayUserId      int64     `db:"homestay_user_id"`      // 店铺房东id
		LiveStartDate       time.Time `db:"live_start_date"`       // 开始入住日期
		LiveEndDate         time.Time `db:"live_end_date"`         // 结束入住日期
		LivePeopleNum       int64     `db:"live_people_num"`       // 实际入住人数
		TradeState          int64     `db:"trade_state"`           // -1: 已取消 0:待支付 1:未使用 2:已使用  3:已退款 4:已过期
		TradeCode           string    `db:"trade_code"`            // 确认码
		Remark              string    `db:"remark"`                // 用户下单备注
		OrderTotalPrice     int64     `db:"order_total_price"`     // 订单总价格（餐食总价格+民宿总价格）(分)
		FoodTotalPrice      int64     `db:"food_total_price"`      // 餐食总价格(分)
		HomestayTotalPrice  int64     `db:"homestay_total_price"`  // 民宿总价格(分)
	}
)

func newHomestayOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultHomestayOrderModel {
	return &defaultHomestayOrderModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`homestay_order`",
	}
}

func (m *defaultHomestayOrderModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	homeNestOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderIdPrefix, id)
	homeNestOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderSnPrefix, data.Sn)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, homeNestOrderHomestayOrderIdKey, homeNestOrderHomestayOrderSnKey)
	return err
}

func (m *defaultHomestayOrderModel) FindOne(ctx context.Context, id int64) (*HomestayOrder, error) {
	homeNestOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderIdPrefix, id)
	var resp HomestayOrder
	err := m.QueryRowCtx(ctx, &resp, homeNestOrderHomestayOrderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", homestayOrderRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayOrderModel) FindOneBySn(ctx context.Context, sn string) (*HomestayOrder, error) {
	homeNestOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderSnPrefix, sn)
	var resp HomestayOrder
	err := m.QueryRowIndexCtx(ctx, &resp, homeNestOrderHomestayOrderSnKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `sn` = ? limit 1", homestayOrderRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, sn); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHomestayOrderModel) Insert(ctx context.Context, data *HomestayOrder) (sql.Result, error) {
	homeNestOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderIdPrefix, data.Id)
	homeNestOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderSnPrefix, data.Sn)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, homestayOrderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Sn, data.UserId, data.HomestayId, data.Title, data.SubTitle, data.Cover, data.Info, data.PeopleNum, data.RowType, data.NeedFood, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.HomestayBusinessId, data.HomestayUserId, data.LiveStartDate, data.LiveEndDate, data.LivePeopleNum, data.TradeState, data.TradeCode, data.Remark, data.OrderTotalPrice, data.FoodTotalPrice, data.HomestayTotalPrice)
	}, homeNestOrderHomestayOrderIdKey, homeNestOrderHomestayOrderSnKey)
	return ret, err
}

func (m *defaultHomestayOrderModel) Update(ctx context.Context, newData *HomestayOrder) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	homeNestOrderHomestayOrderIdKey := fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderIdPrefix, data.Id)
	homeNestOrderHomestayOrderSnKey := fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderSnPrefix, data.Sn)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayOrderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.Sn, newData.UserId, newData.HomestayId, newData.Title, newData.SubTitle, newData.Cover, newData.Info, newData.PeopleNum, newData.RowType, newData.NeedFood, newData.FoodInfo, newData.FoodPrice, newData.HomestayPrice, newData.MarketHomestayPrice, newData.HomestayBusinessId, newData.HomestayUserId, newData.LiveStartDate, newData.LiveEndDate, newData.LivePeopleNum, newData.TradeState, newData.TradeCode, newData.Remark, newData.OrderTotalPrice, newData.FoodTotalPrice, newData.HomestayTotalPrice, newData.Id)
	}, homeNestOrderHomestayOrderIdKey, homeNestOrderHomestayOrderSnKey)
	return err
}

func (m *defaultHomestayOrderModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheHomeNestOrderHomestayOrderIdPrefix, primary)
}

func (m *defaultHomestayOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", homestayOrderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultHomestayOrderModel) tableName() string {
	return m.table
}
