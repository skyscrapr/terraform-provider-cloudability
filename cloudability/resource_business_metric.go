package cloudability

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

func resourceBusinessMetric() *schema.Resource {
	return &schema.Resource{
		Create: resourceBusinessMetricCreate,
		Read:   resourceBusinessMetricRead,
		Update: resourceBusinessMetricUpdate,
		Delete: resourceBusinessMetricDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"index": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"number_format": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(val any, key string) (warns []string, errs []error) {
					switch val.(string) {
					case
						"currency",
						"number":
						return
					}
					errs = append(errs, fmt.Errorf("invalid value for number_format. Must ne 'currency' or 'number'"))
					return
				},
				Default: "number",
			},
			"default_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_value_expression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pre_match_expression": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceBusinessMetricCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	payload := cloudability.BusinessMapping{
		Kind:                   "BUSINESS_METRIC",
		Name:                   d.Get("name").(string),
		NumberFormat:           d.Get("number_format").(string),
		PreMatchExpression:     d.Get("pre_match_expression").(string),
		DefaultValueExpression: d.Get("default_value_expression").(string),
		DefaultValue:           d.Get("default_value").(string),
		Statements:             inflateStatements(d.Get("statement").([]interface{})),
	}
	businessMapping, err := client.BusinessMappings().NewBusinessMetric(&payload)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	tflog.Info(ctx, fmt.Sprintf("New business metric created with index: %d", businessMapping.Index))
	d.SetId(strconv.Itoa(businessMapping.Index))
	return resourceBusinessMetricRead(d, meta)
}

func resourceBusinessMetricRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	index, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	ctx := context.TODO()
	tflog.Info(ctx, fmt.Sprintf("Reading business metric with index: %d", index))
	businessMapping, err := client.BusinessMappings().GetBusinessMetric(index)
	if err != nil {
		return err
	}

	if businessMapping != nil {
		d.Set("index", businessMapping.Index)
		d.Set("name", businessMapping.Name)
		d.Set("number_format", businessMapping.NumberFormat)
		d.Set("default_value", businessMapping.DefaultValue)
		d.Set("pre_match_expression", businessMapping.PreMatchExpression)
		d.Set("statement", flattenStatements(businessMapping.Statements))
		d.Set("updated_at", businessMapping.UpdatedAt)
		d.SetId(strconv.Itoa(businessMapping.Index))
	}
	return nil
}

func resourceBusinessMetricUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil
	}
	payload := cloudability.BusinessMapping{
		Kind:                   "BUSINESS_METRIC",
		Index:                  d.Get("index").(int),
		Name:                   d.Get("name").(string),
		NumberFormat:           d.Get("number_format").(string),
		PreMatchExpression:     d.Get("pre_match_expression").(string),
		DefaultValueExpression: d.Get("default_value_expression").(string),
		Statements:             inflateStatements(d.Get("statement").([]interface{})),
	}
	err = client.BusinessMappings().UpdateBusinessMetric(&payload)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	tflog.Info(ctx, fmt.Sprintf("Updating business metric with index: %d", id))
	return resourceBusinessMetricRead(d, meta)
}

func resourceBusinessMetricDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	ctx := context.TODO()
	tflog.Info(ctx, fmt.Sprintf("Deleting business metric with index: %d", id))
	err = client.BusinessMappings().DeleteBusinessMetric(id)
	if err != nil {
		// Ignore 404 errors (No resource found)
		var apiError cloudability.APIError
		jsonErr := json.Unmarshal([]byte(err.Error()), &apiError)
		if jsonErr == nil && apiError.Error.Status == 404 {
			ctx := context.TODO()
			tflog.Info(ctx, "resourceBusinessMetricDelete Resource not found. Ignoring")
			return nil
		}
	}
	return err
}
