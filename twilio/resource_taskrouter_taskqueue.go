package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceTaskRouterTaskQueue() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceTaskRouterTaskQueueCreate,
		Read:   resourceTaskRouterTaskQueueRead,
		Update: resourceTaskRouterTaskQueueUpdate,
		Delete: resourceTaskRouterTaskQueueDelete,
		Schema: TaskRouterTaskQueueSchema(),
	}
}

// TaskRouterTaskQueueSchema returns the Terraform TaskQueue Schema.
func TaskRouterTaskQueueSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_sid": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"assignment_activity_sid": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"assignment_activity_name": {
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
		"friendly_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"max_reserved_workers": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"reservation_activity_sid": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"reservation_activity_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"target_workers": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"task_order": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"url": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"workspace_sid": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"links": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func resourceTaskRouterTaskQueueCreate(d *schema.ResourceData, m interface{}) error {
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.TaskQueues.Create(*workspaceSID, getTaskQueueParams(d))

	if err != nil {
		return fmt.Errorf("error creating taskqueue: %s", err)
	}

	d.SetId(*r.SID)

	return resourceTaskRouterTaskQueueRead(d, m)
}

func resourceTaskRouterTaskQueueRead(d *schema.ResourceData, m interface{}) error {
	taskQueueSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.TaskQueues.Fetch(*workspaceSID, taskQueueSID)

	if err != nil {
		return err
	}

	d.Set("account_sid", r.AccountSID)
	d.Set("assignment_activity_sid", r.AssignmentActivitySID)
	d.Set("assignment_activity_name", r.AssignmentActivityName)
	d.Set("date_created", r.DateCreated.String)
	d.Set("date_updated", r.DateUpdated.String)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("max_reserved_workers", r.MaxReservedWorkers)
	d.Set("reservation_activity_sid", r.ReservationActivitySID)
	d.Set("reservation_activity_name", r.ReservationActivityName)
	d.Set("target_workers", r.TargetWorkers)
	d.Set("task_order", r.TaskOrder)
	d.Set("url", r.URL)
	d.Set("workspace_sid", r.WorkspaceSID)
	d.Set("links", r.Links)

	return nil
}

func resourceTaskRouterTaskQueueUpdate(d *schema.ResourceData, m interface{}) error {
	taskQueueSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.TaskQueues.Update(
		*workspaceSID,
		taskQueueSID,
		getTaskQueueParams(d),
	)

	if err != nil {
		return fmt.Errorf("error updating taskqueue: %s", err)
	}

	d.SetId(*r.SID)

	return resourceTaskRouterTaskQueueRead(d, m)
}

func resourceTaskRouterTaskQueueDelete(d *schema.ResourceData, m interface{}) error {
	taskQueueSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	if err := m.(*Config).Client.TaskRouter.TaskQueues.Delete(*workspaceSID, taskQueueSID); err != nil {
		return fmt.Errorf("error deleting taskqueue: %s", err)
	}

	return nil
}

func getTaskQueueParams(d *schema.ResourceData) *types.TaskRouterTaskQueueParams {
	params := &types.TaskRouterTaskQueueParams{
		FriendlyName: types.String(d.Get("friendly_name").(string)),
	}

	if v, exists := d.GetOk("assignment_activity_sid"); exists {
		params.AssignmentActivitySID = types.String((v).(string))
	}

	if v, exists := d.GetOk("max_reserved_workers"); exists {
		params.MaxReservedWorkers = types.Int((v).(int))
	}

	if v, exists := d.GetOk("target_workers"); exists {
		params.TargetWorkers = types.String((v).(string))
	}

	if v, exists := d.GetOk("task_order"); exists {
		params.TaskOrder = types.String((v).(string))
	}

	if v, exists := d.GetOk("reservation_activity_sid"); exists {
		params.ReservationActivitySID = types.String((v).(string))
	}

	return params
}
