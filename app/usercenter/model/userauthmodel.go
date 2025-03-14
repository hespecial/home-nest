package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"home-nest/pkg/globalkey"
	"time"
)

var _ UserAuthModel = (*customUserAuthModel)(nil)

type (
	// UserAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthModel.
	UserAuthModel interface {
		userAuthModel
		Insert(ctx context.Context, session sqlx.Session, data *UserAuth) (sql.Result, error)
	}

	customUserAuthModel struct {
		*defaultUserAuthModel
	}
)

// NewUserAuthModel returns a model for the database table.
func NewUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserAuthModel {
	return &customUserAuthModel{
		defaultUserAuthModel: newUserAuthModel(conn, c, opts...),
	}
}

func (m *customUserAuthModel) Insert(ctx context.Context, session sqlx.Session, data *UserAuth) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	homeNestUsercenterUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheHomeNestUsercenterUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	homeNestUsercenterUserAuthIdKey := fmt.Sprintf("%s%v", cacheHomeNestUsercenterUserAuthIdPrefix, data.Id)
	homeNestUsercenterUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheHomeNestUsercenterUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userAuthRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.UserId, data.AuthKey, data.AuthType)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.UserId, data.AuthKey, data.AuthType)
	}, homeNestUsercenterUserAuthAuthTypeAuthKeyKey, homeNestUsercenterUserAuthIdKey, homeNestUsercenterUserAuthUserIdAuthTypeKey)
	return ret, err
}
