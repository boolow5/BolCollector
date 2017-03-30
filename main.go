package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boolow5/BolCollector/models"
)

func main() {
	// start infinite loop to fetch news
	for {
		for _, site := range models.WEBSITES {
			items, err := site.GetNewsItems()
			if err != nil {
				log.Fatalln(err)
			}
			if err == nil {
				go models.SaveNews(items)
			}
			fmt.Println("Next update will takeplace after", models.SETTINGS.Delay, "second(s)")
		}
		time.Sleep(time.Second * time.Duration(models.SETTINGS.Delay))
	}
}
