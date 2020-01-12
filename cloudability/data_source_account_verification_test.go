package cloudability

import (
	"os"
	"testing"
)


func TestDatasourceAccountVerificationRead(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	dataSource := dataSourceAccountVerification()
	d := dataSource.Data(nil)
	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID"))
	d.Set("vendor_key", "aws")
	d.Set("retry_count", 1)
	d.Set("retry_wait", 1)	
	c := config.Client()
	err := dataSourceAccountVerificationRead(d, c)
	if err != nil {
		t.Errorf(err.Error())
	}
}