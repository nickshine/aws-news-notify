package main

import (
	"aws-news-notify/providers"
	"aws-news-notify/providers/discord"
	"fmt"
	"io/ioutil"
	"log"

	awsNews "github.com/circa10a/go-aws-news"
)

func init() {
	b, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	initializeProviders(b)
}

func initializeProviders(configData []byte) {
	discord.Init(configData)
}

func main() {

	news, err := awsNews.Yesterday()
	if err != nil {
		log.Fatal(err)
	}
	// news, _ := awsNews.Fetch(2020, 01)
	fmt.Println("News count: ", len(news))

	providers := providers.GetProviders()

	for _, p := range providers {
		p.Notify(news)
	}
}
