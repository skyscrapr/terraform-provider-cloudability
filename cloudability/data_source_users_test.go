package cloudability

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUsersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccUsersDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloudability_users.test", "id"),
				),
			},
		},
	})
}

const testAccUsersDataSourceConfig = `
data "cloudability_users" "test" {}
`
