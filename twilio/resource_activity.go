package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceActivity() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceActivityCreate,
		Read:   resourceActivityRead,
		Update: resourceActivityUpdate,
		Delete: resourceActivityDelete,
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"available": {
				Type:     schema.TypeString,
				Optional: true,
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
				Required: true,
			},
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"workspace_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceActivityCreate(d *schema.ResourceData, m interface{}) error {
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.ActivityClient.Create(
		getActivityParams(d, m),
		*workspaceSID,
	)

	if err != nil {
		return fmt.Errorf("error creating activity: %s", err)
	}

	d.SetId(r.Sid)

	return resourceActivityRead(d, m)
}

func resourceActivityRead(d *schema.ResourceData, m interface{}) error {
	activitySID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.ActivityClient.Fetch(*workspaceSID, activitySID)

	if err != nil {
		return err
	}

	d.Set("account_sid", r.AccountSid)
	d.Set("available", r.Available)
	d.Set("date_created", r.DateCreated)
	d.Set("date_updated", r.DateUpdated)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("sid", r.Sid)
	d.Set("workspace_sid", r.WorkspaceSid)
	d.Set("url", r.URI)

	return nil
}

func resourceActivityUpdate(d *schema.ResourceData, m interface{}) error {
	activitySID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.ActivityClient.Update(
		getActivityParams(d, m),
		*workspaceSID,
		activitySID,
	)

	if err != nil {
		return fmt.Errorf("error updating activity: %s", err)
	}

	d.SetId(r.Sid)

	return resourceActivityRead(d, m)
}

func resourceActivityDelete(d *schema.ResourceData, m interface{}) error {
	activitySID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	if err := m.(*Config).Client.TaskRouter.ActivityClient.Delete(*workspaceSID, activitySID); err != nil {
		return fmt.Errorf("error deleting activity: %s", err)
	}

	return nil
}

func getActivityParams(d *schema.ResourceData, m interface{}) *types.ActivityParams {
	var available *string

	if v, exists := d.GetOk("available"); exists {
		available = types.String((v).(string))
	}

	params :=
		&types.ActivityParams{
			FriendlyName: *types.String(d.Get("friendly_name").(string)),
			Available:    available,
		}

	return params
}
