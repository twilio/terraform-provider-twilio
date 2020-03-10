TEST?=$$(go list ./... |grep -v 'vendor')

default: build

build:
	go build -o terraform-provider-twilio

test: 
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m



