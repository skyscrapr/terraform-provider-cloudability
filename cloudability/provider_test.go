package cloudability

import (
	"os"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testProviders = map[string]terraform.ResourceProvider{
	"cloudability": Provider(),
}

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"cloudability": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("CLOUDABILITY_APIKEY") != "" {
		return
	}
	// for _, s := range [...]string("CLOUDABILITY_APIKEY") {
	// 	if os.Getenv(s) == "" {
	// 	}
	// }
	t.Fatal("CLOUDABILITY_APIKEY env var is required for tests")
}