package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceIncomingPhoneNumber() *schema.Resource {
	return &schema.Resource{
		Create: resourceIncomingPhoneNumberCreate,
		Read:   resourceIncomingPhoneNumberRead,
		Update: resourceIncomingPhoneNumberUpdate,
		Delete: resourceIncomingPhoneNumberDelete,
		Schema: map[string]*schema.Schema{
			"beta": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"voice_caller_id_lookup": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_requirements": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"api_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"capabilities": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
			"date_created": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"date_updated": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identity_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sid": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
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
		AccountSid:           d.Get("accound_sid").(string),
		APIVersion:           d.Get("api_version").(string),
		FriendlyName:         d.Get("friendly_name").(string),
		SMSApplicationSid:    d.Get("sms_application_sid").(string),
		SMSFallbackMethod:    d.Get("sms_fallback_method").(string),
		PhoneNumber:          d.Get("phone_number").(string),
		AreaCode:             d.Get("area_code").(string),
		SMSFallbackURL:       d.Get("sms_fallback_url").(string),
		SMSMethod:            d.Get("sms_method").(string),
		SMSURL:               d.Get("sms_url").(string),
		StatusCallback:       d.Get("status_callback").(string),
		AddressSid:           d.Get("address_sid").(string),
		StatusCallbackMethod: d.Get("status_callback_method").(string),
		VoiceApplicationSid:  d.Get("voice_application_sid").(string),
		VoiceCallerIDLookup:  d.Get("voice_called_id_lookup").(bool),
		VoiceFallbackMethod:  d.Get("voice_fallback_method").(string),
		VoiceFallbackURL:     d.Get("voice_fallback_url").(string),
		VoiceMethod:          d.Get("voice_method").(string),
		VoiceURL:             d.Get("voice_url").(string),
		VoiceReceiveMode:     d.Get("voice_receive_mode").(string),
		EmergencyStatus:      d.Get("emergency_status").(string),
		EmergencyAddressSid:  d.Get("emergency_address_sid").(string),
		BundleSid:            d.Get("bundle_sid").(string),
		IdentitySid:          d.Get("identity_sid").(string),
		TrunkSid:             d.Get("trunk_sid").(string),
	}
}

func resourceIncomingPhoneNumberCreate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	iPNParams := &types.IncomingPhoneNumberParams{PhoneNumber: d.Get("phone_number").(string)}

	iPN, err := m.(*Config).Client.IncomingPhoneNumbers.Create(iPNParams)
	if err != nil {
		return fmt.Errorf("error creating Phone Number %s", err)

	}

	d.SetId(iPN.Sid)
	d.SetPartial("friendly_name")

	return nil
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
	if err := m.(*Config).Client.Chat.Service.Delete(d.Id()); err != nil {
		return fmt.Errorf("error deleting Incoming Phone Number: %s", err)
	}

	return nil
}
