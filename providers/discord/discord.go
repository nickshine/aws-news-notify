package discord

import (
	"aws-news-notify/providers"
	"fmt"
	"log"
	"net/http"
	"net/url"

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
func Init(configData []byte) {
	var c config
	if err := yaml.Unmarshal(configData, &c); err != nil {
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

	for _, v := range news {
		res, err := http.PostForm(p.WebhookURL, url.Values{
			"username": {"AWS News"},
			"content":  {fmt.Sprintf("Title: %v\nLink: %v\nDate: %v\n", v.Title, v.Link, v.PostDate)},
		})
		if err != nil {
			log.Println(err)
		}

		res.Body.Close()
	}
}
