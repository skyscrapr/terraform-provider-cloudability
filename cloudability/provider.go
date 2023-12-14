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
				Optional:    true,
				Description: "The apikey for API operations",
				DefaultFunc: schema.EnvDefaultFunc("CLOUDABILITY_APIKEY", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cloudability_account_verification": dataSourceAccountVerification(),
			"cloudability_users":                dataSourceUsers(),
			"cloudability_views":                dataSourceViews(),
			"cloudability_view":                 dataSourceView(),
			"cloudability_cluster_config":       dataSourceClusterConfig(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudability_master_account":   resourceMasterAccount(),
			"cloudability_linked_account":   resourceLinkedAccount(),
			"cloudability_business_mapping": resourceBusinessMapping(),
			"cloudability_business_metric":  resourceBusinessMetric(),
			"cloudability_view":             resourceView(),
			"cloudability_cluster":          resourceCluster(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c := NewConfig(d)
	return c.Client(), nil
}
