package cloudability

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/skyscrapr/cloudability-sdk-go/cloudability"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read: resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"email": {
				Type: schema.TypeString,
				Required: true,
				Description: "The email address of the user",
			},
			"full_name": {
				Type: schema.TypeString,
				Required: true,
				Description: "The full name of the user",
			},
			"role": {
				Type: schema.TypeString,
				Optional: true,
				Default: "User",
				Description: "Role assigned to the user: [User|Administrator]",
			},
			"restricted": {
				Type: schema.TypeBool,
				Optional: true,
				Default: false,
				Description: "True if the user is allowed to have no filter set applied, false if they must always have a filter set applied",
			},
			"shared_dimension_filter_set_ids": {
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{Type: schema.TypeInt},
				Description: "Array of filter sets ids available to the user",
			},
			"default_dimension_filter_set_id": {
				Type: schema.TypeInt,
				Optional: true,
				Default: nil,
				Description: "Filter set id used by default for the user",
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)	

	user := &cloudability.User{
		Email: d.Get("email").(string),
		FullName: d.Get("full_name").(string),
		Role: d.Get("role").(string),
		Restricted: d.Get("restricted").(bool),
		// TODO: Fix this
		// SharedDimensionFilterSetIds: d.Get("shared_dimension_filter_set_ids"),
		DefaultDimensionFilterId: d.Get("default_dimension_filter_set_id").(int),
	}
	err := client.Users().NewUser(user)
	if err != nil {
		return err
	}
	return resourceUserRead(d, meta)
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	user, err := client.Users().GetUser(id)
	if err != nil {
		return err
	}

	if user != nil {
		d.Set("email", user.Email)
		d.Set("full_name", user.FullName)
		d.Set("role", user.Role)
		d.Set("restricted", user.Restricted)
		d.Set("shared_dimension_filter_set_ids", user.SharedDimensionFilterSetIds)
		d.Set("default_dimension_filter_set_id", user.DefaultDimensionFilterId)
		d.SetId(strconv.Itoa(user.Id))
	}
	return nil
}
 
func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil
	}
	user := &cloudability.User{
		Id: id,
		Email: d.Get("email").(string),
		FullName: d.Get("full_name").(string),
		Role: d.Get("role").(string),
		Restricted: d.Get("restricted").(bool),
		// TODO: Fix this
		// SharedDimensionFilterSetIds: d.Get("shared_dimension_filter_set_ids"),
		DefaultDimensionFilterId: d.Get("default_dimension_filter_set_id").(int),
	}
	err = client.Users().UpdateUser(user)
	if err != nil {
		return err
	}
	return resourceUserRead(d, meta)
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudability.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil
	}
	err = client.Users().DeleteUser(id)
	return err
}
