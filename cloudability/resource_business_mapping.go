package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
	"log"
	"strconv"
)

func resourceBusinessMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceBusinessMappingCreate,
		Read:   resourceBusinessMappingRead,
		Update: resourceBusinessMappingUpdate,
		Delete: resourceBusinessMappingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"index": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"kind": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"default_value": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"statement": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"match_expression": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value_expression": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceBusinessMappingCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	businessMapping := &cloudability.BusinessMapping{
		Name:         d.Get("name").(string),
		Kind:         d.Get("kind").(string),
		DefaultValue: d.Get("default_value").(string),
		Statements:   inflateStatements(d.Get("statement").([]interface{})),
	}
	newBusinessMapping, err := client.BusinessMappings().NewBusinessMapping(businessMapping)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] New business mapping created with index: %d", newBusinessMapping.Index)
	d.SetId(strconv.Itoa(newBusinessMapping.Index))
	return resourceBusinessMappingRead(d, meta)
}

func resourceBusinessMappingRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	index, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	businessMapping, err := client.BusinessMappings().GetBusinessMapping(index)
	if err != nil {
		return err
	}

	if businessMapping != nil {
		d.Set("index", businessMapping.Index)
		d.Set("kind", businessMapping.Kind)
		d.Set("name", businessMapping.Name)
		d.Set("default_value", businessMapping.DefaultValue)
		d.Set("statement", flattenStatements(businessMapping.Statements))
		d.Set("updated_at", businessMapping.UpdatedAt)
		d.SetId(strconv.Itoa(businessMapping.Index))
	}
	return nil
}

func resourceBusinessMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil
	}
	businessMapping := &cloudability.BusinessMapping{
		Index:        id,
		Kind:         d.Get("kind").(string),
		Name:         d.Get("name").(string),
		DefaultValue: d.Get("default_value").(string),
		// TODO: Statements:
	}
	err = client.BusinessMappings().UpdateBusinessMapping(businessMapping)
	if err != nil {
		return err
	}
	return resourceBusinessMappingRead(d, meta)
}

func resourceBusinessMappingDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil
	}
	return client.BusinessMappings().DeleteBusinessMapping(id)
}
