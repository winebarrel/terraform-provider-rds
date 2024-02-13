HOSTNAME=registry.terraform.io
NAMESPACE=winebarrel
NAME=rds
BINARY=terraform-provider-${NAME}
VERSION=0.1.3
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

default: build

vet:
	go vet ./...

build: vet
	go build -o ${BINARY}

lint:
	golangci-lint run .

json:
	curl -sSfL https://instances.vantage.sh/rds/instances.json \
	| jq 'map({key: .instanceType,value: .memory | tonumber}) | sort_by(.key) | from_entries' > rds/rds.json

dev.tfrc: dev.tfrc.tpl
	sed "s|{{PATH_TO_PROVIDER}}|$(shell pwd)|" dev.tfrc.tpl > dev.tfrc

.PHONY: tf-plan
tf-plan: build dev.tfrc
	TF_CLI_CONFIG_FILE=dev.tfrc terraform plan

.PHONY: tf-apply
tf-apply: build dev.tfrc
	TF_CLI_CONFIG_FILE=dev.tfrc terraform apply -auto-approve

.PHONY: tf-clean
tf-clean: clean
	rm -f dev.tfrc terraform.tfstate*
