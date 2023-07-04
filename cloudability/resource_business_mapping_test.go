package cloudability

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBusinessMappingResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccBusinessMappingResourceConfig("1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloudability_business_mapping.test", "id"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test", "name", "test_1"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test", "kind", "BUSINESS_DIMENSION"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test", "default_value", "Unknown"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "cloudability_business_mapping.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				// ImportStateVerifyIgnore: []string{"attrib"},
			},
			// Update and Read testing
			{
				Config: testAccBusinessMappingResourceConfig("2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cloudability_business_mapping.test", "name", "test_2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccBusinessMappingResourceConfig(name string) string {
	return fmt.Sprintf(`	
resource "cloudability_business_mapping" "test" {
	name = "test_%s"
	default_value = "Unknown"
	kind = "BUSINESS_DIMENSION"
	statement {
		match_expression = "DIMENSION['vendor'] == 'vendor_1'"
		value_expression = "'Vendor1'"
	}
	statement {
		match_expression = "DIMENSION['vendor'] == 'vendor_2'"
		value_expression = "'Vendor2'"
	}
}
`, name)
}
