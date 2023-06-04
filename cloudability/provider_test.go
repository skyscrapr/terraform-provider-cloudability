package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

var testAccProvider *schema.Provider
var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProvider = Provider()
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"cloudability": func() (*schema.Provider, error) { return testAccProvider, nil },
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("CLOUDABILITY_APIKEY") == "" {
		t.Fatal("CLOUDABILITY_APIKEY env var is required for tests")
	}
	// if os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID") == "" {
	// 	t.Fatal("CLOUDABILITY_MASTER_ACCOUNTID env var is required for tests")
	// }
	// if os.Getenv("CLOUDABILITY_ACCOUNTID") == "" {
	// 	t.Fatal("CLOUDABILITY_ACCOUNTID env var is required for tests")
	// }

	// for _, s := range [...]string("CLOUDABILITY_APIKEY") {
	// 	if os.Getenv(s) == "" {
	// 	}
	// }
	return
}
