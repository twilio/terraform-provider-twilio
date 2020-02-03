package twilio

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	types "github.com/twilio/twilio-go"
)

func resourceWorkflow() *schema.Resource { //nolint:golint,funlen
	return &schema.Resource{
		Create: resourceWorkflowCreate,
		Read:   resourceWorkflowRead,
		Update: resourceWorkflowUpdate,
		Delete: resourceWorkflowDelete,
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
				ForceNew: true,
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
			"sid": {
				Type:     schema.TypeString,
				ForceNew: true,
			},
			"task_reservation_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceWorkflowCreate(d *schema.ResourceData, m interface{}) error {
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.WorkflowClient.Create(*workspaceSID,
		&types.WorkflowParams{
			FriendlyName:                  *types.String(d.Get("friendly_name").(string)),
			AssignmentCallbackURL:         types.String(d.Get("assignment_callback_url").(string)),
			Configuration:                 *types.String(d.Get("configuration").(string)),
			TaskReservationTimeout:        types.Int(d.Get("task_reservation_timeout").(int)),
			FallbackAssignmentCallbackURL: types.String(d.Get("fallback_assignment_callback_url").(string)),
		},
	)

	if err != nil {
		return fmt.Errorf("error creating workflow: %s", err)
	}

	d.SetId(r.Sid)

	return resourceWorkflowRead(d, m)
}

func resourceWorkflowRead(d *schema.ResourceData, m interface{}) error {
	workflowSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.WorkflowClient.Fetch(*workspaceSID, workflowSID)

	if err != nil {
		return err
	}

	d.Set("assignment_callback_url", r.AssignmentCallbackURL)
	d.Set("AccountSid", r.AccountSid)
	d.Set("configuration", r.Configuration)
	d.Set("date_created", r.DateCreated)
	d.Set("date_updated", r.DateUpdated)
	d.Set("document_content_type", r.DocumentContentType)
	d.Set("fallback_assignment_callback_url", r.FallbackAssignmentCallbackURL)
	d.Set("friendly_name", r.FriendlyName)
	d.Set("sid", r.Sid)
	d.Set("task_reservation_timeout", r.TaskReservationTimeout)
	d.Set("WorkspaceSid", r.WorkspaceSid)
	d.Set("url", r.URL)
	d.Set("links", r.Links)

	return nil
}

func resourceWorkflowUpdate(d *schema.ResourceData, m interface{}) error {
	workflowSID := d.Id()
	workspaceSID := types.String(d.Get("workspace_sid").(string))

	r, err := m.(*Config).Client.TaskRouter.WorkflowClient.Update(*workspaceSID,
		&types.WorkflowParams{
			FriendlyName:                  *types.String(d.Get("friendly_name").(string)),
			AssignmentCallbackURL:         types.String(d.Get("assignment_callback_url").(string)),
			Configuration:                 *types.String(d.Get("configuration").(string)),
			TaskReservationTimeout:        types.Int(d.Get("task_reservation_timeout").(int)),
			FallbackAssignmentCallbackURL: types.String(d.Get("fallback_assignment_callback_url").(string)),
		}, workflowSID,
	)

	if err != nil {
		return fmt.Errorf("error updating workspace: %s", err)
	}

	d.SetId(r.Sid)

	return resourceWorkflowRead(d, m)
}

func resourceWorkflowDelete(d *schema.ResourceData, m interface{}) error {
	workspaceSID := d.Id()

	if err := m.(*Config).Client.TaskRouter.WorkspaceClient.Delete(workspaceSID); err != nil {
		return fmt.Errorf("error deleting workspace: %s", err)
	}

	return nil
}
