package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/twilio/terraform-provider-twilio/util"
	types "github.com/twilio/twilio-go"
)

func resourceRuntimeEnvironment() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceRuntimeEnvironmentCreate,
		Read:   resourceRuntimeEnvironmentRead,
		Delete: resourceRuntimeEnvironmentDelete,
		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_suffix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"build_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceRuntimeEnvironmentCreate(d *schema.ResourceData, m interface{}) error {
	serviceSID := d.Get("service_sid").(string)
	r, err := m.(*Config).Client.Serverless.Environment.Create(serviceSID, resourceRuntimeEnvironmentParams(d))

	if err != nil {
		return fmt.Errorf("error creating Runtime Environment: %s", err)
	}

	d.SetId(*r.SID)

	return resourceRuntimeEnvironmentRead(d, m)
}

func resourceRuntimeEnvironmentRead(d *schema.ResourceData, m interface{}) error {
	serviceSID := d.Get("service_sid").(string)
	id := d.Id()

	r, err := m.(*Config).Client.Serverless.Environment.Fetch(serviceSID, id)

	d.Set("service_sid", r.ServiceSID)
	d.Set("domain_suffix", r.DomainSuffix)
	d.Set("unique_name", r.UniqueName)
	d.Set("domain_name", r.DomainName)
	d.Set("build_sid", r.BuildSID)

	if err != nil {
		return err
	}

	return nil
}

func resourceRuntimeEnvironmentDelete(d *schema.ResourceData, m interface{}) error {
	serviceSID := d.Get("service_sid").(string)
	if err := m.(*Config).Client.Serverless.Environment.Delete(serviceSID, d.Id()); err != nil {
		return fmt.Errorf("error deleting Runtime Environment: %s", err)
	}

	return nil
}

func resourceRuntimeEnvironmentParams(d *schema.ResourceData) *types.RuntimeEnvironmentParams {
	p := &types.RuntimeEnvironmentParams{}

	if v, ok := d.GetOk("domain_suffix"); ok {
		p.DomainSuffix = util.String(v.(string))
	}

	if v, ok := d.GetOk("unique_name"); ok {
		p.UniqueName = util.String(v.(string))
	}

	return p
}
