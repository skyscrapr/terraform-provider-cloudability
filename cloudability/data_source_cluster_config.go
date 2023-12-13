package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

func dataSourceClusterConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceClusterConfigRead,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceClusterConfigRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)

	clusterID := d.Get("cluster_id").(string)
	config, err := client.Containers().GetClusterConfig(clusterID)
	if err != nil {
		return err
	}

	d.SetId(clusterID)
	d.Set("config", config)
	return nil
}
