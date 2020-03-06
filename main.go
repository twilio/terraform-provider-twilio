package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/twilio/terraform-provider-twilio/twilio"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: twilio.Provider,
	})
}
