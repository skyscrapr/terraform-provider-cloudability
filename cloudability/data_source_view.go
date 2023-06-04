package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

func dataSourceView() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceViewRead,
		Schema: map[string]*schema.Schema{
			"title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the view as it will appear to the end users",
			},
			"shared_with_users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The discrete list of users (by their unique identifier) that the view should be shared with",
			},
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
			"filters": {
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
	}
}

func dataSourceViewRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	view, err := client.Views().GetView(d.Id())
	if err != nil {
		return err
	}

	if view != nil {
		d.Set("title", view.Title)
		d.Set("shared_with_users", view.SharedWithUsers)
		d.Set("shared_with_organization", view.SharedWithOrganization)
		d.Set("owner_id", view.OwnerID)
		d.Set("filters", flattenFilters(view.Filters))
		d.SetId(view.ID)
	}
	return nil
}
