provider "homebrew" {
  path = "/usr/local/bin/brew"
  host = "127.0.0.1"
}

locals {
  developer_tools = [
    { name = "zsh" },
    { name = "vim" },
    { name = "jq" },
  ]
}

module "development_packages" {
  source   = "modules/homebrew_package"
  packages = ["${local.developer_tools}"]
}
