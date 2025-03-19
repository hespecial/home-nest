package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"home-nest/pkg/globalkey"

	"github.com/Masterminds/squirrel"
)

var _ HomestayActivityModel = (*customHomestayActivityModel)(nil)

type (
	// HomestayActivityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomestayActivityModel.
	HomestayActivityModel interface {
		homestayActivityModel
		SelectBuilder() squirrel.SelectBuilder
		FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayActivity, error)
	}

	customHomestayActivityModel struct {
		*defaultHomestayActivityModel
	}
)

// NewHomestayActivityModel returns a model for the database table.
func NewHomestayActivityModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HomestayActivityModel {
	return &customHomestayActivityModel{
		defaultHomestayActivityModel: newHomestayActivityModel(conn, c, opts...),
	}
}

func (m *customHomestayActivityModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *customHomestayActivityModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayActivity, error) {
	builder = builder.Columns(homestayActivityRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayActivity
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
