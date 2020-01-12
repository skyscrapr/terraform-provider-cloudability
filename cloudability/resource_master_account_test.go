package cloudability

import (
	"os"
	"fmt"
	"testing"
)

func TestResourceMasterAccountRead(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	resource := resourceMasterAccount()
	d := resource.Data(nil)
	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID"))
	d.Set("vendor_key", "aws")
	c := config.Client()
	err := resourceMasterAccountRead(d, c)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestResourceMasterAccountCreate(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	resource := resourceMasterAccount()
	d := resource.Data(nil)
	accountId := os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID")
	d.Set("vendor_account_id", accountId)
	d.Set("vendor_key", "aws")
	d.Set("type", "aws_role")
	d.Set("bucket_name", fmt.Sprintf("cloudability-%s", accountId))
	d.Set("report_name", "Cloudability")
	d.Set("report_prefix", "CostAndUsageReports")

	c := config.Client()
	err := resourceMasterAccountCreate(d, c)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestResourceMasterAccountDelete(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	resource := resourceMasterAccount()
	d := resource.Data(nil)
	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID"))
	d.Set("vendor_key", "aws")
	c := config.Client()
	err := resourceMasterAccountDelete(d, c)
	if err != nil {
		t.Errorf(err.Error())
	}
}