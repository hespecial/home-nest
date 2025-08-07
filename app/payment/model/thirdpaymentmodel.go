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

var _ ThirdPaymentModel = (*customThirdPaymentModel)(nil)

type (
	// ThirdPaymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customThirdPaymentModel.
	ThirdPaymentModel interface {
		thirdPaymentModel
		SelectBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*ThirdPayment, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, newData *ThirdPayment) error
	}

	customThirdPaymentModel struct {
		*defaultThirdPaymentModel
	}
)

// NewThirdPaymentModel returns a model for the database table.
func NewThirdPaymentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ThirdPaymentModel {
	return &customThirdPaymentModel{
		defaultThirdPaymentModel: newThirdPaymentModel(conn, c, opts...),
	}
}

func (m *customThirdPaymentModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *customThirdPaymentModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*ThirdPayment, error) {

	builder = builder.Columns(thirdPaymentRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ThirdPayment
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customThirdPaymentModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, newData *ThirdPayment) error {

	oldVersion := newData.Version
	newData.Version += 1

	var sqlResult sql.Result
	var err error

	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	looklookPaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cacheHomeNestPaymentThirdPaymentIdPrefix, data.Id)
	looklookPaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cacheHomeNestPaymentThirdPaymentSnPrefix, data.Sn)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, thirdPaymentRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.Sn, newData.DeleteTime, newData.DelState, newData.Version, newData.UserId, newData.PayMode, newData.TradeType, newData.TradeState, newData.PayTotal, newData.TransactionId, newData.TradeStateDesc, newData.OrderSn, newData.ServiceType, newData.PayStatus, newData.PayTime, newData.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, newData.Sn, newData.DeleteTime, newData.DelState, newData.Version, newData.UserId, newData.PayMode, newData.TradeType, newData.TradeState, newData.PayTotal, newData.TransactionId, newData.TradeStateDesc, newData.OrderSn, newData.ServiceType, newData.PayStatus, newData.PayTime, newData.Id, oldVersion)
	}, looklookPaymentThirdPaymentIdKey, looklookPaymentThirdPaymentSnKey)
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
