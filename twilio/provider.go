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
				Description: "Your Account SID can be found on the Twilio dashboard at www.twilio.com/console.",
				Required:    true,
			},
			"auth_token": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_AUTH_TOKEN", nil),
				Description: "Your Auth Token can be found on the Twilio dashboard at www.twilio.com/console.",
				Required:    true,
			},
			"edge": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_EDGE", nil),
				Description: "https://www.twilio.com/docs/global-infrastructure/edge-locations#public-edge-locations",
				Optional:    true,
				Required:    false,
			},
			"region": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_REGION", nil),
				Description: "https://www.twilio.com/docs/global-infrastructure/edge-locations/legacy-regions",
				Optional:    true,
				Required:    false,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"twilio_studio_flow_v2":                   studioV2.ResourceFlows(),
			"twilio_api_message_v2010":                apiV2010.ResourceAccountsMessages(),
			"twilio_api_call_v2010":                   apiV2010.ResourceAccountsCalls(),
			"twilio_api_incoming_phone_numbers_v2010": apiV2010.ResourceAccountsIncomingPhoneNumbers(),
			"twilio_api_keys_v2010":                   apiV2010.ResourceAccountsSigningKeys(),
			"twilio_serverless_function_v1":           serverlessV1.ResourceServicesFunctions(),
			"twilio_serverless_service_v1":            serverlessV1.ResourceServices(),
		},
	}

	p.ConfigureContextFunc = providerClient(p)

	return p
}

func providerClient(p *schema.Provider) schema.ConfigureContextFunc {
	return func(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		accountSid := d.Get("account_sid").(string)
		authToken := d.Get("auth_token").(string)
		edge := d.Get("edge").(string)
		region := d.Get("region").(string)
		client := client.NewClient(accountSid, authToken)
		client.SetRegion(region)
		client.SetEdge(edge)
		config := &twilio.Config{
			Client: client,
		}

		return config, nil
	}
}
