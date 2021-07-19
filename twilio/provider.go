// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	twilio "github.com/twilio/terraform-provider-twilio/client"
	"github.com/twilio/terraform-provider-twilio/twilio/resources"
	client "github.com/twilio/twilio-go"
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
			"subaccount_sid": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_SUBACCOUNT_SID", nil),
				Description: "Your SubAccount SID can be found on the Twilio dashboard at www.twilio.com/console.",
				Optional:    true,
			},
			"edge": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_EDGE", nil),
				Description: "https://www.twilio.com/docs/global-infrastructure/edge-locations#public-edge-locations",
				Optional:    true,
			},
			"region": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_REGION", nil),
				Description: "https://www.twilio.com/docs/global-infrastructure/edge-locations/legacy-regions",
				Optional:    true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap:   resources.NewTwilioResources().Map,
	}

	p.ConfigureContextFunc = providerClient(p)

	return p
}

func providerClient(p *schema.Provider) schema.ConfigureContextFunc {
	return func(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var TwilioClient *client.RestClient

		username := d.Get("account_sid").(string)
		password := d.Get("auth_token").(string)
		region := d.Get("region").(string)
		edge := d.Get("edge").(string)

		params := client.RestClientParams{
			Username: username,
			Password: password,
		}
		if d.Get("subaccount_sid") != nil {
			params.AccountSid = d.Get("subaccount_sid").(string)
		}
		TwilioClient = client.NewRestClientWithParams(params)

		TwilioClient.SetRegion(region)
		TwilioClient.SetEdge(edge)

		config := &twilio.Config{
			Client: TwilioClient,
		}

		return config, nil
	}
}
