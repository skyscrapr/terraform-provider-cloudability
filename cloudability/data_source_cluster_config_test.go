package cloudability

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestClusterConfigDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testClusterConfigDataSource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloudability_cluster_config.test", "cluster_id"),
				),
			},
		},
	})
}

const testClusterConfigDataSource = `
data "cloudability_cluster_config" "test" {
	cluster_id = "1"
}
`
