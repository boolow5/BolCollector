package models

import (
	"github.com/astaxie/beego/orm"
)

func (this *NewsItem) Save() bool {
	// check if item exists
	existingItems := []orm.Params{}
	o.Raw("SELECT * FROM news_item WHERE link = ?", this.Link).Values(&existingItems)
	// if not save it
	if existingItems != nil {
		verbose("Already existing in the database")
		return false
	}

	i, err := o.Insert(this)
	if err != nil {
		verbose(err.Error())
		return false
	}
	if i < 1 {
		verbose("No rows saved")
		return false
	}
	return true
}

func (this *NewsItem) TableName() string {
	return "news_item"
}
