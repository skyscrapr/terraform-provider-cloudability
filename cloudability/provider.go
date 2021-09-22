package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider - Terraform Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The apikey for API operations",
				DefaultFunc: schema.EnvDefaultFunc("APIKEY", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cloudability_account_verification": dataSourceAccountVerification(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudability_master_account":   resourceMasterAccount(),
			"cloudability_linked_account":   resourceLinkedAccount(),
			"cloudability_business_mapping": resourceBusinessMapping(),
			"cloudability_user":             resourceUser(),
			"cloudability_view":             resourceView(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c := NewConfig(d)
	return c.Client(), nil
}
