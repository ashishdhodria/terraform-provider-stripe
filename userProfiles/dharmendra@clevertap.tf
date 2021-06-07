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

resource "stripe_user" "Dharmendra" {
  email = "dharmendra@clevertap.com"
  first_name = "Dharmendra"
  last_name = ""
}
