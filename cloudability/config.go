package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

// Config - Cloudability Provider config
type Config struct {
	APIKey string
}

// NewConfig - Create Cloudability Provider Config
func NewConfig(d *schema.ResourceData) *Config {
	c := &Config{
		APIKey: d.Get("apikey").(string),
	}
	return c
}

// Client - Create Cloudability API Client
func (c *Config) Client() *cloudability.Client {
	client := cloudability.NewClient(c.APIKey)
	return client
}
