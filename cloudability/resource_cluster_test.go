package cloudability

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestClusterResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Create and Read testing with only kubernetes_version
			{
				Config: testAccClusterResourceConfig("cluster_1", "1.22", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloudability_cluster.test", "id"),
					resource.TestCheckResourceAttr("cloudability_cluster.test", "cluster_name", "cluster_1"),
					resource.TestCheckResourceAttr("cloudability_cluster.test", "kubernetes_version", "1.22"),
					resource.TestCheckResourceAttr("cloudability_cluster.test", "cluster_version", ""),
				),
			},
			// Update and Read testing with only cluster_version
			{
				Config: testAccClusterResourceConfig("cluster_2", "", "2.0"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cloudability_cluster.test", "cluster_name", "cluster_2"),
					resource.TestCheckResourceAttr("cloudability_cluster.test", "kubernetes_version", ""),
					resource.TestCheckResourceAttr("cloudability_cluster.test", "cluster_version", "2.0"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "cloudability_cluster.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccClusterResourceConfig(clusterName string, kubernetesVersion string, clusterVersion string) string {
	return fmt.Sprintf(`
resource "cloudability_cluster" "test" {
  cluster_name = "%s"
  %s
  %s
}
`, clusterName,
		ifString(kubernetesVersion != "", fmt.Sprintf(`kubernetes_version = "%s"`, kubernetesVersion)),
		ifString(clusterVersion != "", fmt.Sprintf(`cluster_version = "%s"`, clusterVersion)),
	)
}

func ifString(condition bool, trueVal string) string {
	if condition {
		return trueVal
	}
	return ""
}
