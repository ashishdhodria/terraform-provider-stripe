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

resource "stripe_user" "SwaraliPankar" {
  email = "swarali@clevertap.com"
  first_name = "Swarali"
  last_name = "Pankar"
}
