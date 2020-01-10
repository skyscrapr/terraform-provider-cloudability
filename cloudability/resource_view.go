package cloudability

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

func resourceView() *schema.Resource {
	return &schema.Resource{
		Create: resourceViewCreate,
		Read: resourceViewRead,
		Update: resourceViewUpdate,
		Delete: resourceViewDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"title": {
				Type: schema.TypeString,
				Required: true,
				Description: "The name of the view as it will appear to the end users",
			},
			"shared_with_users": {
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
				Description: "The discrete list of users (by their unique identifier) that the view should be shared with",
			},
			"shared_with_organization": {
				Type: schema.TypeBool,
				Optional: true,
				Default: true,
				Description: "Whether the view should be accessible to the entire organization",
			},
			"owner_id": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Unique identifier for the user who created the view",
			},
			"filter": {
				Type: schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema {
						"field": &schema.Schema {
							Type: schema.TypeString,
							Required: true,
						},
						"comparator": &schema.Schema {
							Type: schema.TypeString,
							Required: true,
						},
						"value": &schema.Schema {
							Type: schema.TypeString,
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
	// TODO: Implement
	return resourceViewRead(d, meta)
}

func resourceViewRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.CloudabilityClient)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	view, err := client.Views.GetView(id)
	if err != nil {
		return err
	}

	if view != nil {
		d.Set("title", view.Title)
		d.Set("shared_with_users", view.SharedWithUsers)
		d.Set("shared_with_organization", view.SharedWithOrganization)
		d.Set("owner_id", view.OwnerId)
		d.Set("filters", flattenFilters(view.Filters))
		d.SetId(strconv.Itoa(view.Id))
	} else {
		// View not found. Remove from state
		d.SetId("")
	}
	return nil
}
 
func resourceViewUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.CloudabilityClient)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil
	}
	view := &cloudability.View{
		Id: id, 
		Title: d.Get("title").(string),
		// SharedWithUsers: d.Get("shared_with_users").(string),
		SharedWithOrganization: d.Get("shared_with_organization").(bool),
		// Filters: d.Get("filters")
	}
	err = client.Views.UpdateView(view)
	if err != nil {
		return err
	}
	return resourceViewRead(d, meta)
}

func resourceViewDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.CloudabilityClient)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil
	}
	return client.Views.DeleteView(id)
}
