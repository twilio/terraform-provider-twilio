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

	d.SetId(*r.Sid)

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
	p := &types.ProxyServiceParams{}

	if v, ok := d.GetOk("chat_instance_sid"); ok {
		p.ChatInstanceSID = util.String(v.(string))
	}

	if v, ok := d.GetOk("unique_name"); ok {
		p.UniqueName = util.String(v.(string))
	}

	if v, ok := d.GetOk("callback_url"); ok {
		p.CallbackURL = util.String(v.(string))
	}

	if v, ok := d.GetOk("intercept_callback_url"); ok {
		p.InterceptCallbackURL = util.String(v.(string))
	}

	if v, ok := d.GetOk("out_of_session_callback_url"); ok {
		p.OutOfSessionCallbackURL = util.String(v.(string))
	}

	if v, ok := d.GetOk("default_ttl"); ok {
		p.DefaultTTL = util.Int(v.(int))
	}

	if v, ok := d.GetOk("geo_match_level"); ok {
		p.GeoMatchLevel = util.String(v.(string))
	}

	if v, ok := d.GetOk("number_selection_behavior"); ok {
		p.NumberSelectionBehavior = util.String(v.(string))
	}

	return p
}
