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
			},
			"bundle_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIncomingPhoneNumberParams(d *schema.ResourceData) *types.IncomingPhoneNumberParams {
	return &types.IncomingPhoneNumberParams{
		AccountSID:           util.String(d.Get("accound_sid").(string)),
		APIVersion:           util.String(d.Get("api_version").(string)),
		FriendlyName:         util.String(d.Get("friendly_name").(string)),
		SMSApplicationSid:    util.String(d.Get("sms_application_sid").(string)),
		SMSFallbackMethod:    util.String(d.Get("sms_fallback_method").(string)),
		PhoneNumber:          util.String(d.Get("phone_number").(string)),
		AreaCode:             util.String(d.Get("area_code").(string)),
		SMSFallbackURL:       util.String(d.Get("sms_fallback_url").(string)),
		SMSMethod:            util.String(d.Get("sms_method").(string)),
		SMSURL:               util.String(d.Get("sms_url").(string)),
		StatusCallback:       util.String(d.Get("status_callback").(string)),
		AddressSID:           util.String(d.Get("address_sid").(string)),
		StatusCallbackMethod: util.String(d.Get("status_callback_method").(string)),
		VoiceApplicationSID:  util.String(d.Get("voice_application_sid").(string)),
		VoiceCallerIDLookup:  util.Bool(d.Get("voice_called_id_lookup").(bool)),
		VoiceFallbackMethod:  util.String(d.Get("voice_fallback_method").(string)),
		VoiceFallbackURL:     util.String(d.Get("voice_fallback_url").(string)),
		VoiceMethod:          util.String(d.Get("voice_method").(string)),
		VoiceURL:             util.String(d.Get("voice_url").(string)),
		VoiceReceiveMode:     util.String(d.Get("voice_receive_mode").(string)),
		EmergencyStatus:      util.String(d.Get("emergency_status").(string)),
		EmergencyAddressSID:  util.String(d.Get("emergency_address_sid").(string)),
		BundleSID:            util.String(d.Get("bundle_sid").(string)),
		IdentitySID:          util.String(d.Get("identity_sid").(string)),
		TrunkSID:             util.String(d.Get("trunk_sid").(string)),
	}
}

func resourceIncomingPhoneNumberCreate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	iPNParams := &types.IncomingPhoneNumberParams{PhoneNumber: util.String(d.Get("phone_number").(string))}
	iPN, err := m.(*Config).Client.IncomingPhoneNumbers.Create(iPNParams)

	if err != nil {
		return fmt.Errorf("error creating Phone Number %s", err)
	}

	d.SetId(*iPN.Sid)
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
	d.Set("account_sid", iPN.AccountSid)
	d.Set("address_sid", iPN.AddressSid)
	d.Set("address_requirements", iPN.AddressRequirements)
	d.Set("api_version", iPN.APIVersion)
	d.Set("capabilities", iPN.Capabilities)
	d.Set("date_created", iPN.DateCreated)
	d.Set("date_updated", iPN.DateUpdated)
	d.Set("friendly_name", iPN.FriendlyName)
	d.Set("identity_sid", iPN.IdentitySid)
	d.Set("phone_number", iPN.PhoneNumber)
	d.Set("origin", iPN.Origin)
	d.Set("sid", iPN.Sid)
	d.Set("sms_application_sid", iPN.SMSApplicationSid)
	d.Set("sms_fallback_method", iPN.SMSFallbackMethod)
	d.Set("sms_fallback_url", iPN.SMSFallbackURL)
	d.Set("sms_method", iPN.SMSMethod)
	d.Set("sms_url", iPN.SMSURL)
	d.Set("status_callback", iPN.StatusCallback)
	d.Set("status_callback_method", iPN.StatusCallbackMethod)
	d.Set("trunk_sid", iPN.TrunkSid)
	d.Set("uri", iPN.URI)
	d.Set("voice_application_sid", iPN.VoiceApplicationSid)
	d.Set("voice_fallback_method", iPN.VoiceFallbackMethod)
	d.Set("voice_fallback_url", iPN.VoiceFallbackURL)
	d.Set("voice_method", iPN.VoiceMethod)
	d.Set("voice_url", iPN.VoiceURL)
	d.Set("emergency_status", iPN.EmergencyStatus)
	d.Set("emergency_address_sid", iPN.EmergencyAddressSid)
	d.Set("bundle_sid", iPN.BundleSid)

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
