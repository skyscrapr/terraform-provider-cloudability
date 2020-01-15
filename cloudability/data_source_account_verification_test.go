package cloudability

// import (
// 	"os"
// 	"testing"
// 	// 	"fmt"
// 	// 	"regexp"
// 	// 	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
// 	// 	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
//)

// func TestCloudabilityDataSourceAccountVerification(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		IsUnitTest: true,
// 		Providers:    testProviders,
// 		CheckDestroy: testAccCheckCloudabilityDataSourceAccountVerificationDestroy,

// 	})
// }

// // testAccCheckCloudabilityDataSourceAccountVerificationDestroy verifies the
// // datasourceAccountVerification instnace has been destroyed
// func testAccCheckCloudabilityDataSourceAccountVerificationDestroy(s *terraform.State) error {
// 	// retrieve the connection established in Provider configuration
// 	conn := testAccProvider.Meta().(*ExampleClient)

// 	// loop through the resources in state, verifying each widget
// 	// is destroyed
// 	for _, rs := range s.RootModule().Resources {
// 	  if rs.Type != "example_widget" {
// 		continue
// 	  }

// 	  // Retrieve our widget by referencing it's state ID for API lookup
// 	  request := &example.DescribeWidgets{
// 		IDs: []string{rs.Primary.ID},
// 	  }

// 	  response, err := conn.DescribeWidgets(request)
// 	  if err == nil {
// 		if len(response.Widgets) > 0 && *response.Widgets[0].ID == rs.Primary.ID {
// 		  return fmt.Errorf("Widget (%s) still exists.", rs.Primary.ID)
// 		}

// 		return nil
// 	  }

// 	  // If the error is equivalent to 404 not found, the widget is destroyed.
// 	  // Otherwise return the error
// 	  if !strings.Contains(err.Error(), "Widget not found") {
// 		return err
// 	  }
// 	}

// 	return nil
//   }

// func TestAccCloudabilityDataSourceAccountVerification_basic(t *testing.T) {
// 	accountId := os.Getenv("CLOUDABILITY_ACCOUNTID")

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:  func() { testAccPreCheck(t) },
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCloudabilityDataSourceAccountVerificationConfig(accountId),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttrSet("data.cloudability_account_verification.test", "user_id"),
// 					resource.TestCheckResourceAttr("data.cloudability_account_verification.test", "vendor_account_id", accountId),
// 					resource.TestCheckResourceAttr("data.cloudability_account_verification.test", "vendor_key", "aws"),
// 					resource.TestCheckResourceAttr("data.cloudability_account_verification.test", "state", "verified"),
// 					resource.TestMatchResourceAttr("data.cloudability_account_verification.test", "retry_count", 20),
// 					resource.TestMatchResourceAttr("data.cloudability_account_verification.test", "retry_wait", 5),
// 					resource.TestMatchResourceAttr("data.cloudability_account_verification.test", "external_id", regexp.MustCompile(`([a-f0-9]+\-)+[a-f0-9]+`)),
// 					// last_verification_attempted_at
// 					// message
// 				),
// 			},
// 		},
// 	})
// }

// func testAccCloudabilityDataSourceAccountVerificationConfig(accountId string) string {
// 	return fmt.Sprintf(`
// resource "cloudability_linked_account" "account" {
//   vendor_account_id = "%s"
// }

// data "cloudability_account_verification" "test" {
// 	vendor_account_id = cloudability_linked_account.account.vendor_account_id
// }
// `, accountId)
// }

// func TestDatasourceAccountVerificationRead(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	dataSource := dataSourceAccountVerification()
// 	d := dataSource.Data(nil)

// 	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID"))
// 	// d.Set("vendor_key", "aws")
// 	// d.Set("retry_count", 20)
// 	// d.Set("retry_wait", 5)
// 	c := config.Client()
// 	err := dataSourceAccountVerificationRead(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }
