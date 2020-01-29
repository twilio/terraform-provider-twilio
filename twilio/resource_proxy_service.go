package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/twilio/terraform-provider-twilio/util"
	types "github.com/twilio/twilio-go"
)

func resourceProxyService() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceProxyServiceCreate,
		Read:   resourceProxyServiceRead,
		Update: resourceProxyServiceUpdate,
		Delete: resourceProxyServiceDelete,
		Schema: map[string]*schema.Schema{
			"unique_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"chat_instance_sid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_ttl": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"geo_match_level": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"number_selection_behavior": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intercept_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"out_of_session_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceProxyServiceCreate(d *schema.ResourceData, m interface{}) error {
	r, err := m.(*Config).Client.Proxy.Service.Create(resourceProxyServiceParams(d))

	if err != nil {
		return fmt.Errorf("error creating Proxy Service: %s", err)
	}

	d.SetId(r.Sid)

	return resourceProxyServiceRead(d, m)
}

func resourceProxyServiceRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	r, err := m.(*Config).Client.Proxy.Service.Read(id, nil)

	d.Set("chat_instance_sid", r.ChatInstanceSID)
	d.Set("unique_name", r.UniqueName)
	d.Set("default_ttl", r.DefaultTTL)
	d.Set("callback_url", r.CallbackURL)
	d.Set("geo_match_level", r.GeoMatchLevel)
	d.Set("number_selection_behavior", r.NumberSelectionBehavior)
	d.Set("intercept_callback_url", nil)
	d.Set("out_of_session_callback_url", r.OutOfSessionCallbackURL)

	if err != nil {
		return err
	}

	return nil
}

func resourceProxyServiceUpdate(d *schema.ResourceData, m interface{}) error {
	if _, err := m.(*Config).Client.Proxy.Service.Update(d.Id(), resourceProxyServiceParams(d)); err != nil {
		return fmt.Errorf("error updating Proxy Service: %s", err)
	}

	return resourceProxyServiceRead(d, m)
}

func resourceProxyServiceDelete(d *schema.ResourceData, m interface{}) error {
	if err := m.(*Config).Client.Proxy.Service.Delete(d.Id(), nil); err != nil {
		return fmt.Errorf("error deleting Proxy Service: %s", err)
	}

	return nil
}

func resourceProxyServiceParams(d *schema.ResourceData) *types.ProxyServiceParams {
	p := &types.ProxyServiceParams{
		ChatInstanceSID:         util.String(d.Get("chat_instance_sid").(string)),
		UniqueName:              util.String(d.Get("unique_name").(string)),
		DefaultTTL:              util.Int(d.Get("default_ttl").(int)),
		CallbackURL:             util.String(d.Get("callback_url").(string)),
		InterceptCallbackURL:    util.String(d.Get("intercept_callback_url").(string)),
		OutOfSessionCallbackURL: util.String(d.Get("out_of_session_callback_url").(string)),
	}

	if g := d.Get("geo_match_level").(string); g != "" {
		p.GeoMatchLevel = util.String(g)
	}

	if n := d.Get("number_selection_behavior").(string); n != "" {
		p.NumberSelectionBehavior = util.String(n)
	}

	return p
}
