package twilio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func dataSourceAvailablePhoneNumberLocal() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lata": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"locality": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rate_center": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latitude": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"longitude": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"postal_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso_country": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_requirements": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"beta": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"capabilities": {
				Type:     schema.TypeMap,
				Computed: true,
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
			"country_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"area_code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mms_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"voice_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"exclude_all_address_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"exclude_local_address_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"exclude_foreign_address_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"beta": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"near_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"near_lat_long": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"distance": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"in_postal_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_rate_center": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_lata": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_locality": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fax_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"num_pages": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"page": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"page_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"start": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"end": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"first_page_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_page_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"next_page_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"previous_page_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"available_phone_numbers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     dataSourceAvailablePhoneNumberLocal(),
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
	d.SetId(availablePhoneNumbersLocal.URI)
	d.Set("num_pages", availablePhoneNumbersLocal.NumPages)
	d.Set("page", availablePhoneNumbersLocal.Page)
	d.Set("page_size", availablePhoneNumbersLocal.PageSize)
	d.Set("start", availablePhoneNumbersLocal.Start)
	d.Set("total", availablePhoneNumbersLocal.Total)
	d.Set("end", availablePhoneNumbersLocal.End)
	d.Set("uri", availablePhoneNumbersLocal.URI)
	d.Set("first_page_uri", availablePhoneNumbersLocal.FirstPageURI)
	d.Set("last_page_uri", availablePhoneNumbersLocal.LastPageURI)
	d.Set("next_page_uri", availablePhoneNumbersLocal.NextPageURI)
	d.Set("previous_page_uri", availablePhoneNumbersLocal.PreviousPageURI)

	var availablePhoneNumbers []interface{}
	for _, v := range availablePhoneNumbersLocal.AvailablePhoneNumbers {
		availablePhoneNumbers = append(availablePhoneNumbers, flattenPhoneNumber(v))
	}

	d.Set("available_phone_numbers", availablePhoneNumbers)
	return nil
}

func dataSourceAvailablePhoneNumberLocalParams(d *schema.ResourceData) *types.AvailablePhoneNumberLocalReadParams {
	return &types.AvailablePhoneNumberLocalReadParams{
		FaxEnabled:                    d.Get("fax_enabled").(bool),
		SMSEnabled:                    d.Get("sms_enabled").(bool),
		MMSEnabled:                    d.Get("mms_enabled").(bool),
		VoiceEnabled:                  d.Get("voice_enabled").(bool),
		ExcludeAllAddressRequired:     d.Get("exclude_all_address_required").(bool),
		ExcludeLocalAddressRequired:   d.Get("exclude_local_address_required").(bool),
		ExcludeForeignAddressRequired: d.Get("exclude_foreign_address_required").(bool),
		Beta:                          d.Get("beta").(bool),
		Distance:                      d.Get("distance").(int),
		AreaCode:                      d.Get("area_code").(int),
		InPostalCode:                  d.Get("in_postal_code").(string),
		NearNumber:                    d.Get("near_number").(string),
		NearLatLong:                   d.Get("near_lat_long").(string),
		Contains:                      d.Get("contains").(string),
		InRegion:                      d.Get("in_region").(string),
		InRateCenter:                  d.Get("in_rate_center").(string),
		InLATA:                        d.Get("in_lata").(string),
		InLocality:                    d.Get("in_locality").(string),
	}
}

func flattenPhoneNumber(p *types.AvailablePhoneNumberLocal) interface{} {
	values := map[string]interface{}{}
	values["phone_number"] = p.PhoneNumber
	values["friendly_name"] = p.FriendlyName
	values["address_requirements"] = p.AddressRequirements
	values["beta"] = p.Beta
	values["iso_country"] = p.ISOCountry
	values["lata"] = p.LATA
	values["locality"] = p.Locality
	values["latitude"] = p.Latitude
	values["longitude"] = p.Longitude
	values["postal_code"] = p.PostalCode
	values["rate_center"] = p.RateCenter
	values["region"] = p.Region
	return values
}
