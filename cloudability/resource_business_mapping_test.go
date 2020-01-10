package cloudability

import (
	"testing"
	"os"
)


func TestResourceCloudabilityBusinessMappingRead(t *testing.T) {
	apikey, _ := os.LookupEnv("CLOUDABILTIY_APIKEY")
	config := Config{
		ApiKey: apikey,
	}
	resouce := resourceBusinessMapping()
	d := resouce.Data(nil)
	d.SetId("1")
	c := config.Client()
	resourceBusinessMappingRead(d, c)
}