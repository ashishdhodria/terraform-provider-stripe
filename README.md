
# Terraform Stripe Provider

This terraform provider allows to perform Create ,Read ,Update, Delete and Import stripe User(s). 


## Requirements

* [Go](https://golang.org/doc/install) 1.16 <br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Stripe](https://stripe.com/docs/api) Developers account 


## Setup Stripe Account
 :heavy_exclamation_mark:  [IMPORTANT] : This provider can be successfully tested on any Stripe developer account. <br><br>

1. Create a Stripe account. (https://dashboard.stripe.com/register)<br>
2. Sign in to the stripe account (https://dashboard.stripe.com/login)<br>
3. Go to the Get your API keys.<br>
This app will provide us with the Secret Key which will be needed to configure our provider and make request. <br>


## Initialise Stripe Provider in local machine 
1. Clone the repository  to $GOPATH/src/github.com/stripe/terraform-provider-stripe <br>
2. Add the Refresh Secret Key generted in Stripe App to respective fields in `main.tf` <br>
3. Run the following command :
 ```golang
go mod init terraform-provider-stripe
go mod tidy
```
4. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>

## Installation
1. Run the following command to create a vendor subdirectory which will comprise of  all provider dependencies. <br>
```
~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}
``` 
Command: 
```bash
mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/stripe/0.2.0/[OS_ARCH]
```
For eg. `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/stripe/0.2.0/windows_amd64`<br>

2. Run `go build -o terraform-provider-stripe.exe`. This will save the binary (`.exe`) file in the main/root directory. <br>
3. Run this command to move this binary file to appropriate location.
 ```
 move terraform-provider-stripe.exe %APPDATA%\terraform.d\plugins\hashicorp.com\edu\stripe\0.2.0\[OS_ARCH]
 ``` 
Otherwise you can manually move the file from current directory to destination directory.<br>


[OR]

1. Download required binaries <br>
2. move binary `~/.terraform.d/plugins/[architecture name]/`


## Run the Terraform provider

#### Create User
1. Add the user email, first name, last name in the respective field in `main.tf`
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully created.

#### Update the user
Update the data of the user in the `main.tf` file and apply using `terraform apply`

#### Read the User Data
Add data and output blocks in the `main.tf` file and run `terraform plan` to read user data.

#### Delete the user
Delete the resource block of the particular user from `main.tf` file and run `terraform apply`.

#### Import a User Data
1. Write manually a resource configuration block for the User in `main.tf`, to which the imported object will be mapped.
2. Run the command `terraform import stripe_user.user1 [EMAIL_ID]`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in resource block.


### Testing the Provider
1. Navigate to the test file directory.
2. Run command `go test` . This command will give combined test result for the execution or errors if any failure occur.
3. If you want to see test result of each test function individually while running test in a single go, run command `go test -v`
4. To check test cover run `go test -cover`


## Example Usage
```terraform
terraform {
  required_providers{
    stripe ={
      version ="0.2"
      source = "hashicorp.com/edu/stripe"
    }
  }
}

provider "stripe" {
  token = "[Secret Key]"
}


data "stripe_user" "user" {
  email = "[EMAIL]"
}

output "user" {
  value = data.stripe_user.user
}

resource "stripe_user" "user1" {
  email = "[EMAIL]"
  first_name = "[FIRST NAME]"
  last_name = "[LAST_NAME]"
}
```
