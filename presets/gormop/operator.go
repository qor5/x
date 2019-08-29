package gormop

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/sunfmin/bran/ui"

	"github.com/jinzhu/gorm"
	"github.com/sunfmin/bran/presets"
)

func DataOperator(db *gorm.DB) (r presets.DataOperator) {
	r = &dataOperatorImpl{db: db}
	return
}

type dataOperatorImpl struct {
	db *gorm.DB
}

func (op *dataOperatorImpl) Search(obj interface{}, params *presets.SearchParams, ctx *ui.EventContext) (r interface{}, totalCount int, err error) {
	wh := op.db.Model(obj)
	if len(params.KeywordColumns) > 0 && len(params.Keyword) > 0 {
		var segs []string
		var args []interface{}
		for _, c := range params.KeywordColumns {
			segs = append(segs, fmt.Sprintf("%s ilike ?", c))
			args = append(args, fmt.Sprintf("%%%s%%", params.Keyword))
		}
		wh = wh.Where(strings.Join(segs, " OR "), args...)
	}

	for _, cond := range params.SQLConditions {
		wh = wh.Where(cond.Query, cond.Args...)
	}

	err = wh.Count(&totalCount).Error
	if err != nil {
		return
	}

	if params.PerPage > 0 {
		wh = wh.Limit(params.PerPage)
		page := params.Page
		if page == 0 {
			page = 1
		}
		offset := (page - 1) * params.PerPage
		wh = wh.Offset(offset)
	}

	orderBy := params.OrderBy
	if len(orderBy) > 0 {
		wh = wh.Order(orderBy)
	}

	err = wh.Find(obj).Error
	if err != nil {
		return
	}
	r = reflect.ValueOf(obj).Elem().Interface()
	return
}

func (op *dataOperatorImpl) Fetch(obj interface{}, id string, ctx *ui.EventContext) (r interface{}, err error) {
	err = op.db.Model(obj).Find(obj, "id = ?", id).Error
	if err != nil {
		return
	}
	r = obj
	return
}

func (op *dataOperatorImpl) Save(obj interface{}, id string, ctx *ui.EventContext) (err error) {
	err = op.db.Save(obj).Error
	return
}

func (op *dataOperatorImpl) Delete(obj interface{}, id string, ctx *ui.EventContext) (err error) {
	err = op.db.Delete(obj, id).Error
	return
}
