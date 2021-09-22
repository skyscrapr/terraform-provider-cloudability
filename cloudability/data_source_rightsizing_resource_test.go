package cloudability

import (
	// 	"os"
	"testing"
	// "fmt"
	// 	// 	"regexp"
	// 	// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	// "github.com/skyscrapr/cloudability-sdk-go/cloudability"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// func TestCloudabilityDataSourceRightsizingResource(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		IsUnitTest: true,
// 		Providers:    testProviders,
// 		CheckDestroy: testAccCheckCloudabilityDataSourceAccountVerificationDestroy,

// 	})
// }

func TestDatasourceRightsizingResourceReadExistingRightsize(t *testing.T) {
	vendor := "aws"
	service := "ec2"
	resourceIdentifier := "i-abcde12345"
	resourceJSON := []byte(`{
		"result": {
			"service": "ec2",
			"resourceIdentifier": "i-abcde12345",
			"recommendations": [
				{
					"action": "Rightsize",
					"nodeType": "x1e.xlarge",
					"risk": 0
				}
			]
		}
	}`)
	dataSource := dataSourceRightsizingResource()
	d := dataSource.Data(nil)
	d.Set("vendor", vendor)
	d.Set("service", service)
	d.Set("resource_identifier", resourceIdentifier)

	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(resourceJSON))
	}))
	defer testServer.Close()
	config := Config{
		APIKey: "MOCK",
	}
	client := config.Client()
	client.V3BaseURL, _ = url.Parse(testServer.URL)

	err := dataSourceRightsizingResourceRead(d, client)
	if err != nil {
		t.Errorf(err.Error())
	}
}
