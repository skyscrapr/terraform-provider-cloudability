package cloudability

import (
	"os"
	"testing"
)

func TestResourceLinkedAccountRead(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	resource := resourceLinkedAccount()
	d := resource.Data(nil)
	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_ACCOUNTID"))
	d.Set("vendor_key", "aws")
	c := config.Client()
	err := resourceLinkedAccountRead(d, c)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestResourceLinkedAccountCreate(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	resource := resourceLinkedAccount()
	d := resource.Data(nil)
	accountId := os.Getenv("CLOUDABILITY_ACCOUNTID")
	d.Set("vendor_account_id", accountId)
	d.Set("vendor_key", "aws")
	d.Set("type", "aws_role")

	c := config.Client()
	err := resourceLinkedAccountCreate(d, c)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestResourceLinkedAccountDelete(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	resource := resourceLinkedAccount()
	d := resource.Data(nil)
	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_ACCOUNTID"))
	d.Set("vendor_key", "aws")
	c := config.Client()
	err := resourceLinkedAccountDelete(d, c)
	if err != nil {
		t.Errorf(err.Error())
	}
}