terraform {
  required_providers{
    stripe ={
      version ="0.2"
      source = "hashicorp.com/edu/stripe"
    }
  }
}

provider "stripe" {
  secretkey = "[SECRET_KEY]"
}


data "stripe_user" "user" {
  email = "[EMAIL]"
}

output "user" {
  value = data.stripe_user.user
}


resource "stripe_user" "user1" {
  email = "[EMAIL]"
  first_name = "[FIRST_NAME]"
  last_name = "[LAST_NAME]"
}
