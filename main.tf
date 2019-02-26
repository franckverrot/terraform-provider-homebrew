provider "homebrew" {
  path = "/usr/local/bin/brew"
  host = "127.0.0.1"
}

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
