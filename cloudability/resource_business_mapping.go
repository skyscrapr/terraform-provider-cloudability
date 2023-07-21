package cloudability

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

func resourceBusinessMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceBusinessMappingCreate,
		Read:   resourceBusinessMappingRead,
		Update: resourceBusinessMappingUpdate,
		Delete: resourceBusinessMappingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Read:   schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},
	}
}

func resourceBusinessMappingCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	client.SetTimeout(d.Timeout("Create"))
	payload := cloudability.BusinessMapping{
		Name:         d.Get("name").(string),
		Kind:         d.Get("kind").(string),
		DefaultValue: d.Get("default_value").(string),
		Statements:   inflateStatements(d.Get("statement").([]interface{})),
	}

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	businessMapping, err := client.BusinessMappings().NewBusinessDimension(&payload)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	tflog.Info(ctx, fmt.Sprintf("New business dimension created with index: %d", businessMapping.Index))
	d.SetId(strconv.Itoa(businessMapping.Index))

	return resourceBusinessMappingRead(d, meta)
}

func resourceBusinessMappingRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	client.SetTimeout(d.Timeout("Read"))
	index, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	ctx := context.TODO()
	tflog.Info(ctx, fmt.Sprintf("Reading business dimension with index: %d", index))
	businessMapping, err := client.BusinessMappings().GetBusinessDimension(index)
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
	client.SetTimeout(d.Timeout("Update"))
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil
	}
	payload := cloudability.BusinessMapping{
		Index:        d.Get("index").(int),
		Kind:         d.Get("kind").(string),
		Name:         d.Get("name").(string),
		DefaultValue: d.Get("default_value").(string),
		Statements:   inflateStatements(d.Get("statement").([]interface{})),
	}
	err = client.BusinessMappings().UpdateBusinessDimension(&payload)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	tflog.Info(ctx, fmt.Sprintf("Updating business dimension with index: %d", id))
	return resourceBusinessMappingRead(d, meta)
}

func resourceBusinessMappingDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	client.SetTimeout(d.Timeout("Delete"))
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	ctx := context.TODO()
	tflog.Info(ctx, fmt.Sprintf("Deleting business mapping with index: %d", id))
	err = client.BusinessMappings().DeleteBusinessDimension(id)
	if err != nil {
		// Ignore 404 errors (No resource found)
		var apiError cloudability.APIError
		jsonErr := json.Unmarshal([]byte(err.Error()), &apiError)
		if jsonErr == nil && apiError.Error.Status == 404 {
			ctx := context.TODO()
			tflog.Info(ctx, "resourceBusinessDimensionDelete Resource not found. Ignoring")
			return nil
		}
	}
	return err
}
