package twilio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	client "github.com/twilio/twilio-go"
)

// Provider initializes terraform-provider-twilio.
func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("ACCOUNT_SID", nil),
				Description: "Your Account SID can be found on the Twilio dashboard.",
				Required:    true,
			},
			"auth_token": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("AUTH_TOKEN", nil),
				Description: "Your Auth Token can be found on the Twilio dashboard.",
				Required:    true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"twilio_available_phone_numbers_local": dataSourceAvailablePhoneNumbersLocal(),
			"twilio_taskrouter_taskqueues":         dataSourceTaskRouterTaskQueues(),
			"twilio_taskrouter_workflows":          dataSourceTaskRouterTaskQueues(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"twilio_chat_service":          resourceChatService(),
			"twilio_incoming_phone_number": resourceIncomingPhoneNumber(),
			"twilio_proxy_service":         resourceProxyService(),
			"twilio_proxy_phone_number":    resourceProxyPhoneNumber(),
			"twilio_studio_flow":           resourceStudioFlow(),
			"twilio_sync_service":          resourceSyncService(),
			"twilio_taskrouter_workspace":  resourceTaskRouterWorkspace(),
			"twilio_taskrouter_workflow":   resourceTaskRouterWorkflow(),
			"twilio_taskrouter_taskqueue":  resourceTaskRouterTaskQueue(),
			"twilio_taskrouter_activity":   resourceTaskRouterActivity(),
			"twilio_runtime_service":       resourceRuntimeService(),
			"twilio_runtime_environment":   resourceRuntimeEnvironment(),
			"twilio_flex_flow":             resourceFlexFlow(),
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
