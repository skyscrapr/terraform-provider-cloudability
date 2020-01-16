package cloudability

// import (
// 	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
// 	"github.com/hashicorp/terraform-plugin-sdk/terraform"
// 	"github.com/terraform-providers/terraform-provider-aws/aws"
// 	"os"
// 	"testing"
// )

// var testAccProviders map[string]terraform.ResourceProvider
// var testAccProvider, testAccAWSProvider *schema.Provider

// func init() {
// 	testAccProvider = Provider()
// 	testAccAWSProvider = aws.Provider().(*schema.Provider)
// 	// testAccAWSCURProvider = aws.Provider().(*schema.Provider)
// 	// testAccAWSCURProvider.Schema["region"].Elem = "us-east-1"
// 	// // testAccAWSCURProvider.Schema["alias"].Elem = "us-east-1"
// 	testAccProviders = map[string]terraform.ResourceProvider{
// 		"cloudability": testAccProvider,
// 		"aws":          testAccAWSProvider,
// 		// "aws.us-east-1": testAccAWSCURProvider,
// 	}
// }

// func testAccPreCheck(t *testing.T) {
// 	if os.Getenv("CLOUDABILITY_APIKEY") == "" {
// 		t.Fatal("CLOUDABILITY_APIKEY env var is required for tests")
// 	}
// 	if os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID") == "" {
// 		t.Fatal("CLOUDABILITY_MASTER_ACCOUNTID env var is required for tests")
// 	}
// 	if os.Getenv("CLOUDABILITY_ACCOUNTID") == "" {
// 		t.Fatal("CLOUDABILITY_ACCOUNTID env var is required for tests")
// 	}

// 	// for _, s := range [...]string("CLOUDABILITY_APIKEY") {
// 	// 	if os.Getenv(s) == "" {
// 	// 	}
// 	// }
// 	return
// }
