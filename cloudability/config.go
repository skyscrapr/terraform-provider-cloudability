package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

// Config - Provider Config
type Config struct {
	APIKey string
}

// NewConfig - Return a new Config instance
func NewConfig(d *schema.ResourceData) *Config {
	c := &Config{
		APIKey: d.Get("apikey").(string),
	}
	return c
}

// Client - Retrun a new Cloudabiity Client instance
func (c *Config) Client() *cloudability.Client {
	client := cloudability.NewClient(c.APIKey)
	return client
}
