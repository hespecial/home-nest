package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"home-nest/pkg/globalkey"
)

var _ HomestayModel = (*customHomestayModel)(nil)

type (
	// HomestayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayModel.
	HomestayModel interface {
		homestayModel
		FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Homestay, error)
		SelectBuilder() squirrel.SelectBuilder
	}

	customHomestayModel struct {
		*defaultHomestayModel
	}
)

// NewHomestayModel returns a model for the database table.
func NewHomestayModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HomestayModel {
	return &customHomestayModel{
		defaultHomestayModel: newHomestayModel(conn, c, opts...),
	}
}

func (m *customHomestayModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *customHomestayModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Homestay, error) {
	builder = builder.Columns(homestayRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Homestay
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
