.PHONY: help

help:
	@echo 'make <target>'
	@echo ''
	@echo 'Targets:'
	@echo ''
	@echo '    help    Show this help'
	@echo ''
	@echo '    build   Build the Terraform module'
	@echo '    clean   Clean up build files'
	@echo '    init    Initialize the Terraform module'
	@echo '    install Install the Terraform binary'
	@echo '    lint    Run the golangci-lint utility'
	@echo '    test    Run the Terraform module tests'
	@echo ''

build: clean init
	@cd fixtures/example1 && terraform apply --auto-approve

clean:
	@cd fixtures/example1 && rm -rf .terraform *.tfstate* .terraform.lock.hcl
	@cd test && rm -f go.mod go.sum

init: clean
	@cd fixtures/example1 && terraform init
	@cd test && go mod init module_test.go; go mod tidy

install:
	@chmod +x ./fixtures/scripts/install_terraform.sh
	@sudo ./fixtures/scripts/install_terraform.sh

lint: clean init
	@cd test && golangci-lint run --print-linter-name --verbose module_test.go

test: clean init
	@cd test && go test -v -destroy
