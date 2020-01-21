package twilio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceChatService() *schema.Resource {
	return &schema.Resource{
		Create: resourceChatServiceCreate,
		Read:   resourceChatServiceRead,
		Update: resourceChatServiceUpdate,
		Delete: resourceChatServiceDelete,

		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"default_service_role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_channel_role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_channel_creator_role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"read_status_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"reachability_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"typing_indicator_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"consumption_report_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pre_webhook_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"post_webhook_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webhook_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webhook_filters": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pre_webhook_retry_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"post_webhook_retry_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceChatServiceParams(d *schema.ResourceData) *types.ChatServiceParams {
	return &types.ChatServiceParams{
		FriendlyName:                 d.Get("friendly_name").(string),
		DefaultServiceRoleSid:        d.Get("default_service_role_sid").(string),
		DefaultChannelRoleSid:        d.Get("default_channel_role_sid").(string),
		DefaultChannelCreatorRoleSid: d.Get("default_channel_creator_role_sid").(string),
		ReadStatusEnabled:            d.Get("read_status_enabled").(bool),
		ReachabilityEnabled:          d.Get("reachability_enabled").(bool),
		TypingIndicatorTimeout:       d.Get("typing_indicator_timeout").(int),
		ConsumptionReportInterval:    d.Get("consumption_report_interval").(int),
		PreWebhookURL:                d.Get("pre_webhook_url").(string),
		PostWebhookURL:               d.Get("post_webhook_url").(string),
		WebhookMethod:                d.Get("webhook_method").(string),
		WebhookFilters:               d.Get("webhook_filters").(*schema.Set).List(),
		PreWebhookRetryCount:         d.Get("pre_webhook_retry_count").(int),
		PostWebhookRetryCount:        d.Get("post_webhook_retry_count").(int),
	}
}

func resourceChatServiceCreate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	chatService, err := m.(*Config).Client.Chat.Create(&types.ChatServiceParams{
		FriendlyName: d.Get("friendly_name").(string),
	})

	if err != nil {
		return err
	}

	d.SetPartial("friendly_name")

	if _, err := m.(*Config).Client.Chat.Update(chatService.Sid, resourceChatServiceParams(d)); err != nil {
		return err
	}

	d.SetPartial("default_service_role_sid")
	d.SetPartial("default_channel_role_sid")
	d.SetPartial("default_channel_creator_role_sid")
	d.SetPartial("read_status_enabled")
	d.SetPartial("reachability_enabled")
	d.SetPartial("typing_indicator_timeout")
	d.SetPartial("consumption_report_interval")
	d.SetPartial("pre_webhook_url")
	d.SetPartial("post_webhook_url")
	d.SetPartial("webhook_method")
	d.SetPartial("webhook_filters")

	d.Partial(false)

	d.SetId(chatService.Sid)

	return resourceChatServiceRead(d, m)
}

func resourceChatServiceRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	chatService, err := m.(*Config).Client.Chat.Read(id, nil)

	if err != nil {
		return err
	}

	d.Set("sid", chatService.Sid)
	d.Set("account_sid", chatService.AccountSid)
	d.Set("friendly_name", chatService.FriendlyName)
	d.Set("date_created", chatService.DateCreated)
	d.Set("date_updated", chatService.DateUpdated)
	d.Set("default_service_role_sid", chatService.DefaultServiceRoleSid)
	d.Set("default_channel_role_sid", chatService.DefaultChannelRoleSid)
	d.Set("default_channel_creator_role_sid", chatService.DefaultChannelCreatorRoleSid)
	d.Set("read_status_enabled", chatService.ReadStatusEnabled)
	d.Set("reachability_enabled", chatService.ReachabilityEnabled)
	d.Set("typing_indicator_timeout", chatService.TypingIndicatorTimeout)
	d.Set("consumption_report_interval", chatService.ConsumptionReportInterval)
	d.Set("limits", chatService.Limits)
	d.Set("pre_webhook_url", chatService.PreWebhookURL)
	d.Set("post_webhook_url", chatService.PostWebhookURL)
	d.Set("webhook_method", chatService.WebhookMethod)
	d.Set("webhook_filters", chatService.WebhookFilters)
	d.Set("pre_webhook_retry_count", chatService.PreWebhookRetryCount)
	d.Set("post_webhook_retry_count", chatService.PostWebhookRetryCount)
	d.Set("notifications", chatService.Notifications)
	d.Set("media", chatService.Media)
	d.Set("url", chatService.URL)
	d.Set("links", chatService.Links)

	return nil
}

func resourceChatServiceUpdate(d *schema.ResourceData, m interface{}) error {
	chatService, err := m.(*Config).Client.Chat.Update(d.Id(), resourceChatServiceParams(d))

	if err != nil {
		return err
	}

	d.SetId(chatService.Sid)

	return resourceChatServiceRead(d, m)
}

func resourceChatServiceDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	_, err := m.(*Config).Client.Chat.Delete(id, nil)

	if err != nil {
		return err
	}

	return nil
}
