package main

import (
	"aws-news-notify/providers"
	"fmt"
	"log"

	awsNews "github.com/circa10a/go-aws-news"
)

func init() {
	fmt.Println("main init")
}

func main() {

	// news, err := awsNews.Fetch(2020, 01)
	news, err := awsNews.Yesterday()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("News count: ", len(news))

	providers := providers.GetProviders()

	// fmt.Printf("%#v\n", providers)
	for _, p := range providers {
		p.Notify(news)
	}
}
