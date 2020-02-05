package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceWorkspace() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceWorkspaceCreate,
		Read:   resourceWorkspaceRead,
		Update: resourceWorkspaceUpdate,
		Delete: resourceWorkspaceDelete,
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_activity_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_activity_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"links": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"event_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"events_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multi_task_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"template": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"prioritize_queue_order": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWorkspaceCreate(d *schema.ResourceData, m interface{}) error {
	r, err := m.(*Config).Client.TaskRouter.WorkspaceClient.Create(
		getWorkspaceParams(d),
	)

	if err != nil {
		return fmt.Errorf("error creating workspace: %s", err)
	}

	d.SetId(r.Sid)

	return resourceWorkspaceRead(d, m)
}

func resourceWorkspaceRead(d *schema.ResourceData, m interface{}) error {
	workspaceSID := d.Id()
	r, err := m.(*Config).Client.TaskRouter.WorkspaceClient.Fetch(workspaceSID)

	if err != nil {
		return err
	}

	d.Set("account_sid", r.AccountSid)
	d.Set("date_created", r.DateCreated.String)
	d.Set("date_updated", r.DateUpdated.String)
	d.Set("default_activity_name", r.DefaultActivityName)
	d.Set("default_activity_sid", r.DefaultActivitySid)
	d.Set("url", r.URL)
	d.Set("links", r.Links)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("event_callback_url", r.EventCallbackURL)
	d.Set("events_filter", r.EventsFilter)
	d.Set("multi_task_enabled", r.MultitaskEnabled)
	d.Set("template", r.Template)
	d.Set("prioritize_queue_order", r.PrioritizeQueueOrder)

	return nil
}

func resourceWorkspaceUpdate(d *schema.ResourceData, m interface{}) error {
	workspaceSID := d.Id()

	if _, err := m.(*Config).Client.TaskRouter.WorkspaceClient.Update(workspaceSID, getWorkspaceParams(d)); err != nil {
		return fmt.Errorf("error updating workspace: %s", err)
	}

	return resourceWorkspaceRead(d, m)
}

func resourceWorkspaceDelete(d *schema.ResourceData, m interface{}) error {
	workspaceSID := d.Id()

	if err := m.(*Config).Client.TaskRouter.WorkspaceClient.Delete(workspaceSID); err != nil {
		return fmt.Errorf("error deleting workspace: %s", err)
	}

	return nil
}

func getWorkspaceParams(d *schema.ResourceData) *types.WorkspaceParams {
	var eventCallBackURL *string

	var eventsFilter *string

	var multitaskEnabled *bool

	var prioritizeQueueOrder *string

	if v, exists := d.GetOk("event_callback_url"); exists {
		eventCallBackURL = types.String((v).(string))
	}

	if v, exists := d.GetOk("events_filter"); exists {
		eventsFilter = types.String((v).(string))
	}

	if v, exists := d.GetOk("multi_task_enabled"); exists {
		multitaskEnabled = types.Bool((v).(bool))
	}

	if v, exists := d.GetOk("prioritize_queue_order"); exists {
		prioritizeQueueOrder = types.String((v).(string))
	}

	params :=
		&types.WorkspaceParams{
			FriendlyName:         *types.String(d.Get("friendly_name").(string)),
			EventCallbackURL:     eventCallBackURL,
			EventsFilter:         eventsFilter,
			MultitaskEnabled:     multitaskEnabled,
			PrioritizeQueueOrder: prioritizeQueueOrder,
		}

	return params
}
