package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceTaskQueue() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceTaskQueueCreate,
		Read:   resourceTaskQueueRead,
		Update: resourceTaskQueueUpdate,
		Delete: resourceTaskQueueDelete,
		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourceTaskQueueCreate(d *schema.ResourceData, m interface{}) error {
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.TaskQueueClient.Create(*workspaceSID, getTaskQueueParams(d))

	if err != nil {
		return fmt.Errorf("error creating taskqueue: %s", err)
	}

	d.SetId(r.Sid)

	return resourceTaskQueueRead(d, m)
}

func resourceTaskQueueRead(d *schema.ResourceData, m interface{}) error {
	taskQueueSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.TaskQueueClient.Fetch(*workspaceSID, taskQueueSID)

	if err != nil {
		return err
	}

	d.Set("account_sid", r.AccountSid)
	d.Set("assignment_activity_sid", r.AssignmentActivitySid)
	d.Set("assignment_activity_name", r.AssignmentActivityName)
	d.Set("date_created", r.DateCreated.String)
	d.Set("date_updated", r.DateUpdated.String)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("max_reserved_workers", r.MaxReservedWorkers)
	d.Set("reservation_activity_sid", r.ReservationActivitySid)
	d.Set("reservation_activity_name", r.ReservationActivityName)
	d.Set("target_workers", r.TargetWorkers)
	d.Set("task_order", r.TaskOrder)
	d.Set("url", r.URI)
	d.Set("workspace_sid", r.WorkspaceSid)
	d.Set("links", r.Links)

	return nil
}

func resourceTaskQueueUpdate(d *schema.ResourceData, m interface{}) error {
	taskQueueSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.TaskQueueClient.Update(
		*workspaceSID,
		taskQueueSID,
		getTaskQueueParams(d),
	)

	if err != nil {
		return fmt.Errorf("error updating taskqueue: %s", err)
	}

	d.SetId(r.Sid)

	return resourceTaskQueueRead(d, m)
}

func resourceTaskQueueDelete(d *schema.ResourceData, m interface{}) error {
	taskQueueSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	if err := m.(*Config).Client.TaskRouter.TaskQueueClient.Delete(*workspaceSID, taskQueueSID); err != nil {
		return fmt.Errorf("error deleting taskqueue: %s", err)
	}

	return nil
}

func getTaskQueueParams(d *schema.ResourceData) *types.TaskQueueParams {
	var assignmentActivitySid *string

	var maxReservedWorkers *int

	var targetWorkers *string

	var taskOrder *string

	var reservationActivitySid *string

	if v, exists := d.GetOk("assignment_activity_sid"); exists {
		assignmentActivitySid = types.String((v).(string))
	}

	if v, exists := d.GetOk("max_reserved_workers"); exists {
		maxReservedWorkers = types.Int((v).(int))
	}

	if v, exists := d.GetOk("target_workers"); exists {
		targetWorkers = types.String((v).(string))
	}

	if v, exists := d.GetOk("task_order"); exists {
		taskOrder = types.String((v).(string))
	}

	if v, exists := d.GetOk("reservation_activity_sid"); exists {
		reservationActivitySid = types.String((v).(string))
	}

	params := &types.TaskQueueParams{
		FriendlyName:           *types.String(d.Get("friendly_name").(string)),
		AssignmentActivitySid:  assignmentActivitySid,
		MaxReservedWorkers:     maxReservedWorkers,
		TargetWorkers:          targetWorkers,
		TaskOrder:              taskOrder,
		ReservationActivitySid: reservationActivitySid,
	}

	return params
}
