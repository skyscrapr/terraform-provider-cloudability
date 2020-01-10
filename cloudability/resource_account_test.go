package cloudability

import (
	"os"
	"testing"
)

func TestResourceAccountRead(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	resource := resourceAccount()
	d := resource.Data(nil)
	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_ACCOUNTID"))
	d.Set("vendor_key", "aws")
	c := config.Client()
	resourceAccountRead(d, c)
}

func TestResourceAccountCreate(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	resource := resourceAccount()
	d := resource.Data(nil)
	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_ACCOUNTID"))
	d.Set("vendor_key", "aws")
	c := config.Client()
	err := resourceAccountCreate(d, c)
	if err != nil {
		t.Fail()
	}
}

func TestResourceAccountDelete(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	resource := resourceAccount()
	d := resource.Data(nil)
	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_ACCOUNTID"))
	d.Set("vendor_key", "aws")
	c := config.Client()
	resourceAccountDelete(d, c)
}