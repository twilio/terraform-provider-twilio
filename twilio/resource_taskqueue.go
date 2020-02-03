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
				ForceNew: true,
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
			},
			"reservation_activity_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reservation_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"target_workers": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_order": {
				Type:     schema.TypeString,
				Optional: true,
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

	r, err := m.(*Config).Client.TaskRouter.TaskQueueClient.Create(*workspaceSID,
		&types.TaskQueueParams{
			FriendlyName:           *types.String(d.Get("friendly_name").(string)),
			AssignmentActivitySid:  types.String(d.Get("assignment_activity_sid").(string)),
			MaxReservedWorkers:     types.Int(d.Get("max_reserved_workers").(int)),
			TargetWorkers:          types.String(d.Get("target_workers").(string)),
			TaskOrder:              types.String(d.Get("task_order").(string)),
			ReservationActivitySid: types.String(d.Get("reservation_activity_sid").(string)),
		},
	)

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
	d.Set("date_created", r.DateCreated)
	d.Set("date_updated", r.DateUpdated)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("max_reserved_workers", r.MaxReservedWorkers)
	d.Set("reservation_activity_sid", r.ReservationActivitySid)
	d.Set("reservation_activity_name", r.ReservationActivityName)
	d.Set("sid", r.Sid)
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
		&types.TaskQueueParams{
			FriendlyName:           *types.String(d.Get("friendly_name").(string)),
			AssignmentActivitySid:  types.String(d.Get("assignment_activity_sid").(string)),
			MaxReservedWorkers:     types.Int(d.Get("max_reserved_workers").(int)),
			TargetWorkers:          types.String(d.Get("target_workers").(string)),
			TaskOrder:              types.String(d.Get("task_order").(string)),
			ReservationActivitySid: types.String(d.Get("reservation_activity_sid").(string)),
		},
		*workspaceSID,
		taskQueueSID,
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
