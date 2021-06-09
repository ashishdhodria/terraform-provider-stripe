terraform {
  required_providers{
    stripe ={
      version ="0.2"
      source = "hashicorp.com/edu/stripe"
    }
  }
}

provider "stripe" {
  token = ""
}

data "stripe_user" "user" {
  email = "ashishdhodria1999@gmail.com"
}

output "user" {
  value = data.stripe_user.user
}

resource "stripe_user" "user1" {
  email = ""
  first_name = ""
  last_name = ""
}

