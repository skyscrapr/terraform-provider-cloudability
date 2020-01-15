package cloudability

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testAccAWSProvider terraform.ResourceProvider

func init() {
	testAccProvider = Provider()
	testAccAWSProvider = aws.Provider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"cloudability": testAccProvider,
		"aws":          testAccAWSProvider,
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
