package weixin

import "gorm.io/gorm"

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
