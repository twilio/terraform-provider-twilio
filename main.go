package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/joho/godotenv"
	"github.com/twilio/terraform-provider-twilio/twilio"
)

func main() {
	if fileExists("./conf/.env") {
		err := godotenv.Load(`./conf/.env`)
		if err != nil {
			_ = fmt.Errorf("error on loading env %s", err)
			os.Exit(-1)
		}
	}
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: twilio.Provider,
	})
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
