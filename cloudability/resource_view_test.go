package cloudability

// import (
// 	"os"
// 	"testing"
// )

// func TestResourceViewRead(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceView()
// 	d := resource.Data(nil)
// 	d.SetId("48190")
// 	c := config.Client()
// 	err := resourceViewRead(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }

// func TestResourceViewCreate(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceView()
// 	d := resource.Data(nil)
// 	d.Set("title", "test-view")
// 	d.Set("shared_with_organization", true)
// 	filters := []map[string]interface{}{
// 		{
// 			"field":      "category1",
// 			"comparator": "==",
// 			"value":      "Amazon",
// 		},
// 	}
// 	d.Set("filter", filters)
// 	c := config.Client()
// 	err := resourceViewCreate(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }

// func TestResourceViewDelete(t *testing.T) {
// 	apikey := os.Getenv("CLOUDABILITY_APIKEY")
// 	config := Config{
// 		ApiKey: apikey,
// 	}
// 	resource := resourceMasterAccount()
// 	d := resource.Data(nil)
// 	d.SetId("48190")
// 	c := config.Client()
// 	err := resourceViewDelete(d, c)
// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}
// }
