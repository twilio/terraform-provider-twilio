package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/twilio/terraform-provider-twilio/util"
	types "github.com/twilio/twilio-go"
)

func resourceIncomingPhoneNumber() *schema.Resource { // nolint:golint,funlen
	return &schema.Resource{
		Create: resourceIncomingPhoneNumberCreate,
		Read:   resourceIncomingPhoneNumberRead,
		Update: resourceIncomingPhoneNumberUpdate,
		Delete: resourceIncomingPhoneNumberDelete,
		Schema: map[string]*schema.Schema{
			"beta": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"voice_caller_id_lookup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_requirements": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_version": {
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
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sms_application_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_fallback_method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_fallback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_callback": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_callback_method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trunk_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"voice_application_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"voice_fallback_method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"voice_fallback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"voice_method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"voice_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"emergency_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"emergency_address_sid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bundle_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIncomingPhoneNumberParams(d *schema.ResourceData) *types.IncomingPhoneNumberParams {
	p := &types.IncomingPhoneNumberParams{}

	if v, ok := d.GetOk("api_version"); ok {
		p.APIVersion = util.String(v.(string))
	}

	if v, ok := d.GetOk("friendly_name"); ok {
		p.FriendlyName = util.String(v.(string))
	}

	if v, ok := d.GetOk("sms_application_sid"); ok {
		p.SMSApplicationSID = util.String(v.(string))
	}

	if v, ok := d.GetOk("sms_fallback_method"); ok {
		p.SMSFallbackMethod = util.String(v.(string))
	}

	if v, ok := d.GetOk("phone_number"); ok {
		p.PhoneNumber = util.String(v.(string))
	}

	if v, ok := d.GetOk("area_code"); ok {
		p.AreaCode = util.String(v.(string))
	}

	if v, ok := d.GetOk("sms_fallback_url"); ok {
		p.SMSFallbackURL = util.String(v.(string))
	}

	if v, ok := d.GetOk("sms_method"); ok {
		p.SMSMethod = util.String(v.(string))
	}

	if v, ok := d.GetOk("sms_url"); ok {
		p.SMSURL = util.String(v.(string))
	}

	if v, ok := d.GetOk("status_callback"); ok {
		p.StatusCallback = util.String(v.(string))
	}

	if v, ok := d.GetOk("address_sid"); ok {
		p.AddressSID = util.String(v.(string))
	}

	if v, ok := d.GetOk("status_callback_method"); ok {
		p.StatusCallbackMethod = util.String(v.(string))
	}

	if v, ok := d.GetOk("voice_application_sid"); ok {
		p.VoiceApplicationSID = util.String(v.(string))
	}

	if v, ok := d.GetOk("voice_fallback_method"); ok {
		p.VoiceFallbackMethod = util.String(v.(string))
	}

	if v, ok := d.GetOk("voice_fallback_url"); ok {
		p.VoiceFallbackURL = util.String(v.(string))
	}

	if v, ok := d.GetOk("voice_method"); ok {
		p.VoiceMethod = util.String(v.(string))
	}

	if v, ok := d.GetOk("voice_url"); ok {
		p.VoiceURL = util.String(v.(string))
	}

	if v, ok := d.GetOk("voice_receive_mode"); ok {
		p.VoiceReceiveMode = util.String(v.(string))
	}

	if v, ok := d.GetOk("emergency_status"); ok {
		p.EmergencyStatus = util.String(v.(string))
	}

	if v, ok := d.GetOk("emergency_address_sid"); ok {
		p.EmergencyAddressSID = util.String(v.(string))
	}

	if v, ok := d.GetOk("bundle_sid"); ok {
		p.BundleSID = util.String(v.(string))
	}

	if v, ok := d.GetOk("identity_sid"); ok {
		p.IdentitySID = util.String(v.(string))
	}

	if v, ok := d.GetOk("trunk_sid"); ok {
		p.TrunkSID = util.String(v.(string))
	}

	if v, ok := d.GetOk("voice_called_id_lookup"); ok {
		p.VoiceCallerIDLookup = util.Bool(v.(bool))
	}

	return p
}

func resourceIncomingPhoneNumberCreate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	iPNParams := &types.IncomingPhoneNumberParams{PhoneNumber: util.String(d.Get("phone_number").(string))}
	iPN, err := m.(*Config).Client.IncomingPhoneNumbers.Create(iPNParams)

	if err != nil {
		return fmt.Errorf("error creating Phone Number %s", err)
	}

	d.SetId(*iPN.SID)
	d.SetPartial("friendly_name")
	d.Partial(false)

	return resourceIncomingPhoneNumberRead(d, m)
}

func resourceIncomingPhoneNumberRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	iPN, err := m.(*Config).Client.IncomingPhoneNumbers.Read(id)
	if err != nil {
		return err
	}

	d.Set("beta", iPN.Beta)
	d.Set("voice_caller_id_lookup", iPN.VoiceCallerIDLookup)
	d.Set("account_sid", iPN.AccountSID)
	d.Set("address_sid", iPN.AddressSID)
	d.Set("address_requirements", iPN.AddressRequirements)
	d.Set("api_version", iPN.APIVersion)
	d.Set("capabilities", iPN.Capabilities)
	d.Set("date_created", iPN.DateCreated)
	d.Set("date_updated", iPN.DateUpdated)
	d.Set("friendly_name", iPN.FriendlyName)
	d.Set("identity_sid", iPN.IdentitySID)
	d.Set("phone_number", iPN.PhoneNumber)
	d.Set("origin", iPN.Origin)
	d.Set("sid", iPN.SID)
	d.Set("sms_application_sid", iPN.SMSApplicationSID)
	d.Set("sms_fallback_method", iPN.SMSFallbackMethod)
	d.Set("sms_fallback_url", iPN.SMSFallbackURL)
	d.Set("sms_method", iPN.SMSMethod)
	d.Set("sms_url", iPN.SMSURL)
	d.Set("status_callback", iPN.StatusCallback)
	d.Set("status_callback_method", iPN.StatusCallbackMethod)
	d.Set("trunk_sid", iPN.TrunkSID)
	d.Set("uri", iPN.URI)
	d.Set("voice_application_sid", iPN.VoiceApplicationSID)
	d.Set("voice_fallback_method", iPN.VoiceFallbackMethod)
	d.Set("voice_fallback_url", iPN.VoiceFallbackURL)
	d.Set("voice_method", iPN.VoiceMethod)
	d.Set("voice_url", iPN.VoiceURL)
	d.Set("emergency_status", iPN.EmergencyStatus)
	d.Set("emergency_address_sid", iPN.EmergencyAddressSID)
	d.Set("bundle_sid", iPN.BundleSID)

	return nil
}

func resourceIncomingPhoneNumberUpdate(d *schema.ResourceData, m interface{}) error {
	if _, err := m.(*Config).Client.IncomingPhoneNumbers.Update(d.Id(), resourceIncomingPhoneNumberParams(d)); err != nil {
		return fmt.Errorf("error updating Incoming Phone Number: %s", err)
	}

	return resourceIncomingPhoneNumberRead(d, m)
}

func resourceIncomingPhoneNumberDelete(d *schema.ResourceData, m interface{}) error {
	if err := m.(*Config).Client.IncomingPhoneNumbers.Delete(d.Id()); err != nil {
		return fmt.Errorf("error deleting Incoming Phone Number: %s", err)
	}

	return nil
}
