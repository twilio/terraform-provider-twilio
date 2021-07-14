.PHONY: default githooks build goimports govet golint terrafmt install test testacc

TEST?=$$(go list ./... |grep -v 'vendor')
REGISTRY=local
NAMESPACE=twilio
NAME=twilio
BINARY=terraform-provider-${NAME}
VERSION=0.5.0
OS_ARCH=darwin_amd64

default: build

githooks:
	ln -sf ../../githooks/pre-commit .git/hooks/pre-commit

build: githooks goimports terrafmt
	go build -o ${BINARY}

goimports:
	go get golang.org/x/tools/cmd/goimports
	goimports -w .

govet: goimports
	go vet ./...

golint: govet
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0
	golangci-lint run

terrafmt:
	terraform fmt -recursive

install: build
	mkdir -p ~/.terraform.d/plugins/${REGISTRY}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${REGISTRY}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test: build
	go test $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
