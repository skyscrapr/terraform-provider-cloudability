package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

type Config struct {
	ApiKey string
}

func NewConfig(d *schema.ResourceData) *Config {
	c := &Config{
		ApiKey: d.Get("apikey").(string),
	}
	return c
}

func (c *Config) Client() *cloudability.Client {
	client := cloudability.NewClient(c.ApiKey)
	return client
}
