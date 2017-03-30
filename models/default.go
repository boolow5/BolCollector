package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DEBUG        bool
	AUTO_MIGRATE bool
	WEBSITES     []*Website
	SETTINGS     Settings
	o            orm.Ormer
)

type (
	Settings struct {
		Delay            int    `json:"delay"`
		NewsItemsUrl     string `json:"news_items_url"`
		NewsDetailUrl    string `json:"news_detail_url"`
		DatabaseFileName string `json:"database_file_name"`
	}
	Selector struct {
		Base       string `json:"base"`
		TargetBase string `json:"target_base"`
		TargetText string `json:"target_text"`
		TargetLink string `json:"target_link"`
	}
	Website struct {
		Name     string    `json:"name"`
		RootUrl  string    `json:"root_url"`
		Selector *Selector `json:"selector"`
	}
	NewsItem struct {
		Id          int    `json:"id" orm:"auto"`
		Title       string `json:"title" orm:"size(500)"`
		Link        string `json:"link" orm:"unique"`
		WebsiteName string `json:"website_name"`
		WebsiteUrl  string `json:"website_url"`
	}
)

func init() {
	DEBUG = false
	AUTO_MIGRATE = false
	// load and open config files
	config, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}

	// parse the config files
	json.Unmarshal(config, &SETTINGS)
	verbose("SETTINGS: %v", SETTINGS)

	dir := "config"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "site") {
			fileData, err := ioutil.ReadFile(dir + "/" + file.Name())
			site := Website{}
			err = json.Unmarshal(fileData, &site)
			if err == nil {
				WEBSITES = append(WEBSITES, &site)
			}
		}
	}

	// database access
	DatabaseFileName := SETTINGS.DatabaseFileName
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", DatabaseFileName)
	orm.RegisterModel(new(NewsItem))
	o = orm.NewOrm()

	// auto create database tables;
	if AUTO_MIGRATE {
		name := "default"                      // Database alias.
		force := true                          // Drop table and re-create.
		log := true                            // Print log
		err := orm.RunSyncdb(name, force, log) // Sync with database
		if err != nil {
			verbose(err.Error())
		}
	}
}

func SaveNews(items []*NewsItem) {
	if len(items) < 1 {
		return
	}
	savedItems := 0
	skippedItems := 0
	for _, item := range items {
		verbose("Saving item %s", item.Link)
		if item.Save() {
			savedItems += 1
		} else {
			skippedItems += 1
		}
	}
	fmt.Printf("Saved %d items\nskipped %d items\n", savedItems, skippedItems)
}

func verbose(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format+"\n", args...)
	}
}
