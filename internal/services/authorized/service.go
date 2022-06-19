package authorized

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/mysql/authorized"
	"github.com/singcl/gin-taro-api/internal/repository/mysql/authorized_api"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Detail(ctx core.Context, id int32) (info *authorized.Authorized, err error)
	DetailByKey(ctx core.Context, key string) (data *CacheAuthorizedData, err error)
	Create(ctx core.Context, authorizedData *CreateAuthorizedData) (id int32, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*authorized.Authorized, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	Delete(ctx core.Context, id int32) (err error)

	ListAPI(ctx core.Context, searchAPIData *SearchAPIData) (listData []*authorized_api.AuthorizedApi, err error)
	CreateAPI(ctx core.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int32, err error)
	DeleteAPI(ctx core.Context, id int32) (err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
