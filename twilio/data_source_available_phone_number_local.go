package twilio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/twilio/terraform-provider-twilio/util"
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

func dataSourceAvailablePhoneNumbersLocal() *schema.Resource { // nolint:golint,funlen
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
	a, err := m.(*Config).Client.AvailablePhoneNumbers.Read(params)

	if err != nil {
		return err
	}

	d.SetId(*a.URI)
	d.Set("num_pages", a.NumPages)
	d.Set("page", a.Page)
	d.Set("page_size", a.PageSize)
	d.Set("start", a.Start)
	d.Set("total", a.Total)
	d.Set("end", a.End)
	d.Set("uri", a.URI)
	d.Set("first_page_uri", a.FirstPageURI)
	d.Set("last_page_uri", a.LastPageURI)
	d.Set("next_page_uri", a.NextPageURI)
	d.Set("previous_page_uri", a.PreviousPageURI)

	availablePhoneNumbers := make([]interface{}, 0, len(a.AvailablePhoneNumbers))

	for _, v := range a.AvailablePhoneNumbers {
		availablePhoneNumbers = append(availablePhoneNumbers, flattenPhoneNumber(v))
	}

	d.Set("available_phone_numbers", availablePhoneNumbers)

	return nil
}

func dataSourceAvailablePhoneNumberLocalParams(d *schema.ResourceData) *types.AvailablePhoneNumberLocalReadParams {
	p := new(types.AvailablePhoneNumberLocalReadParams)

	if v, ok := d.GetOk("fax_enabled"); ok {
		p.FaxEnabled = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("sms_enabled"); ok {
		p.SMSEnabled = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("mms_enabled"); ok {
		p.MMSEnabled = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("voice_enabled"); ok {
		p.VoiceEnabled = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("exclude_all_address_required"); ok {
		p.ExcludeAllAddressRequired = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("exclude_local_address_required"); ok {
		p.ExcludeLocalAddressRequired = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("exclude_foreign_address_required"); ok {
		p.ExcludeForeignAddressRequired = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("beta"); ok {
		p.Beta = util.Bool(v.(bool))
	}

	if v, ok := d.GetOk("distance"); ok {
		p.Distance = util.Int(v.(int))
	}

	if v, ok := d.GetOk("area_code"); ok {
		p.AreaCode = util.Int(v.(int))
	}

	if v, ok := d.GetOk("in_postal_code"); ok {
		p.InPostalCode = util.String(v.(string))
	}

	if v, ok := d.GetOk("near_number"); ok {
		p.NearNumber = util.String(v.(string))
	}

	if v, ok := d.GetOk("near_lat_long"); ok {
		p.NearLatLong = util.String(v.(string))
	}

	if v, ok := d.GetOk("contains"); ok {
		p.Contains = util.String(v.(string))
	}

	if v, ok := d.GetOk("in_region"); ok {
		p.InRegion = util.String(v.(string))
	}

	if v, ok := d.GetOk("in_rate_center"); ok {
		p.InRateCenter = util.String(v.(string))
	}

	if v, ok := d.GetOk("in_lata"); ok {
		p.InLATA = util.String(v.(string))
	}

	if v, ok := d.GetOk("in_locality"); ok {
		p.InLocality = util.String(v.(string))
	}

	return p
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
