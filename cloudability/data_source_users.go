package cloudability

import (
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUsersRead,
		Schema: map[string]*schema.Schema{
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the user",
						},
						"frontdoor_user_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Frontdoor user id",
						},
						"frontdoor_login": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Frontdoor login",
						},
						"email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The email address of the user",
						},
						"full_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The full name of the user",
						},
						"shared_dimension_filter_set_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Description: "Array of filter sets ids available to the user",
						},
						"default_dimension_filter_set_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Filter set id used by default for the user",
						},
					},
				},
			},
		},
	}
}

func dataSourceUsersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	users, err := client.Users().GetUsers()
	if err != nil {
		return err
	}

	if users != nil {
		var data = make([]map[string]interface{}, len(users))
		for i, user := range users {
			m := make(map[string]interface{})
			m["id"] = user.ID
			m["frontdoor_user_id"] = user.FrontdoorUserId
			m["frontdoor_login"] = user.FrontdoorLogin
			m["email"] = user.Email
			m["full_name"] = user.FullName
			m["shared_dimension_filter_set_ids"] = user.SharedDimensionFilterSetIDs
			m["default_dimension_filter_set_id"] = user.DefaultDimensionFilterSetID
			data[i] = m
		}
		d.Set("users", data)
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	}

	return nil
}
