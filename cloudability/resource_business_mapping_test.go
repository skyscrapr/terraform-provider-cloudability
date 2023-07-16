package cloudability

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBusinessMappingResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				// VersionConstraint: "...", // last version of old schema version
				Source:            "registry.terraform.io/hashicorp/time",
			},
		},
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
			// Multiple business mappings at once
			{
				Config: testAccBusinessMappingMultipleConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloudability_business_mapping.test1", "id"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test1", "name", "test__1"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test1", "kind", "BUSINESS_DIMENSION"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test1", "default_value", "Unknown1"),
					resource.TestCheckResourceAttrSet("cloudability_business_mapping.test2", "id"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test2", "name", "test__2"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test2", "kind", "BUSINESS_DIMENSION"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test2", "default_value", "Unknown2"),
					resource.TestCheckResourceAttrSet("cloudability_business_mapping.test3", "id"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test3", "name", "test__3"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test3", "kind", "BUSINESS_DIMENSION"),
					resource.TestCheckResourceAttr("cloudability_business_mapping.test3", "default_value", "Unknown3"),
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

func testAccBusinessMappingMultipleConfig() string {
	return `
provider "time" {
	version = "~> 0.7"
}

resource "cloudability_business_mapping" "test1" {
	name = "test__1"
	default_value = "Unknown1"
	kind = "BUSINESS_DIMENSION"
	statement {
		match_expression = "DIMENSION['vendor'] == 'vendor1_1'"
		value_expression = "'Vendor1_1'"
	}
	statement {
		match_expression = "DIMENSION['vendor'] == 'vendor1_2'"
		value_expression = "'Vendor1_2'"
	}
}

resource "time_sleep" "wait_test1" {
	depends_on = [cloudability_business_mapping.test1]
	create_duration = "2s"
}

resource "cloudability_business_mapping" "test2" {
	name = "test__2"
	default_value = "Unknown2"
	kind = "BUSINESS_DIMENSION"
	statement {
		match_expression = "DIMENSION['vendor'] == 'vendor2_1'"
		value_expression = "'Vendor2_1'"
	}
	statement {
		match_expression = "DIMENSION['vendor'] == 'vendor2_2'"
		value_expression = "'Vendor2_2'"
	}
	depends_on = [time_sleep.wait_test1]
}

resource "time_sleep" "wait_test2" {
	depends_on = [cloudability_business_mapping.test2]
	create_duration = "2s"
}

resource "cloudability_business_mapping" "test3" {
	name = "test__3"
	default_value = "Unknown3"
	kind = "BUSINESS_DIMENSION"
	statement {
		match_expression = "DIMENSION['vendor'] == 'vendor3_1'"
		value_expression = "'Vendor3_1'"
	}
	statement {
		match_expression = "DIMENSION['vendor'] == 'vendor3_2'"
		value_expression = "'Vendor3_2'"
	}
	depends_on = [time_sleep.wait_test2]
}
`
}
