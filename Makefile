default: build

.PHONY=build
build:
	go build -o terraform-provider-homebrew

test: build
	terraform init
	terraform plan -out terraform.tfplan
	terraform apply terraform.tfplan
	terraform show

.PHONY: authors
authors:
	./scripts/generate_authors
