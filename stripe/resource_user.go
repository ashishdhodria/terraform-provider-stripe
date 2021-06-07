package stripe

import (
	"context"
	"terraform-provider-stripe/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stripe/stripe-go/v72"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func setData(user *stripe.Account, data *schema.ResourceData) {
	data.Set("id", user.Email)
	data.Set("email", user.Individual.Email)
	data.Set("first_name", user.Individual.FirstName)
	data.Set("last_name", user.Individual.LastName)
	data.SetId(user.Email)
}

func resourceUserCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := i.(*client.Client)
	params := &stripe.AccountParams{
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
		Country:      stripe.String("IN"),
		Email:        stripe.String(data.Get("email").(string)),
		Type:         stripe.String("custom"),
		BusinessType: stripe.String("individual"),
		Individual: &stripe.PersonParams{
			Email:     stripe.String(data.Get("email").(string)),
			FirstName: stripe.String(data.Get("first_name").(string)),
			LastName:  stripe.String(data.Get("last_name").(string)),
		},
	}
	user, err := apiClient.NewItem(params)
	if err != nil {
		return diag.FromErr(err)
	}
	setData(user, data)
	return diags
}

func resourceUserRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := i.(*client.Client)
	Email := data.Id()
	user, err := apiClient.GetItem(Email)
	if err != nil {
		return diag.FromErr(err)
	}
	setData(user, data)
	return diags
}

func resourceUserUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := i.(*client.Client)
	Email := data.Id()
	params := &stripe.AccountParams{
		Individual: &stripe.PersonParams{
			FirstName: stripe.String(data.Get("first_name").(string)),
			LastName:  stripe.String(data.Get("last_name").(string)),
		},
	}
	if data.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "User not allowed to change email",
			Detail:   "User not allowed to change email",
		})
		return diags
	}
	if data.HasChanges("first_name") || data.HasChange("last_name") {
		user, err := apiClient.UpdateItem(params, Email)
		if err != nil {
			return diag.FromErr(err)
		}
		setData(user, data)
	}
	return diags
}

func resourceUserDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := i.(*client.Client)
	Email := data.Id()
	user, err := apiClient.DeleteItem(Email)
	if err != nil {
		return diag.FromErr(err)
	}
	if user.ID == "" {
		data.SetId("")
	}
	return diags
}
