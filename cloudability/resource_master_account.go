package cloudability

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
	"log"
	"time"
)

func resourceMasterAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceMasterAccountCreate,
		Read:   resourceMasterAccountRead,
		Delete: resourceMasterAccountDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vendor_account_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name given to your AWS account",
			},
			"vendor_account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "12 digit string corresponding to your AWS account ID",
			},
			"vendor_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "aws",
				ForceNew:    true,
				Description: "'aws'",
			},
			"verification": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Examples: unverified, verified, error",
						},
						"last_verification_attempted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date timestamp, example: 1970-01-01T00:00:00.000Z",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error message for credentials in error state",
						},
					},
				},
				Description: "Object containing details of verification state",
			},
			"authorization": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "'aws_role' or 'aws_user'",
						},
						"role_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "currently hardcoded to 'CloudabilityRole'",
						},
						"external_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The external ID used to prevent confused deputies. Generated by Cloudability",
						},
					},
				},
				Description: "Object contain vendor specific authorization details",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "aws_role",
				ForceNew:    true,
				Description: "'aws_role' or 'aws_user'",
			},
			"external_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The external ID used to prevent confused deputies. Generated by Cloudability",
			},
			"parent_account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "12 digit string representing parent's account ID (if current cred is a linked account)",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date timestamp corresponding to cloudability credential creation time",
			},
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of s3 bucket containing cost usage reports",
			},
			"report_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of cost and usage report",
			},
			"report_prefix": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cost and usage report prefix",
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Read:   schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},
	}
}

func resourceMasterAccountCreate(d *schema.ResourceData, meta interface{}) error {
	vendorKey := d.Get("vendor_key").(string)
	accountID := d.Get("vendor_account_id").(string)
	credType := d.Get("type").(string)
	bucketName := d.Get("bucket_name").(string)
	reportName := d.Get("report_name").(string)
	reportPrefix := d.Get("report_prefix").(string)

	client := meta.(*cloudability.Client)
	client.SetTimeout(d.Timeout("Create"))
	log.Printf("[DEBUG] resourceAccountCreate NewAccount [account_id: %q]", accountID)
	params := &cloudability.NewMasterAccountParams{
		NewLinkedAccountParams: &cloudability.NewLinkedAccountParams{
			VendorAccountID: accountID,
			Type:            credType,
		},
		BucketName: bucketName,
		CostAndUsageReport: &cloudability.CostAndUsageReport{
			Name:   reportName,
			Prefix: reportPrefix,
		},
	}
	_, err := client.Vendors().NewMasterAccount(vendorKey, params)
	if err != nil {
		return err
	}
	return resourceMasterAccountRead(d, meta)
}

func resourceMasterAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	client.SetTimeout(d.Timeout("Read"))
	vendorKey := d.Get("vendor_key").(string)
	accountID := d.Get("vendor_account_id").(string)
	log.Printf("[DEBUG] resourceMasterAccountRead [account_id: %q]", accountID)
	account, err := client.Vendors().GetAccount(vendorKey, accountID)
	if err != nil {
		// Ignore 404 errors (No account found)
		var apiError cloudability.APIError
		jsonErr := json.Unmarshal([]byte(err.Error()), &apiError)
		if jsonErr == nil && apiError.Error.Status == 404 {
			log.Print("[DEBUG] resourceMasterAccountRead Account not found. Ignoring")
			err = nil
		} else {
			return err
		}
	}

	if account != nil {
		d.Set("vendor_account_name", account.VendorAccountName)
		d.Set("vendor_account_id", account.VendorAccountID)
		d.Set("vendor_key", account.VendorKey)
		d.Set("parent_account_id", account.ParentAccountID)
		d.Set("created_at", account.CreatedAt)

		if account.Verification != nil {
			d.Set("verification", flattenVerification(account.Verification))
		}
		if account.Authorization != nil {
			d.Set("authorization", flattenAuthorization(account.Authorization))
			d.Set("external_id", account.Authorization.ExternalID)
			if account.Authorization.CostAndUsageReport != nil {
				d.Set("report_name", account.Authorization.CostAndUsageReport.Name)
				d.Set("report_prefix", account.Authorization.CostAndUsageReport.Prefix)
				d.Set("bucket_name", account.Authorization.BucketName)
			}
		}
		d.SetId(account.ID)
	}
	return nil
}

func resourceMasterAccountDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	client.SetTimeout(d.Timeout("Delete"))
	vendorKey := d.Get("vendor_key").(string)
	accountID := d.Get("vendor_account_id").(string)
	err := client.Vendors().DeleteAccount(vendorKey, accountID)
	return err
}
