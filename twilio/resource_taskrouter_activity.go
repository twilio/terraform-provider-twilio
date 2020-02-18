package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceTaskRouterActivity() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceTaskRouterActivityCreate,
		Read:   resourceTaskRouterActivityRead,
		Update: resourceTaskRouterActivityUpdate,
		Delete: resourceTaskRouterActivityDelete,
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"available": {
				Type:     schema.TypeBool,
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

func resourceTaskRouterActivityCreate(d *schema.ResourceData, m interface{}) error {
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.Activities.Create(
		*workspaceSID,
		getActivityParams(d),
	)

	if err != nil {
		return fmt.Errorf("error creating activity: %s", err)
	}

	d.SetId(*r.SID)

	return resourceTaskRouterActivityRead(d, m)
}

func resourceTaskRouterActivityRead(d *schema.ResourceData, m interface{}) error {
	activitySID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.Activities.Fetch(*workspaceSID, activitySID)

	if err != nil {
		return err
	}

	d.Set("account_sid", r.AccountSID)
	d.Set("available", r.Available)
	d.Set("date_created", r.DateCreated)
	d.Set("date_updated", r.DateUpdated)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("workspace_sid", r.WorkspaceSID)
	d.Set("url", r.URL)

	return nil
}

func resourceTaskRouterActivityUpdate(d *schema.ResourceData, m interface{}) error {
	activitySID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.Activities.Update(
		*workspaceSID,
		activitySID,
		getActivityParams(d),
	)

	if err != nil {
		return fmt.Errorf("error updating activity: %s", err)
	}

	d.SetId(*r.SID)

	return resourceTaskRouterActivityRead(d, m)
}

func resourceTaskRouterActivityDelete(d *schema.ResourceData, m interface{}) error {
	activitySID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	if err := m.(*Config).Client.TaskRouter.Activities.Delete(*workspaceSID, activitySID); err != nil {
		return fmt.Errorf("error deleting activity: %s", err)
	}

	return nil
}

func getActivityParams(d *schema.ResourceData) *types.TaskRouterActivityParams {
	params := &types.TaskRouterActivityParams{
		FriendlyName: types.String(d.Get("friendly_name").(string)),
		Available:    types.Bool(d.Get("available").(bool)),
	}

	return params
}
