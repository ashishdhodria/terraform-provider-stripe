package stripe

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccUserDataSource_basic(t *testing.T) {
	t.Log("Test")
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserDataExists("data.stripe_user.user"),
					resource.TestCheckResourceAttr("data.stripe_user.user", "email", "ashishdhodria1999@gmail.com"),
					resource.TestCheckResourceAttr("data.stripe_user.user", "name", "ashish dhodria"),
				),
			},
		},
	})
	log.Println("Test final")
}

func testAccUserDataSourceConfig() string {
	return fmt.Sprintf(
		`data "stripe_user" "user" {
		 email = "ashishdhodria1999@gmail.com"
	}`)
}

func testAccCheckUserDataExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		for _, rs := range state.RootModule().Resources {
			log.Println(rs.Type)
		}
		rs, ok := state.RootModule().Resources[resource]
		log.Println("data exist check ", resource)
		if !ok {
			return fmt.Errorf(" Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf(" No Record ID is set")
		}
		return nil
	}
}
