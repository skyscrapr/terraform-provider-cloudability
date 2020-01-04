package cloudability

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

func resourceBusinessMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceBusinessMappingCreate,
		Read: resourceBusinessMappingRead,
		Update: resourceBusinessMappingUpdate,
		Delete: resourceBusinessMappingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"index": {
				Type: schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"kind": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"statement": {
				Type: schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"match_expression": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"value_expression": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"updated_at": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceBusinessMappingCreate(d *schema.ResourceData, meta interface{}) error {
	// TODO: Implement
	return resourceBusinessMappingRead(d, meta)
}

func resourceBusinessMappingRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.CloudabilityClient)
	index, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	businessMapping, err := client.BusinessMappings.BusinessMapping(index)
	if err != nil {
		return err
	}

	if businessMapping != nil {
		d.Set("index", businessMapping.Index)
		d.Set("kind", businessMapping.Kind)
		d.Set("name", businessMapping.Name)
		d.Set("statement", flattenStatements(businessMapping.Statements))
		d.Set("updated_at", businessMapping.UpdatedAt)
		d.SetId(strconv.Itoa(businessMapping.Index))
	} else {
		// BusinessMapping not found. Remove from state
		d.SetId("")
	}
	return nil
}
 
func resourceBusinessMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	// TODO: Implement
	return resourceBusinessMappingRead(d, meta)
}

func resourceBusinessMappingDelete(d *schema.ResourceData, meta interface{}) error {
	// TODO: Implement
	return nil
}

func flattenStatements(in []cloudability.BusinessMappingStatement) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["match_expression"] = v.MatchExpression
		m["value_expression"] = v.ValueExpression
		out[i] = m
	}
	return out
}