package twilio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	client "github.com/twilio/twilio-go"
)

// Provider initializes terraform-provider-twilio.
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				Description: "Your Account SID can be found on the Twilio dashboard.",
				Required:    true,
			},
			"auth_token": {
				Type:        schema.TypeString,
				Description: "Your Auth Token can be found on the Twilio dashboard.",
				Required:    true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"twilio_chat_service": resourceChatService(),
		},
	}

	p.ConfigureFunc = providerClient(p)

	return p
}

// Config is provided as context to the underlying resources.
type Config struct {
	Client *client.Twilio
}

func providerClient(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := &Config{
			Client: client.NewClient(d.Get("account_sid").(string), d.Get("auth_token").(string)),
		}

		return config, nil
	}
}
