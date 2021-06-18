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
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_name": &schema.Schema{
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
	data.Set("email", user.Individual.Email)
	data.Set("first_name", user.Individual.FirstName)
	data.Set("last_name", user.Individual.LastName)
	data.SetId(user.Email)
	return diags
}
