package menu

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/mysql/menu"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, menuData *CreateMenuData) (id int32, err error)
	Modify(ctx core.Context, id int32, menuData *UpdateMenuData) (err error)
	List(ctx core.Context, searchData *SearchData) (listData []*menu.Menu, err error)
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
