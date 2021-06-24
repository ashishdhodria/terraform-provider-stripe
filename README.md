This terraform provider allows to perform Create ,Read ,Update, Delete and Import stripe Users. 


## Requirements

* [Go](https://golang.org/doc/install) 1.16 (To build the provider plugin)<br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Stripe](https://stripe.com/docs/api) Stripe API Documentation 

## Application Account

### Setup
1. Create a Stripe account. (https://dashboard.stripe.com/register)<br>
2. Sign in to the stripe account (https://dashboard.stripe.com/login)<br>

### API Authentication
Go to the Developers -> API Keys -> Standard Keys and select Secret Key.<br>
This app will provide us with the Secret Key which will be needed to configure our provider and make request. <br>

## Building The Provider
1. Clone the repository, add all the dependencies and create a vendor directory that contains all dependencies. For this, run the following commands: <br>
 ```golang
cd terraform-provider-stripe
go mod init terraform-provider-stripe
go mod tidy
go mod vendor
```

## Managing terraform plugins
*For Windows:*
1. Run the following command to create a vendor sub-directory (`%APPDATA%/terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${OS_ARCH}`) which will consist of all terraform plugins. <br> 
Command: 
```bash
mkdir -p %APPDATA%/terraform.d/plugins/hashicorp.com/edu/stripe/0.2.0/windows_amd64
```
2. Run `go build -o terraform-provider-stripe.exe` to generate the binary in present working directory. <br>
3. Run this command to move this binary file to the appropriate location.
 ```
 move terraform-provider-stripe.exe %APPDATA%\terraform.d\plugins\hashicorp.com\edu\stripe\0.2.0\windows_amd64
 ``` 
<p align="center">[OR]</p>
 
3. Manually move the file from current directory to destination directory (`%APPDATA%\terraform.d\plugins\hashicorp.com\edu\stripe\0.2.0\windows_amd64`).<br>

## Working with terraform


### Application Credential Integration in terraform
1. Add `terraform` block and `provider` block as shown in [example usage](#example-usage).
2. Get the credentials: secretkey
3. Assign the above credentials to the respective field in the `provider` block.

### Basic Terraform Commands
1. `terraform init` - To initialize a working directory containing Terraform configuration files.
2. `terraform plan` - To create an execution plan. Displays the changes to be done.
3. `terraform apply` - To execute the actions proposed in a Terraform plan. Apply the changes.

### Create User
1. Add the user email, name in the respective field in `resource` block as shown in [example usage](#example-usage).
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully created.

### Update the user
Update the data of the user in `resource` block as show in [example usage](#example-usage) file and apply using `terraform apply`

### Read the User Data
Add data and output blocks as shown in the [example usage](#example-usage) and run `terraform plan` to read user data.

### Delete the user
Delete the resource block of the particular user from `main.tf` file and run `terraform apply`.

### Import a User Data
1. Write manually a `resource` configuration block for the user as shown in [example usage](#example-usage). Imported user will be mapped to this block.
2. Run the command `terraform import stripe_user.user1 [EMAIL_ID]`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in resource block.

## Example Usage<a id="example-usage"></a>
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
  secretkey = _REPLACE_STRIPE_USER_SECRETKEY"
}


data "stripe_user" "user" {
  email = user@domain.com"
}

output "user" {
  value = data.stripe_user.user
}

resource "stripe_user" "user1" {
  email = "user@domain.com"
  name = "User_Name"
}
```
## Argument Reference

* `secretkey`(Required,string)     - The stripe secret Key from created application
* `name`(Required,string) - Name of the User.
* `email`(Required,string)         - Email of the user.
