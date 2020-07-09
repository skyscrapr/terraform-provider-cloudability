package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
	"log"
)

func dataSourceRightsizingResource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRightsizingResourceRead,
		Schema: map[string]*schema.Schema{
			"vendor": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cloud Vendor",
			},
			"service": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cloud Service",
			},
			"resource_identifier": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cloud instance id",
			},
			"default_instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Default value for instance type",
			},
			"recommendations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"nodeType": {
							Type:     schema.TypeString,
							Required: true,
						},
						"risk": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
				Description: "List of recommendation objects.",
			},
		},
	}
}

func dataSourceRightsizingResourceRead(d *schema.ResourceData, meta interface{}) error {
	vendor := d.Get("vendor").(string)
	service := d.Get("service").(string)
	resourceIdentifier := d.Get("resource_identifier").(string)

	client := meta.(*cloudability.Client)
	log.Printf("[DEBUG] resourceRightsizingResourceRead [resourceIdentifier: %q]", resourceIdentifier)
	resource, err := client.Rightsizing().GetResource(vendor, service, resourceIdentifier)
	// TODO: Handle default value for resource if the rightsizing instance or recommendation not found.
	// Need a test for this too.
	// Don't know what is returned if the instance does not exist.
	if err != nil {
		log.Printf("[DEBUG] Could not get Cloudability Rightsizing Resource: %q", err)
		return err
	}
	d.SetId(resource.ResourceIdentifier)
	d.Set("service", resource.Service)
	return nil
}
