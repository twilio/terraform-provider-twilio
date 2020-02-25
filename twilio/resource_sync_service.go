package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/twilio/terraform-provider-twilio/util"
	types "github.com/twilio/twilio-go"
)

func resourceSyncService() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceSyncServiceCreate,
		Read:   resourceSyncServiceRead,
		Update: resourceSyncServiceUpdate,
		Delete: resourceSyncServiceDelete,
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"webhook_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reachability_webhooks_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"acl_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"reachability_debouncing_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"reachability_debouncing_window": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"webhooks_from_rest_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSyncServiceCreate(d *schema.ResourceData, m interface{}) error {

	r, err := m.(*Config).Client.Sync.Service.Create(resourceSyncServiceParams(d))

	if err != nil {
		return fmt.Errorf("error creating Sync Service: %s", err)
	}

	d.SetId(*r.SID)

	return resourceSyncServiceRead(d, m)
}

func resourceSyncServiceRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	r, err := m.(*Config).Client.Sync.Service.Fetch(id)

	d.Set("account_sid", r.AccountSID)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("unique_name", r.UniqueName)
	d.Set("webhook_url", r.WebhookURL)
	d.Set("reachability_webhooks_enabled", r.ReachabilityWebhooksEnabled)
	d.Set("acl_enabled", r.ACLEnabled)
	d.Set("reachability_debouncing_enabled", r.ReachabilityDebouncingEnabled)
	d.Set("reachability_debouncing_window", r.ReachabilityDebouncingWindow)
	d.Set("webhooks_from_rest_enabled", r.WebhooksFromRestEnabled)

	if err != nil {
		return err
	}

	return nil
}

func resourceSyncServiceUpdate(d *schema.ResourceData, m interface{}) error {
	if _, err := m.(*Config).Client.Sync.Service.Update(d.Id(), resourceSyncServiceParams(d)); err != nil {
		return fmt.Errorf("error updating Sync Service: %s", err)
	}

	return resourceSyncServiceRead(d, m)
}

func resourceSyncServiceDelete(d *schema.ResourceData, m interface{}) error {
	if err := m.(*Config).Client.Sync.Service.Delete(d.Id()); err != nil {
		return fmt.Errorf("error deleting Sync Service: %s", err)
	}

	return nil
}

func resourceSyncServiceParams(d *schema.ResourceData) *types.SyncServiceParams {
	p := &types.SyncServiceParams{}

	if v, ok := d.GetOk("friendly_name"); ok {
		p.FriendlyName = util.String(v.(string))
	}

	if v, ok := d.GetOk("webhook_url"); ok {
		p.WebhookURL = util.String(v.(string))
	}

	if v, ok := d.GetOk("reachability_webhooks_enabled"); ok {
		p.ReachabilityWebhooksEnabled = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("acl_enabled"); ok {
		p.ACLEnabled = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("reachability_webhooks_enabled"); ok {
		p.ReachabilityDebouncingEnabled = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("reachability_webhooks_window"); ok {
		p.ReachabilityDebouncingWindow = util.Int(v.(int))
	}

	if v, ok := d.GetOk("reachability_debouncing_window"); ok {
		p.ReachabilityDebouncingWindow = util.Int(v.(int))
	}

	if v, ok := d.GetOk("webhooks_from_rest_enabled"); ok {
		p.WebhooksFromRestEnabled = util.Bool(v.(bool))
	}

	return p
}
