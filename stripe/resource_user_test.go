package stripe

import (
	"fmt"
	"log"
	"terraform-provider-stripe/client"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccItem_Basic(t *testing.T) {
	t.Log("Test")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("stripe_user.user", "first_name", "ashish"),
					resource.TestCheckResourceAttr("stripe_user.user", "last_name", "dhodria"),
					resource.TestCheckResourceAttr("stripe_user.user", "email", "ashishdhodria27@gmail.com"),
				),
			},
		},
	})
	log.Println("Test final")
}

func testAccCheckItemDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*client.Client)
	time.Sleep(10 * time.Second)
	log.Println("destroy")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "stripe_user_instance" {
			continue
		}
		Id := rs.Primary.ID
		_, err := apiClient.DeleteItem(Id)
		if err != nil {
			return fmt.Errorf("status: %v", err)
		}
	}
	log.Println("destroy final")
	return nil
}

func TestAccItem_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("stripe_user.user_update", "first_name", "ashish"),
					resource.TestCheckResourceAttr("stripe_user.user_update", "last_name", "dhodria"),
					resource.TestCheckResourceAttr("stripe_user.user_update", "email", "ashishdhodria27@gmail.com"),
				),
			},
			{
				Config: testAccCheckItemUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("stripe_user.user_update", "first_name", "ashu"),
					resource.TestCheckResourceAttr("stripe_user.user_update", "last_name", "kumar"),
					resource.TestCheckResourceAttr("stripe_user.user_update", "email", "ashishdhodria27@gmail.com"),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
resource "stripe_user" "user" {
	first_name = "ashish"
	last_name  = "dhodria"
	email      = "ashishdhodria27@gmail.com"
}
`)
}

func testAccCheckItemUpdatePre() string {
	return fmt.Sprintf(`
resource "stripe_user" "user_update" {
	   first_name = "ashish"
	   last_name  = "dhodria"
	   email      = "ashishdhodria27@gmail.com"
}
`)
}

func testAccCheckItemUpdatePost() string {
	return fmt.Sprintf(`
resource "stripe_user" "user_update" {
	first_name = "ashu"
	last_name  = "kumar"
	email      = "ashishdhodria27@gmail.com"
}
`)
}
