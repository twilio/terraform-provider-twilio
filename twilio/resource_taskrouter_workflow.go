package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceTaskRouterWorkflow() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceTaskRouterWorkflowCreate,
		Read:   resourceTaskRouterWorkflowRead,
		Update: resourceTaskRouterWorkflowUpdate,
		Delete: resourceTaskRouterWorkflowDelete,
		Schema: map[string]*schema.Schema{
			"assignment_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"configuration": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"document_content_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fallback_assignment_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_reservation_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceTaskRouterWorkflowCreate(d *schema.ResourceData, m interface{}) error {
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.Workflows.Create(*workspaceSID,
		getWorkflowParams(d),
	)

	if err != nil {
		return fmt.Errorf("error creating workflow: %s", err)
	}

	d.SetId(*r.SID)

	return resourceTaskRouterWorkflowRead(d, m)
}

func resourceTaskRouterWorkflowRead(d *schema.ResourceData, m interface{}) error {
	workflowSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.Workflows.Fetch(*workspaceSID, workflowSID)

	if err != nil {
		return err
	}

	d.Set("assignment_callback_url", r.AssignmentCallbackURL)
	d.Set("AccountSid", r.AccountSID)
	d.Set("configuration", r.Configuration)
	d.Set("date_created", r.DateCreated)
	d.Set("date_updated", r.DateUpdated)
	d.Set("document_content_type", r.DocumentContentType)
	d.Set("fallback_assignment_callback_url", r.FallbackAssignmentCallbackURL)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("task_reservation_timeout", r.TaskReservationTimeout)
	d.Set("WorkspaceSid", r.WorkspaceSID)
	d.Set("url", r.URL)
	d.Set("links", r.Links)

	return nil
}

func resourceTaskRouterWorkflowUpdate(d *schema.ResourceData, m interface{}) error {
	workflowSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.Workflows.Update(*workspaceSID, workflowSID,
		getWorkflowParams(d),
	)

	if err != nil {
		return fmt.Errorf("error updating workflow: %s", err)
	}

	d.SetId(*r.SID)

	return resourceTaskRouterWorkflowRead(d, m)
}

func resourceTaskRouterWorkflowDelete(d *schema.ResourceData, m interface{}) error {
	workflowSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	if err := m.(*Config).Client.TaskRouter.Workflows.Delete(*workspaceSID, workflowSID); err != nil {
		return fmt.Errorf("error deleting workflow: %s", err)
	}

	return nil
}

func getWorkflowParams(d *schema.ResourceData) *types.TaskRouterWorkflowParams {
	params := &types.TaskRouterWorkflowParams{
		FriendlyName:  types.String(d.Get("friendly_name").(string)),
		Configuration: types.String(d.Get("configuration").(string)),
	}

	if v, exists := d.GetOk("assignment_callback_url"); exists {
		params.AssignmentCallbackURL = types.String((v).(string))
	}

	if v, exists := d.GetOk("task_reservation_timeout"); exists {
		params.TaskReservationTimeout = types.Int((v).(int))
	}

	if v, exists := d.GetOk("fallback_assignment_callback_url"); exists {
		params.AssignmentCallbackURL = types.String((v).(string))
	}

	return params
}