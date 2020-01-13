package providers

import (
	awsNews "github.com/circa10a/go-aws-news"
)

// Provider is implemented in each webhook provider in providers/.
type Provider interface {
	Notify(awsNews.Announcements)
	GetName() string
	Enabled() bool
}

// Providers is a slice of registered providers.
var providers []Provider

// GetProviders returns a list of registered providers.
func GetProviders() []Provider {
	return providers
}

// RegisterProvider adds the provider to the list of registered providers.
func RegisterProvider(p Provider) {
	if p.Enabled() {
		providers = append(providers, p)
	}
}
