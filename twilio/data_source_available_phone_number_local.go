package twilio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceAvailablePhoneNumberLocal() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lata": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"locality": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_center": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  true,
			},
			"latitude": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"longitude": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"postal_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"iso_country": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_requirements": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"beta": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"capabilities": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
		},
	}
}

func dataSourceAvailablePhoneNumbersLocal() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAvailablePhoneNumbersLocalRead,
		Schema: map[string]*schema.Schema{
			"num_pages": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"page": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"page_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"start": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"end": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"first_page_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_page_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"next_page_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"previous_page_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"available_phone_number": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem:     resourceAvailablePhoneNumberLocal(),
			},
		},
	}
}

func dataSourceAvailablePhoneNumbersLocalRead(d *schema.ResourceData, m interface{}) error {
	params := dataSourceAvailablePhoneNumberLocalParams(d)
	availablePhoneNumbersLocal, err := m.(*Config).Client.AvailablePhoneNumbers.ReadMultiple(params)

	if err != nil {
		return err
	}

	d.Set("num_pages", availablePhoneNumbersLocal.NumPages)
	d.Set("page", availablePhoneNumbersLocal.Page)
	d.Set("page_size", availablePhoneNumbersLocal.PageSize)
	d.Set("start", availablePhoneNumbersLocal.Start)
	d.Set("total", availablePhoneNumbersLocal.Total)
	d.Set("end", availablePhoneNumbersLocal.End)
	d.Set("first_page_uri", availablePhoneNumbersLocal.FirstPageURI)
	d.Set("last_page_uri", availablePhoneNumbersLocal.LastPageURI)
	d.Set("next_page_uri", availablePhoneNumbersLocal.NextPageURI)
	d.Set("uri", availablePhoneNumbersLocal.URI)
	d.Set("previous_page_uri", availablePhoneNumbersLocal.PreviousPageURI)
	d.Set("available_phone_numbers", availablePhoneNumbersLocal.AvailablePhoneNumbers)

	return nil
}

func dataSourceAvailablePhoneNumberLocalParams(d *schema.ResourceData) *types.AvailablePhoneNumberLocalReadParams {
	return &types.AvailablePhoneNumberLocalReadParams{
		FaxEnabled:                    d.Get("fax_enabled").(bool),
		SMSEnabled:                    d.Get("sms_enabled").(bool),
		MMSEnabled:                    d.Get("mms_enabled").(bool),
		VoiceEnabled:                  d.Get("voice_enabled").(bool),
		ExcludeAllAddressRequired:     d.Get("exclude_all_address_required").(bool),
		ExcludeLocalAddressRequired:   d.Get("exclude_all_local_address_required").(bool),
		ExcludeForeignAddressRequired: d.Get("exclude_all_foreign_address_required").(bool),
		Beta:                          d.Get("beta").(bool),
		Distance:                      d.Get("distance").(int),
		AreaCode:                      d.Get("area_code").(int),
		InPostalCode:                  d.Get("postal_code").(string),
		NearNumber:                    d.Get("near_number").(string),
		NearLatLong:                   d.Get("near_lat_long").(string),
		Contains:                      d.Get("contains").(string),
		InRegion:                      d.Get("in_region").(string),
		InRateCenter:                  d.Get("in_rate_center").(string),
		InLATA:                        d.Get("in_lata").(string),
		InLocality:                    d.Get("in_locality").(string),
	}
}
