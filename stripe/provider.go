package stripe

import (
	"context"
	"terraform-provider-stripe/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("STRIPE_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"stripe_user": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"stripe_user": dataSourceUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	var diags diag.Diagnostics
	return client.NewClient(token), diags
}
