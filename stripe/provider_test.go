package stripe

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	token := ""
	os.Setenv("STRIPE_TOKEN", token)
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"stripe": testAccProvider,
	}
}
func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	t.Log("called")
	if v := os.Getenv("STRIPE_TOKEN"); v == "" {
		t.Fatal("token must be set for acceptance tests")
	}

}
