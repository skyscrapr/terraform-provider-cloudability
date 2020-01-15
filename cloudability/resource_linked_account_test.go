package cloudability

// import (
// 	"fmt"
// 	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
// 	"github.com/hashicorp/terraform-plugin-sdk/terraform"
// 	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
// 	"os"
// 	"testing"
// )

// func TestAccLinkedAccount(t *testing.T) {
// 	var account *cloudability.Account
// 	accountId := os.Getenv("CLOUDABILITY_ACCOUNTID")
// 	payer_accountId := os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID")
// 	cloudability_apikey := os.Getenv("CLOUDABILITY_APIKEY")

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckLinkedAccountDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccLinkedAccount(cloudability_apikey, accountId, payer_accountId),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckLinkedAccountResourceExists("cloudability_linked_account.test", account),
// 					// TODO: Complete this
// 					// testAccCheckExampleWidgetValues(widget, rName),
// 					resource.TestCheckResourceAttr("cloudability_linked_account.test", "vendor_account_id", accountId),
// 					resource.TestCheckResourceAttr("cloudability_linked_account.test", "vendor_account_name", "CloudabilityLinkedTest"),
// 				),
// 			},
// 		},
// 	})
// }

// // testAccCheckLinkedAccountDestroy verifies the LinkedAccount resource has been destroyed.
// // Linked accounts will still remain in the account list. But verification and authorization fields will be nil.
// // Linked accounts will be fully destroyed when the master account is destroyed.
// func testAccCheckLinkedAccountDestroy(s *terraform.State) error {
// 	client := testAccProvider.Meta().(*cloudability.Client)
// 	// loop through the resources in state, verifying each resource is destroyed
// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "cloudability_linked_account" {
// 			continue
// 		}
// 		account, err := client.Vendors().GetAccount(rs.Primary.Attributes["vendor_key"], rs.Primary.ID)
// 		if err != nil {
// 			return err
// 		} else if account != nil {
// 			if account.Verification != nil || account.Authorization != nil {
// 				return fmt.Errorf("Linked Account (%s) still exists.", rs.Primary.ID)
// 			}
// 		}
// 	}
// 	return nil
// }

// // testAccCheckLinkedAccountResourceExists uses the Cloudability SDK to retrieve the Account.
// func testAccCheckLinkedAccountResourceExists(resourceName string, account *cloudability.Account) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		// retrieve the resource by name from state
// 		rs, ok := s.RootModule().Resources[resourceName]
// 		if !ok {
// 			return fmt.Errorf("Not found: %s", resourceName)
// 		}
// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("LinkedAccount ID is not set")
// 		}
// 		// retrieve the client from the test provider
// 		client := testAccProvider.Meta().(*cloudability.Client)
// 		var err error
// 		account, err = client.Vendors().GetAccount(rs.Primary.Attributes["vendor_key"], rs.Primary.ID)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// }

// // testAccLinkedAccount returns a configuration for an LinkedAccount
// func testAccLinkedAccount(cloudability_apikey string, accountId string, payer_accountId string) string {
// 	return fmt.Sprintf(`
// provider cloudability {
// 	apikey = "%s"
// }

// resource "cloudability_master_account" "aws_payer_account" {
//     vendor_account_id = "%s"
//     bucket_name = "cloudability-949237620074-20200113000147075900000001"
//     report_name = "Cloudability"
//     report_prefix = "CostAndUsageReports"
// }

// module "cloudability_payer_aws_iam_role" {
// 	source = "github.com/skyscrapr/terrafrom-cloudability-modules/cloudability-aws-iam-role"

// 	aws_region = var.aws_region
// 	cloudability_account_external_id = cloudability_master_account.aws_payer_account.external_id
// }

// module "cloudability_aws_iam_role" {
// 	source = "github.com/skyscrapr/terrafrom-cloudability-modules/cloudability-aws-iam-role"

// 	aws_region = var.aws_region
// 	cloudability_account_external_id = cloudability_master_account.aws_payer_account.external_id
// }

// data "cloudability_account_verification" "aws_payer_account" {
//     vendor_account_id = "%s"
//     depends_on = [
//         cloudability_master_account.aws_payer_account,
//     ]
// }

// resource "cloudability_linked_account" "test" {
// 	vendor_account_id = "%s"
// }`, cloudability_apikey, payer_accountId, payer_accountId, accountId)
// }

// func TestResourceLinkedAccountRead(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceLinkedAccount()
// 	d := resource.Data(nil)
// 	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_ACCOUNTID"))
// 	d.Set("vendor_key", "aws")
// 	c := config.Client()
// 	err := resourceLinkedAccountRead(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }

// func TestResourceLinkedAccountCreate(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceLinkedAccount()
// 	d := resource.Data(nil)
// 	accountId := os.Getenv("CLOUDABILITY_ACCOUNTID")
// 	d.Set("vendor_account_id", accountId)
// 	d.Set("vendor_key", "aws")
// 	d.Set("type", "aws_role")

// 	c := config.Client()
// 	err := resourceLinkedAccountCreate(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }

// func TestResourceLinkedAccountDelete(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceLinkedAccount()
// 	d := resource.Data(nil)
// 	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_ACCOUNTID"))
// 	d.Set("vendor_key", "aws")
// 	c := config.Client()
// 	err := resourceLinkedAccountDelete(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }
