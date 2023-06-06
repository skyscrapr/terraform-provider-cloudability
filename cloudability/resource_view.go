package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
	"log"
)

func resourceView() *schema.Resource {
	return &schema.Resource{
		Create: resourceViewCreate,
		Read:   resourceViewRead,
		Update: resourceViewUpdate,
		Delete: resourceViewDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"title": {
				Type:        schema.TypeString,
				Required:    true,
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
				Optional:    true,
				Default:     true,
				Description: "Whether the view should be accessible to the entire organization",
			},
			"owner_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Unique identifier for the user who created the view",
			},
			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:     schema.TypeString,
							Required: true,
						},
						"comparator": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Description: "list of filter objects. If multiple filters are applied on the same dimension they are OR'd, however if they are on different dimensions they are AND'd. See below regarding filter specifics.",
			},
		},
	}
}

func resourceViewCreate(d *schema.ResourceData, meta interface{}) error {
	title := d.Get("title").(string)

	client := meta.(*cloudability.Client)
	log.Printf("[DEBUG] resourceViewCreate [title]: %q]", title)
	view := &cloudability.View{
		Title:                  title,
		SharedWithUsers:        inflateStrings(d.Get("shared_with_users").([]interface{})),
		SharedWithOrganization: d.Get("shared_with_organization").(bool),
		Filters:                inflateFilters(d.Get("filter").([]interface{})),
	}
	view, err := client.Views().NewView(view)
	if err != nil {
		return err
	}
	d.SetId(view.ID)
	return resourceViewRead(d, meta)
}

func resourceViewRead(d *schema.ResourceData, meta interface{}) error {
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

func resourceViewUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	view := &cloudability.View{
		ID:                     d.Id(),
		Title:                  d.Get("title").(string),
		SharedWithUsers:        d.Get("shared_with_users").([]string),
		SharedWithOrganization: d.Get("shared_with_organization").(bool),
		Filters:                inflateFilters(d.Get("filter").([]interface{})),
	}
	err := client.Views().UpdateView(view)
	if err != nil {
		return err
	}
	return resourceViewRead(d, meta)
}

func resourceViewDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	return client.Views().DeleteView(d.Id())
}
