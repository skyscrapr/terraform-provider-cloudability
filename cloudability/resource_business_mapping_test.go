package cloudability

// import (
// 	"os"
// 	"testing"
// )

// func TestResourceCloudabilityBusinessMappingRead(t *testing.T) {
// 	apikey, _ := os.LookupEnv("CLOUDABILTIY_APIKEY")
// 	config := Config{
// 		APIKey: apikey,
// 	}
// 	resouce := resourceBusinessMapping()
// 	d := resouce.Data(nil)
// 	d.SetId("1")
// 	c := config.Client()
// 	resourceBusinessMappingRead(d, c)
// }

// func TestResourceCloudabilityBusinessMappingCreate(t *testing.T) {
// 	apikey, _ := os.LookupEnv("CLOUDABILTIY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceBusinessMapping()
// 	d := resource.Data(nil)
// 	d.Set("name", "Cloud Service Provider")
// 	d.Set("kind", "BUSINESS_DIMENSION")
// 	d.Set("default_value", "Unknown Cloud Service Provider")
// 	statements := []map[string]interface{}{
// 		{
// 			"match_expression": "DIMENSION['vendor'] == 'Amazon'",
// 			"value_expression": "'Amazon'",
// 		},
// 		{
// 			"match_expression": "DIMENSION['vendor'] == 'Azure'",
// 			"value_expression": "'Azure'",
// 		},
// 	}
// 	d.Set("statement", statements)
// 	c := config.Client()
// 	err := resourceBusinessMappingCreate(d, c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// }

// func TestResourceCloudabilityBusinessMappingDelete(t *testing.T) {
// 	apikey, _ := os.LookupEnv("CLOUDABILTIY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceBusinessMapping()
// 	d := resource.Data(nil)
// 	d.SetId("1")
// 	c := config.Client()
// 	err := resourceBusinessMappingDelete(d, c)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// }
