package twilio

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/twilio/terraform-provider-twilio/util"
	types "github.com/twilio/twilio-go"
)

func resourceStudioFlow() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceStudioFlowCreate,
		Read:   resourceStudioFlowRead,
		Update: resourceStudioFlowUpdate,
		Delete: resourceStudioFlowDelete,
		Schema: map[string]*schema.Schema{
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Required: true,
			},
			"definition": {
				Type:     schema.TypeString,
				Required: true,
			},
			"commit_message": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceStudioFlowCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("STRUCT: %+v", resourceStudioFlowParams(d))
	r, err := m.(*Config).Client.Studio.Flow.Create(resourceStudioFlowParams(d))

	if err != nil {
		return fmt.Errorf("error creating Proxy Phone Number: %s", err)
	}

	d.SetId(*r.SID)

	return resourceStudioFlowRead(d, m)
}

func resourceStudioFlowRead(d *schema.ResourceData, m interface{}) error {
	r, err := m.(*Config).Client.Studio.Flow.Read(d.Id())

	d.Set("account_sid", r.AccountSID)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("definition", r.Definition)
	d.Set("status", r.Status)
	d.Set("commit_message", r.CommitMessage)

	if err != nil {
		return err
	}

	return nil
}

func resourceStudioFlowUpdate(d *schema.ResourceData, m interface{}) error {
	if _, err := m.(*Config).Client.Studio.Flow.Update(d.Id(), resourceStudioFlowParams(d)); err != nil {
		return fmt.Errorf("error updating Proxy Phone Number: %s", err)
	}

	return resourceStudioFlowRead(d, m)
}

func resourceStudioFlowDelete(d *schema.ResourceData, m interface{}) error {
	if err := m.(*Config).Client.Studio.Flow.Delete(d.Id()); err != nil {
		return fmt.Errorf("error deleting Proxy Phone Number: %s", err)
	}

	return nil
}

func resourceStudioFlowParams(d *schema.ResourceData) *types.StudioFlowParams {
	p := &types.StudioFlowParams{
		FriendlyName: types.String(d.Get("friendly_name").(string)),
		Status:       types.String(d.Get("status").(string)),
		Definition:   types.String(d.Get("definition").(string)),
	}

	if v, ok := d.GetOk("commit_message"); ok {
		p.CommitMessage = util.String(v.(string))
	}

	return p
}
