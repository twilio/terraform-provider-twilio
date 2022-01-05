package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/joho/godotenv"
	"github.com/twilio/terraform-provider-twilio/twilio"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if fileExists("./conf/.env") {
		err := godotenv.Load(`./conf/.env`)
		if err != nil {
			_ = fmt.Errorf("error on loading env %s", err)
			os.Exit(-1)
		}
	}

	opts := &plugin.ServeOpts{
		ProviderFunc: twilio.Provider,
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/twilio/twilio", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
