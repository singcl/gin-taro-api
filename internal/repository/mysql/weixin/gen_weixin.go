package weixin

import (
	"fmt"

	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Weixin {
	return new(Weixin)
}

type weixinQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func NewQueryBuilder() *weixinQueryBuilder {
	return new(weixinQueryBuilder)
}

func (qb *weixinQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *weixinQueryBuilder) First(db *gorm.DB) (*Weixin, error) {
	ret := &Weixin{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *weixinQueryBuilder) QueryOne(db *gorm.DB) (*Weixin, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *weixinQueryBuilder) QueryAll(db *gorm.DB) ([]*Weixin, error) {
	var ret []*Weixin
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *weixinQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *weixinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *weixinQueryBuilder) WhereOpenid(p mysql.Predicate, value string) *weixinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "openid", p),
		value,
	})
	return qb
}
func (qb *weixinQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Weixin{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *weixinQueryBuilder) WhereUnionid(p mysql.Predicate, value string) *weixinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "unionid", p),
		value,
	})
	return qb
}

func (qb *weixinQueryBuilder) WhereUsername(p mysql.Predicate, value string) *weixinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", p),
		value,
	})
	return qb
}

func (qb *weixinQueryBuilder) WhereNickname(p mysql.Predicate, value string) *weixinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", p),
		value,
	})
	return qb
}

func (qb *weixinQueryBuilder) WhereAvatarUrl(p mysql.Predicate, value string) *weixinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "avatar_url", p),
		value,
	})
	return qb
}

func (qb *weixinQueryBuilder) WhereMobile(p mysql.Predicate, value string) *weixinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mobile", p),
		value,
	})
	return qb
}

func (qb *weixinQueryBuilder) WhereIsUsed(p mysql.Predicate, value int32) *weixinQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", p),
		value,
	})
	return qb
}

func (t *Weixin) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}
