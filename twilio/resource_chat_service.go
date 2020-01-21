package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	twilio "github.com/twilio/twilio-go"
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
			"read_status_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
			},
			"reachability_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
			},
		},
	}
}

func resourceChatServiceParams(d *schema.ResourceData) *twilio.ChatServiceParams {
	return &twilio.ChatServiceParams{
		FriendlyName:                 d.Get("friendly_name").(string),
		DefaultServiceRoleSid:        d.Get("default_service_role_sid").(string),
		DefaultChannelRoleSid:        d.Get("default_channel_role_sid").(string),
		DefaultChannelCreatorRoleSid: d.Get("default_channel_creator_role_sid").(string),
		ReadStatusEnabled:            d.Get("read_status_enabled").(string),
		ReachabilityEnabled:          d.Get("reachability_enabled").(string),
		TypingIndicatorTimeout:       d.Get("typing_indicator_timeout").(string),
		ConsumptionReportInterval:    d.Get("consumption_report_interval").(string),
		PreWebhookURL:                d.Get("pre_webhook_url").(string),
		PostWebhookURL:               d.Get("post_webhook_url").(string),
		WebhookMethod:                d.Get("webhook_method").(string),
		WebhookFilters:               d.Get("webhook_filters").(string),
		PreWebhookRetryCount:         d.Get("pre_webhook_retry_count").(int),
		PostWebhookRetryCount:        d.Get("post_webhook_retry_count").(int),
	}
}

func resourceChatServiceCreate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	chatService, err := m.(*Config).Client.Chat.Create(&twilio.ChatServiceParams{
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
