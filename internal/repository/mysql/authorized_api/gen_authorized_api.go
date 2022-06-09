///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package authorized_api

import (
	"fmt"

	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"gorm.io/gorm"
)

func NewQueryBuilder() *authorizedApiQueryBuilder {
	return new(authorizedApiQueryBuilder)
}

type authorizedApiQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *authorizedApiQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *authorizedApiQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *authorizedApiQueryBuilder) WhereBusinessKey(p mysql.Predicate, value string) *authorizedApiQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", p),
		value,
	})
	return qb
}

func (qb *authorizedApiQueryBuilder) OrderById(asc bool) *authorizedApiQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *authorizedApiQueryBuilder) QueryAll(db *gorm.DB) ([]*AuthorizedApi, error) {
	var ret []*AuthorizedApi
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *authorizedApiQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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
