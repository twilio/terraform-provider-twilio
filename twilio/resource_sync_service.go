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
			},
			"webhooks_from_rest_enabled": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSyncServiceCreate(d *schema.ResourceData, m interface{}) error {

	r, err := m.(*Config).Client.Serverless.Service.Create(resourceSyncServiceParams(d))

	if err != nil {
		return fmt.Errorf("error creating Runtime Service: %s", err)
	}

	d.SetId(*r.SID)

	return resourceSyncServiceRead(d, m)
}

func resourceSyncServiceRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	r, err := m.(*Config).Client.Serverless.Service.Fetch(id)

	d.Set("account_sid", r.AccountSID)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("unique_name", r.UniqueName)
	d.Set("include_credentials", r.IncludeCredentials)

	if err != nil {
		return err
	}

	return nil
}

func resourceSyncServiceUpdate(d *schema.ResourceData, m interface{}) error {
	if _, err := m.(*Config).Client.Serverless.Service.Update(d.Id(), resourceSyncServiceParams(d)); err != nil {
		return fmt.Errorf("error updating Runtime Service: %s", err)
	}

	return resourceSyncServiceRead(d, m)
}

func resourceSyncServiceDelete(d *schema.ResourceData, m interface{}) error {
	if err := m.(*Config).Client.Serverless.Service.Delete(d.Id()); err != nil {
		return fmt.Errorf("error deleting Runtime Service: %s", err)
	}

	return nil
}

func resourceSyncServiceParams(d *schema.ResourceData) *types.RuntimeServiceParams {
	p := &types.RuntimeServiceParams{}

	if v, ok := d.GetOk("friendly_name"); ok {
		p.FriendlyName = util.String(v.(string))
	}

	if v, ok := d.GetOk("unique_name"); ok {
		p.UniqueName = util.String(v.(string))
	}

	if v, ok := d.GetOk("included_credentials"); ok {
		p.IncludeCredentials = util.Bool(v.(bool))
	}

	return p
}
