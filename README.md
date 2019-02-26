# Terraform Homebrew Provider

## DISCLAIMER

This is Work In Progress: Use it at your own risk :-)

This provider makes it easy to provision a MacOS setup.

It requires having:

  * A SSH access to that box (as easy as `sudo systemsetup -setremotelogin on` on MacOS)
  * [Homebrew](https://brew.sh/) to be already installed


## Requirements

*	[Terraform](https://www.terraform.io/downloads.html) 0.11.x
*	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)


## Building The Provider

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-homebrew`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-homebrew
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-homebrew
$ make build
```

## Using the provider

### Basic Usage

Installing a simple package could be done like this:

```hcl
resource "homebrew_package" "zsh" {
  name = "zsh"
}
```

But you can also create a module out of it, like the one you can found in
`modules` in this repository, which then allows you to automate the
installation of more than one package!

```hcl
variable "developer_tools" {
  default = [
    { name = "zsh" },
    { name = "vim" },
    { name = "jq" },
  ]
}

module "development_packages" {
  source   = "modules/homebrew_package"
  packages = ["${var.developer_tools}"]
}
```

Running this content gives:

```
λ go build -o terraform-provider-homebrew

λ terraform init
Initializing modules...
- module.development_packages

Initializing provider plugins...

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.


λ terraform plan -out terraform.tfplan
Refreshing Terraform state in-memory prior to plan...
The refreshed state will be used to calculate this plan, but will not be
persisted to local or remote state storage.


------------------------------------------------------------------------

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  + module.development_packages.homebrew_package.packages[0]
      id:   <computed>
      name: "zsh"

  + module.development_packages.homebrew_package.packages[1]
      id:   <computed>
      name: "vim"

  + module.development_packages.homebrew_package.packages[2]
      id:   <computed>
      name: "jq"


Plan: 3 to add, 0 to change, 0 to destroy.

------------------------------------------------------------------------

This plan was saved to: terraform.tfplan

To perform exactly these actions, run the following command to apply:
    terraform apply "terraform.tfplan"


λ terraform apply terraform.tfplan
module.development_packages.homebrew_package.packages[2]: Creating...
  name: "" => "jq"
module.development_packages.homebrew_package.packages[0]: Creating...
  name: "" => "zsh"
module.development_packages.homebrew_package.packages[1]: Creating...
  name: "" => "vim"
module.development_packages.homebrew_package.packages[0]: Creation complete after 6s (ID: zsh)
module.development_packages.homebrew_package.packages[1]: Creation complete after 6s (ID: vim)
module.development_packages.homebrew_package.packages[2]: Creation complete after 7s (ID: jq)

Apply complete! Resources: 3 added, 0 changed, 0 destroyed.


λ terraform show

module.development_packages.homebrew_package.packages.0:
  id = zsh
  name = zsh
module.development_packages.homebrew_package.packages.1:
  id = vim
  name = vim
module.development_packages.homebrew_package.packages.2:
  id = jq
  name = jq
```

### Known gotchas

* It's kinda slow
* If you're going with the module, you can easily add new packages but
  reordering them might end up with a lot of deletions/recreations.


## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-homebrew
...
```


## License

Mozilla Public License Version 2.0 – Franck Verrot – Copyright 2019