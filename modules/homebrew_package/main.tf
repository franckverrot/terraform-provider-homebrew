resource "homebrew_package" "packages" {
  count = "${length(var.packages)}"
  name  = "${lookup(var.packages[count.index], "name")}"
}