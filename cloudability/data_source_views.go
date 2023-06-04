package cloudability

import (
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

func dataSourceViews() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceViewsRead,
		Schema: map[string]*schema.Schema{
			"views": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"title": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the view as it will appear to the end users",
						},
						// TODO:
						// "shared_with_users": {
						// 	Type: schema.TypeList,
						// 	Optional: true,
						// 	Elem: &schema.Schema {
						// 		Type: schema.TypeString,
						// 	},
						// 	Description: "The discrete list of users (by their unique identifier) that the view should be shared with",
						// },
						"shared_with_organization": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the view should be accessible to the entire organization",
						},
						"owner_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier for the user who created the view",
						},
						"filter": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"field": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"comparator": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
							Description: "list of filter objects. If multiple filters are applied on the same dimension they are OR'd, however if they are on different dimensions they are AND'd. See below regarding filter specifics.",
						},
					},
				},
			},
		},
	}
}

func dataSourceViewsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	views, err := client.Views().GetViews()
	if err != nil {
		return err
	}

	if views != nil {
		var data = make([]map[string]interface{}, len(views))
		for i, view := range views {
			m := make(map[string]interface{})
			m["title"] = view.Title
			// m["shared_with_users"] = view.SharedWithUsers
			m["shared_with_organization"] = view.SharedWithOrganization
			m["owner_id"] = view.OwnerID
			m["filters"] = flattenFilters(view.Filters)
			data[i] = m
		}
		d.Set("views", data)
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	}

	return nil
}
