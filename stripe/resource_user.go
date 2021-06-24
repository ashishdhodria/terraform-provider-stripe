package stripe

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"terraform-provider-stripe/client"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stripe/stripe-go/v72"
)

func validateEmail(v interface{}, k string) (warns []string, errs []error) {
	value := v.(string)
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !(emailRegex.MatchString(value)) {
		errs = append(errs, fmt.Errorf("Expected EmailId is not valid %s", k))
		return warns, errs
	}
	return
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateEmail,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invoice_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_invoice_sequence": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceUserImporter,
		},
	}
}

func setData(user *stripe.Customer, data *schema.ResourceData) {
	data.Set("object", user.Object)
	data.Set("balance", user.Balance)
	data.Set("created", user.Created)
	data.Set("email", user.Email)
	data.Set("name", user.Name)
	data.Set("description", user.Description)
	data.Set("phone", user.Phone)
	data.Set("invoice_prefix", user.InvoicePrefix)
	data.Set("next_invoice_sequence", user.NextInvoiceSequence)
	data.SetId(user.Email)
}

func resourceUserCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := i.(*client.Client)
	params := &stripe.CustomerParams{
		Email:       stripe.String(data.Get("email").(string)),
		Name:        stripe.String(data.Get("name").(string)),
		Description: stripe.String(data.Get("description").(string)),
		Phone:       stripe.String(data.Get("phone").(string)),
	}

	var err error
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		user, err := apiClient.NewItem(params)
		if err != nil {
			if apiClient.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		setData(user, data)
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		return diag.FromErr(retryErr)
	}
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId(*params.Email)
	return diags
}

func resourceUserRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := i.(*client.Client)
	Email := data.Id()
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		user, err := apiClient.GetItem(Email)
		if err != nil {
			if apiClient.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		setData(user, data)
		return nil
	})
	if retryErr != nil {
		if strings.Contains(retryErr.Error(), "user does not exist") {
			data.SetId("")
			return diags
		}
		return diag.FromErr(retryErr)
	}
	return diags
}

func resourceUserUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := i.(*client.Client)
	Email := data.Id()
	params := &stripe.CustomerParams{
		Name:        stripe.String(data.Get("name").(string)),
		Description: stripe.String(data.Get("description").(string)),
		Phone:       stripe.String(data.Get("phone").(string)),
	}
	if data.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "User not allowed to change email",
			Detail:   "User not allowed to change email",
		})
		return diags
	}
	if data.HasChanges("name") || data.HasChanges("description") || data.HasChanges("phone") {
		var err error
		retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
			user, err := apiClient.UpdateItem(params, Email)
			if err != nil {
				if apiClient.IsRetry(err) {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			setData(user, data)
			return nil
		})
		if retryErr != nil {
			time.Sleep(2 * time.Second)
			return diag.FromErr(retryErr)
		}
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceUserDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := i.(*client.Client)
	Email := data.Id()
	var err error
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		_, err := apiClient.DeleteItem(Email)
		if err != nil {
			if apiClient.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		return diag.FromErr(retryErr)
	}
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId("")
	return diags
}

func resourceUserImporter(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	apiClient := i.(*client.Client)
	Email := data.Id()
	user, err := apiClient.GetItem(Email)
	if err != nil {
		return nil, fmt.Errorf("error fetching user. Make sure the user exists: %s ", err)
	}
	setData(user, data)
	return []*schema.ResourceData{data}, nil
}
