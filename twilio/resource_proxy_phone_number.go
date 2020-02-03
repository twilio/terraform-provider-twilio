package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceProxyPhoneNumber() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceProxyPhoneNumberCreate,
		Read:   resourceProxyPhoneNumberRead,
		Update: resourceProxyPhoneNumberUpdate,
		Delete: resourceProxyPhoneNumberDelete,
		Schema: map[string]*schema.Schema{
			"phone_number_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"capabilities": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"is_reserved": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"in_use": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceProxyPhoneNumberCreate(d *schema.ResourceData, m interface{}) error {
	r, err := m.(*Config).Client.Proxy.PhoneNumber.Create(
		d.Get("service_sid").(string),
		&types.ProxyPhoneNumberCreateParams{
			PhoneNumberSID: types.String(d.Get("phone_number_sid").(string)),
			IsReserved:     types.Bool(d.Get("is_reserved").(bool)),
		},
	)

	if err != nil {
		return fmt.Errorf("error creating Proxy Phone Number: %s", err)
	}

	d.SetId(r.SID)

	return resourceProxyPhoneNumberRead(d, m)
}

func resourceProxyPhoneNumberRead(d *schema.ResourceData, m interface{}) error {
	SID := d.Id()
	serviceSID := d.Get("service_sid").(string)
	r, err := m.(*Config).Client.Proxy.PhoneNumber.Read(serviceSID, SID)

	d.Set("sid", SID)
	d.Set("service_sid", serviceSID)
	d.Set("phone_numer", r.PhoneNumber)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("capabilities", r.Capabilities)
	d.Set("is_reserved", r.IsReserved)
	d.Set("in_use", r.InUse)

	if err != nil {
		return err
	}

	return nil
}

func resourceProxyPhoneNumberUpdate(d *schema.ResourceData, m interface{}) error {
	SID := d.Id()
	serviceSID := d.Get("service_sid").(string)
	params := &types.ProxyPhoneNumberUpdateParams{IsReserved: types.Bool(d.Get("is_reserved").(bool))}

	if _, err := m.(*Config).Client.Proxy.PhoneNumber.Update(serviceSID, SID, params); err != nil {
		return fmt.Errorf("error updating Proxy Phone Number: %s", err)
	}

	return resourceProxyPhoneNumberRead(d, m)
}

func resourceProxyPhoneNumberDelete(d *schema.ResourceData, m interface{}) error {
	SID := d.Id()
	serviceSID := d.Get("service_sid").(string)

	if err := m.(*Config).Client.Proxy.PhoneNumber.Delete(serviceSID, SID); err != nil {
		return fmt.Errorf("error deleting Proxy Phone Number: %s", err)
	}

	return nil
}
