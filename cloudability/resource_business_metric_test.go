package cloudability

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBusinessMetricResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccBusinessMetricResourceConfig("A"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloudability_business_metric.test", "id"),
					resource.TestCheckResourceAttr("cloudability_business_metric.test", "name", "Cost (Surcharge) A"),
					resource.TestCheckResourceAttr("cloudability_business_metric.test", "default_value", "METRIC['unblended_cost']"),
					resource.TestCheckResourceAttr("cloudability_business_metric.test", "default_value_expression", "METRIC['unblended_cost']"),
					resource.TestCheckResourceAttr("cloudability_business_metric.test", "pre_match_expression", "DIMENSION['vendor'] == 'Amazon'"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "cloudability_business_metric.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				ImportStateVerifyIgnore: []string{"default_value_expression"},
			},
			// Update and Read testing
			{
				Config: testAccBusinessMetricResourceConfig("B"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cloudability_business_metric.test", "name", "Cost (Surcharge) B"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccBusinessMetricResourceConfig(name string) string {
	return fmt.Sprintf(`
resource "cloudability_business_metric" "test" {
	name = "Cost (Surcharge) %s"
	number_format = "number"
	default_value_expression = "METRIC['unblended_cost']"
	pre_match_expression = "DIMENSION['vendor'] == 'Amazon'"
	statement {
		match_expression = "DIMENSION['vendor'] == 'Amazon' || DIMENSION['vendor'] == 'Azure'"
		value_expression = "METRIC['unblended_cost'] * 1.15"
	}

	// lifecycle {
	// 	ignore_changes = [
	// 		default_value_expression,
	// 	  	default_value
	// 	]
	// }
  
	// name = "Cost (Storage Backup)"
	// // "kind": "BUSINESS_METRIC"
	// number_format = "currency"
	// pre_match_expression = "(DIMENSION['vendor'] == 'amazon' && DIMENSION['usage_family'] == 'storage'"
	// default_value_expression = "METRIC['unblended_cost']"
	// statement {
	// 	match_expression = "DIMENSION['invoice_date'] == '2019-09-01' && DIMENSION['transaction_type'] == 'usage' && DIMENSION['usage_type'] CONTAINS 'snapshot'"
	// 	value_expression = "METRIC['unblended_cost'] * 1.10"
	// }
	// statement {
	// 	match_expression = "DIMENSION['invoice_date'] == '2019-10-01' && DIMENSION['transaction_type'] == 'usage' && DIMENSION['usage_type'] CONTAINS 'snapshot'"
	// 	value_expression = "METRIC['unblended_cost'] * 1.10"
	// }
	// statement {
	// 	match_expression = "DIMENSION['invoice_date'] == '2019-11-01' && DIMENSION['transaction_type'] == 'usage' && DIMENSION['usage_type'] CONTAINS 'snapshot'"
	// 	value_expression = "METRIC['unblended_cost'] * 1.10"
	// }
}
`, name)
}

// func testAccBusinessMetricMultipleConfig() string {
// 	return `
// resource "cloudability_business_mapping" "test1" {
// 	name = "test1"
// 	default_value = "Unknown1"
// 	kind = "BUSINESS_DIMENSION"
// 	statement {
// 		match_expression = "DIMENSION['vendor'] == 'vendor1_1'"
// 		value_expression = "'Vendor1_1'"
// 	}
// 	statement {
// 		match_expression = "DIMENSION['vendor'] == 'vendor1_2'"
// 		value_expression = "'Vendor1_2'"
// 	}
// }

// resource "cloudability_business_mapping" "test2" {
// 	name = "test2"
// 	default_value = "Unknown2"
// 	kind = "BUSINESS_DIMENSION"
// 	statement {
// 		match_expression = "DIMENSION['vendor'] == 'vendor2_1'"
// 		value_expression = "'Vendor2_1'"
// 	}
// 	statement {
// 		match_expression = "DIMENSION['vendor'] == 'vendor2_2'"
// 		value_expression = "'Vendor2_2'"
// 	}
// 	depends_on = [cloudability_business_mapping.test1]
// }

// resource "cloudability_business_mapping" "test3" {
// 	name = "test3"
// 	default_value = "Unknown3"
// 	kind = "BUSINESS_DIMENSION"
// 	statement {
// 		match_expression = "DIMENSION['vendor'] == 'vendor3_1'"
// 		value_expression = "'Vendor3_1'"
// 	}
// 	statement {
// 		match_expression = "DIMENSION['vendor'] == 'vendor3_2'"
// 		value_expression = "'Vendor3_2'"
// 	}
// 	depends_on = [cloudability_business_mapping.test2]
// }
// `
// }
