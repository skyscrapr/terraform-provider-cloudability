package cloudability

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
        return &schema.Provider{
                Schema: map[string]*schema.Schema{
                        "apikey": {
                                Type: schema.TypeString,
                                Required: true,
                                Description: "The apikey for API operations",
                                DefaultFunc: schema.EnvDefaultFunc("APIKEY", nil),
                        },
                },              
                ResourcesMap: map[string]*schema.Resource{
                        // "cloudability_aws_credential": resourceAWSCredential(),
                        "cloudability_business_mapping": resourceBusinessMapping(),
                        "cloudability_user": resourceUser(),
                        // "cloudability_view": resourceView(),
                },
                ConfigureFunc: providerConfigure,
        }
}


func providerConfigure(d *schema.ResourceData) (interface{}, error) {
        c := NewConfig(d)
        return c.Client(), nil
}
