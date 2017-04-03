package models

import (
	"encoding/json"

	"github.com/astaxie/beego/orm"

	"gopkg.in/resty.v0"
)

func (this *Website) GetNewsItems() ([]*NewsItem, error) {
	verbose("Getting news items")
	newsItems := []*NewsItem{}
	// marshal this website struct to json
	requestBody, err := json.Marshal(this)
	if err != nil {
		return newsItems, err
	}
	// get the api url from the configurations
	url := SETTINGS.NewsItemsUrl
	// add time out to prevent some problems
	verbose("Setting up http client")
	res, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(url)

	// parse json from the byte array
	verbose("status: %v\nError: %v", res.StatusCode(), res.Error())
	result := map[string]interface{}{}
	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		verbose("Body unmarshalling error:\n%s\nRequest url: %s\nBody: %v", err.Error(), url, res.String())
		return newsItems, err
	}

	if result["items"] != nil {
		for _, item := range result["items"].([]interface{}) {
			it := item.(map[string]interface{})
			newItem := NewsItem{
				WebsiteName: it["website_name"].(string),
				WebsiteUrl:  it["website_url"].(string),
				Title:       it["title"].(string),
				Link:        it["link"].(string),
			}
			newsItems = append(newsItems, &newItem)
		}
	}

	// no error yet
	verbose("Returning %d items", len(newsItems))
	return newsItems, nil
}

func (this *Website) Save() bool {
	// check if item exists
	existingWebsite := []orm.Params{}
	o.Raw("SELECT * FROM website WHERE name = ? OR root_url = ?", this.Name, this.RootUrl).Values(&existingWebsite)
	if existingWebsite != nil {
		verbose("Website already exists in the database")
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
