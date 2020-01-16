package discord

import (
	"aws-news-notify/providers"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	awsNews "github.com/circa10a/go-aws-news"
	"gopkg.in/yaml.v2"
)

type config struct {
	Providers struct {
		Provider Provider `yaml:"discord"`
	} `yaml:"providers"`
}

// Provider is an implementation of the `aws-news-notify` Provider interface.
type Provider struct {
	IsEnabled  bool   `yaml:"enabled"`
	WebhookURL string `yaml:"webhookURL"`
}

// Init initializes the provider from the provided config.
func init() {
	var c config
	if err := yaml.Unmarshal(providers.Config, &c); err != nil {
		panic(err)
	}

	providers.RegisterProvider(&c.Providers.Provider)
}

// Enabled returns true if the provider is enabled in the config.
func (p *Provider) Enabled() bool {
	return p.IsEnabled
}

// GetName returns the Provider's name.
func (*Provider) GetName() string {
	return "discord"
}

// Notify is the function executed to POST to a provider's webhook url.
func (p *Provider) Notify(news awsNews.Announcements) {

	var b strings.Builder
	for _, v := range news {
		b.WriteString(fmt.Sprintf("**Title:** *%v*\n**Link:** %v\n**Date:** %v\n", v.Title, v.Link, v.PostDate))
	}

	res, err := http.PostForm(p.WebhookURL, url.Values{
		"username": {"AWS News"},
		"content":  {b.String()},
	})
	if err != nil {
		log.Println(err)
	}
	res.Body.Close()
}
