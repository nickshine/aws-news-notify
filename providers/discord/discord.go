package discord

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"aws-news-notify/providers"

	awsNews "github.com/circa10a/go-aws-news"
	"gopkg.in/yaml.v2"
)

// Provider is an implementation of the `aws-news-notify` Provider interface.
type Provider struct {
	IsEnabled  bool   `yaml:"enabled"`
	WebhookURL string `yaml:"webhookURL"`
}

// Init initializes the provider from the provided config.
func Init(configData []byte) {

	d := struct {
		Provider Provider `yaml:"discord"`
	}{}
	if err := yaml.Unmarshal(configData, &d); err != nil {
		panic(err)
	}

	providers.RegisterProvider(&d.Provider)
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
