module github.com/twilio/terraform-provider-twilio

go 1.16

replace github.com/twilio/terraform-provider-twilio/core => ./core

require (
	github.com/golangci/golangci-lint v1.39.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/joho/godotenv v1.3.0
	github.com/stretchr/testify v1.7.0
	github.com/twilio/terraform-provider-twilio/core v0.0.0 // indirect
	github.com/twilio/twilio-go v0.10.1-0.20210623172131-41aed5988cfd
	golang.org/x/tools v0.1.4 // indirect
)
