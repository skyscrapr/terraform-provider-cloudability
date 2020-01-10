package cloudability

import (
	"os"
	"testing"
)


func TestDatasourceAccountRead(t *testing.T) {
	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	dataSource := dataSourceAccount()
	d := dataSource.Data(nil)
	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_ACCOUNTID"))
	d.Set("vendor_key", "aws")
	c := config.Client()
	dataSourceAccountRead(d, c)
}