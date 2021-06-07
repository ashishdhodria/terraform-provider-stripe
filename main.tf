terraform {
  required_providers{
    stripe ={
      version ="0.2"
      source = "hashicorp.com/edu/stripe"
    }
  }
}

provider "stripe" {
  token = "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"
}

/*
data "stripe_user" "user" {
  email = "ashishdhodria1999@gmail.com"
}

output "user" {
  value = data.stripe_user.user
}
*/
resource "stripe_user" "user1" {
  email = "vausdhodria@gmail.com"
  first_name = "vasu"
  last_name = "dhodria"
}

