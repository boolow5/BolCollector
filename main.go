package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func init() {
	// load and open config files
	config, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}

	// parse the config files
	settings := Settings{}
	json.Unmarshal(config, &settings)

	dir := "config"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	websites := []*Website{}
	for index, file := range files {
		if strings.HasPrefix(file.Name(), "site") {
			fmt.Println("Website configuration found", file.Name(), "index", index)
			fileData, err := ioutil.ReadFile(dir + "/" + file.Name())
			site := Website{}
			err = json.Unmarshal(fileData, &site)
			if err == nil {
				websites = append(websites, &site)
			}
		}
	}
	for _, site := range websites {
		fmt.Println(site.RootUrl)
		fmt.Println(site.Name)
		fmt.Println(site.Selector.TargetBase)
	}
}

func main() {
	// start infinite loop to fetch news

	// make a call to any site every DELAY_TIME from the settings
	// use the basic configurations from config.json
	// with every call use the json sites.json from the file
}
