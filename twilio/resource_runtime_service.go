package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/twilio/terraform-provider-twilio/util"
	types "github.com/twilio/twilio-go"
)

func resourceRuntimeService() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceRuntimeServiceCreate,
		Read:   resourceRuntimeServiceRead,
		Update: resourceRuntimeServiceUpdate,
		Delete: resourceRuntimeServiceDelete,
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"include_credentials": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRuntimeServiceCreate(d *schema.ResourceData, m interface{}) error {

	r, err := m.(*Config).Client.Serverless.Service.Create(resourceRuntimeServiceParams(d))

	if err != nil {
		return fmt.Errorf("error creating Runtime Service: %s", err)
	}

	d.SetId(*r.SID)

	return resourceRuntimeServiceRead(d, m)
}

func resourceRuntimeServiceRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	r, err := m.(*Config).Client.Serverless.Service.Fetch(id)

	d.Set("account_sid", r.AccountSID)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("unique_name", r.UniqueName)
	d.Set("include_credentials", r.IncludeCredentials)

	if err != nil {
		return err
	}

	return nil
}

func resourceRuntimeServiceUpdate(d *schema.ResourceData, m interface{}) error {
	if _, err := m.(*Config).Client.Serverless.Service.Update(d.Id(), resourceRuntimeServiceParams(d)); err != nil {
		return fmt.Errorf("error updating Runtime Service: %s", err)
	}

	return resourceRuntimeServiceRead(d, m)
}

func resourceRuntimeServiceDelete(d *schema.ResourceData, m interface{}) error {
	if err := m.(*Config).Client.Serverless.Service.Delete(d.Id()); err != nil {
		return fmt.Errorf("error deleting Runtime Service: %s", err)
	}

	return nil
}

func resourceRuntimeServiceParams(d *schema.ResourceData) *types.RuntimeServiceParams {
	p := &types.RuntimeServiceParams{}

	if v, ok := d.GetOk("friendly_name"); ok {
		p.FriendlyName = util.String(v.(string))
	}

	if v, ok := d.GetOk("unique_name"); ok {
		p.UniqueName = util.String(v.(string))
	}

	if v, ok := d.GetOk("included_credentials"); ok {
		p.IncludeCredentials = util.Bool(v.(bool))
	}

	return p
}
