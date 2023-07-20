package cloudability

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccViewDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccViewDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloudability_view.test", "id"),
					resource.TestCheckResourceAttr("data.cloudability_view.test", "title", "TestTitle"),
				),
			},
		},
	})
}

const testAccViewDataSourceConfig = `
resource "cloudability_business_mapping" "test" {
    name = "Test Business Mapping"
    default_value = "TestDefaultValue"
    kind = "BUSINESS_DIMENSION"
    statement {
        match_expression = "DIMENSION['vendor'] == 'Amazon'"
        value_expression = "'Amazon'"
    }
}

resource "cloudability_view" "test" {
    title = "TestTitle"
    filter {
        field = "category${cloudability_business_mapping.test.index}"
		comparator = "=="
		value = "test"
    }
}

data "cloudability_view" "test" {
	id = cloudability_view.test.id
}
`
