package stripe

import (
	"context"

	"terraform-provider-stripe/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"object": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"balance": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := i.(*client.Client)
	Email := data.Get("email").(string)
	user, err := apiClient.GetItem(Email)
	if err != nil {
		return diag.FromErr(err)
	}
	data.Set("object", user.Object)
	data.Set("balance", user.Balance)
	data.Set("created", user.Created)
	data.Set("email", user.Email)
	data.Set("name", user.Name)
	data.Set("description", user.Description)
	data.Set("phone", user.Phone)
	data.SetId(user.Email)
	return diags
}
