package twilio

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	twilio "github.com/twilio/terraform-provider-twilio/client"
	apiV2010 "github.com/twilio/terraform-provider-twilio/twilio/resources/api/v2010"
	serverlessV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/serverless/v1"
	studioV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/studio/v2"
	client "github.com/twilio/twilio-go/twilio"
)

// Provider initializes terraform-provider-twilio.
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_ACCOUNT_SID", nil),
				Description: "Your Account SID can be found on the Twilio dashboard.",
				Required:    true,
			},
			"auth_token": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_AUTH_TOKEN", nil),
				Description: "Your Auth Token can be found on the Twilio dashboard.",
				Required:    true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"twilio_studio_flow_v2":    studioV2.ResourceFlows(),
			"twilio_api_message_v2010": apiV2010.ResourceAccountsMessages(),
			"twilio_api_call_v2010":    apiV2010.ResourceAccountsCalls(),
			"twilio_serverless_v1":     serverlessV1.ResourceServicesFunctions(),
		},
	}

	p.ConfigureContextFunc = providerClient(p)

	return p
}

func providerClient(p *schema.Provider) schema.ConfigureContextFunc {
	return func(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		config := &twilio.Config{
			Client: client.NewClient(d.Get("account_sid").(string), d.Get("auth_token").(string)),
		}

		return config, nil
	}
}
