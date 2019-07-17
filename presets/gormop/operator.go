package gormop

import (
	"fmt"
	"strings"

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

func (op *dataOperatorImpl) Search(obj interface{}, params *presets.SearchParams) (r interface{}, err error) {
	wh := op.db
	if len(params.KeywordColumns) > 0 && len(params.Keyword) > 0 {
		var segs []string
		var args []interface{}
		for _, c := range params.KeywordColumns {
			segs = append(segs, fmt.Sprintf("%s ilike ?", c))
			args = append(args, fmt.Sprintf("%%%s%%", params.Keyword))
		}
		wh = wh.Where(strings.Join(segs, " OR "), args...)
	}
	err = wh.Find(obj).Error
	if err != nil {
		return
	}
	r = obj
	return
}

func (op *dataOperatorImpl) Fetch(obj interface{}, id string) (r interface{}, err error) {
	err = op.db.Model(obj).Find(obj, "id = ?", id).Error
	if err != nil {
		return
	}
	r = obj
	return
}

func (op *dataOperatorImpl) UpdateField(obj interface{}, id string, fieldName string, value interface{}) (err error) {
	err = op.db.Model(obj).UpdateColumn(fieldName, value).Error
	return
}

func (op *dataOperatorImpl) Save(obj interface{}, id string) (err error) {
	err = op.db.Save(obj).Error
	return
}
