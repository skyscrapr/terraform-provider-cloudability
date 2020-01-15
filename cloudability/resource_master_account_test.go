package cloudability

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
	"os"
	"testing"
)

func TestAccMasterAccount(t *testing.T) {
	var account *cloudability.Account
	accountId := os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID")
	cloudabilityApikey := os.Getenv("CLOUDABILITY_APIKEY")
	awsRegion := os.Getenv("AWS_DEFAULT_REGION")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMasterAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMasterAccount(cloudabilityApikey, accountId, awsRegion),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMasterAccountResourceExists("cloudability_master_account.test", account),
					// TODO: Complete this
					// testAccCheckExampleWidgetValues(widget, rName),
					resource.TestCheckResourceAttr("cloudability_master_account.test", "vendor_account_id", accountId),
					resource.TestCheckResourceAttr("cloudability_master_account.test", "vendor_account_name", "9492-3762-0074"),
				),
			},
		},
	})
}

// testAccCheckMasterAccountDestroy verifies the MasterAccount resource has been destroyed.
func testAccCheckMasterAccountDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*cloudability.Client)
	// loop through the resources in state, verifying each resource is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudability_master_account" {
			continue
		}
		account, err := client.Vendors().GetAccount(rs.Primary.Attributes["vendor_key"], rs.Primary.ID)
		if err != nil {
			return err
		} else if account != nil {
			if account.Verification != nil || account.Authorization != nil {
				return fmt.Errorf("Master Account (%s) still exists.", rs.Primary.ID)
			}
		}
	}
	return nil
}

// testAccCheckMasterAccountResourceExists uses the Cloudability SDK to retrieve the Account.
func testAccCheckMasterAccountResourceExists(resourceName string, account *cloudability.Account) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("MasterAccount ID is not set")
		}
		// retrieve the client from the test provider
		client := testAccProvider.Meta().(*cloudability.Client)
		var err error
		account, err = client.Vendors().GetAccount(rs.Primary.Attributes["vendor_key"], rs.Primary.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

// source = "github.com/skyscrapr/terraform-cloudability-modules/cloudability-master-account"

// testAccMasterAccount returns a configuration for an MasterAccount
func testAccMasterAccount(cloudabilityApikey string, accountId string, awsRegion string) string {
	return fmt.Sprintf(`
module "clouability_master_account" {
	source = "github.com/skyscrapr/terraform-cloudability-modules//cloudability-master-account"
	aws_payer_account_id = "%s"
	aws_region = "%s"
	cloudability_apikey = "%s"
}`, accountId, awsRegion, cloudabilityApikey)
}

// func TestResourceMasterAccountRead(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceMasterAccount()
// 	d := resource.Data(nil)
// 	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID"))
// 	d.Set("vendor_key", "aws")
// 	c := config.Client()
// 	err := resourceMasterAccountRead(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }

// func TestResourceMasterAccountCreate(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceMasterAccount()
// 	d := resource.Data(nil)
// 	accountId := os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID")
// 	d.Set("vendor_account_id", accountId)
// 	d.Set("vendor_key", "aws")
// 	d.Set("type", "aws_role")
// 	d.Set("bucket_name", fmt.Sprintf("cloudability-%s", accountId))
// 	d.Set("report_name", "Cloudability")
// 	d.Set("report_prefix", "CostAndUsageReports")

// 	c := config.Client()
// 	err := resourceMasterAccountCreate(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }

// func TestResourceMasterAccountDelete(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceMasterAccount()
// 	d := resource.Data(nil)
// 	d.Set("vendor_account_id", os.Getenv("CLOUDABILITY_MASTER_ACCOUNTID"))
// 	d.Set("vendor_key", "aws")
// 	c := config.Client()
// 	err := resourceMasterAccountDelete(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }
