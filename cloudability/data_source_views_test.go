package cloudability

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccViewsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccViewsDataSourceConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(
				// resource.TestCheckResourceAttr("data.cloudability_views.test", "id", "testid"),
				),
			},
		},
	})
}

const testAccViewsDataSourceConfig = `
data "cloudability_views" "test" {}
`
