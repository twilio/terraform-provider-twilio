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

const (
	AccountSid    = "TWILIO_ACCOUNT_SID"
	AuthToken     = "TWILIO_AUTH_TOKEN"
	ApiKey        = "TWILIO_API_KEY"
	ApiSecret     = "TWILIO_API_SECRET"
	SubAccountSid = "TWILIO_SUBACCOUNT_SID"
	Edge          = "TWILIO_EDGE"
	Region        = "TWILIO_REGION"
)

// Provider initializes terraform-provider-twilio.
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{ApiKey, AccountSid}, nil),
				Required:    true,
				Description: "Your Account SID/ API Key can be found on the Twilio dashboard at www.twilio.com/console.",
			},
			"password": {
				Type:        schema.TypeString,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{ApiSecret, AuthToken}, nil),
				Description: "Your Auth Token/ API Secret can be found on the Twilio dashboard at www.twilio.com/console.",
				Required:    true,
			},
			"account_sid": {
				Type:        schema.TypeString,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{SubAccountSid, AccountSid}, nil),
				Description: "Your SubAccount SID can be found on the Twilio dashboard at www.twilio.com/console.",
				Optional:    true,
			},
			"edge": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc(Edge, nil),
				Description: "https://www.twilio.com/docs/global-infrastructure/edge-locations#public-edge-locations",
				Optional:    true,
			},
			"region": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc(Region, nil),
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

		username := d.Get("username").(string)
		password := d.Get("password").(string)
		region := d.Get("region").(string)
		edge := d.Get("edge").(string)

		params := client.ClientParams{
			Username: username,
			Password: password,
		}

		if d.Get("account_sid") != nil {
			params.AccountSid = d.Get("account_sid").(string)
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
