HOSTNAME=github.com
NAMESPACE=winebarrel
NAME=rds
BINARY=terraform-provider-${NAME}
VERSION=0.1.0
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

default: install

vet:
	go vet ./...

build: vet
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${GOOS}_${GOARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${GOOS}_${GOARCH}
