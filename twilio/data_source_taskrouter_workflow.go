package twilio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/twilio/terraform-provider-twilio/util"
	types "github.com/twilio/twilio-go"
)

func dataSourceTaskRouterWorkflows() *schema.Resource { // nolint:golint,funlen
	return &schema.Resource{
		Read: dataSourceTaskRouterWorkflowRead,
		Schema: map[string]*schema.Schema{
			"workspace_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workflows": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: TaskRouterTaskQueueSchema(),
				},
			},
		},
	}
}

func dataSourceTaskRouterWorkflowRead(d *schema.ResourceData, m interface{}) error {
	workspaceSID := d.Get("workspace_sid").(string)
	t, err := m.(*Config).Client.TaskRouter.Workflows.Read(workspaceSID, dataSourceTaskRouterWorkflowQueryParams(d))
	if err != nil {
		return err
	}

	workflows := []interface{}{}
	for _, v := range t.Workflows {
		workflows = append(workflows, flattenWorkflow(v))
	}

	d.Set("workflows", workflows)

	return nil
}

func dataSourceTaskRouterWorkflowQueryParams(d *schema.ResourceData) *types.TaskRouterWorkflowQueryParams {
	t := &types.TaskRouterWorkflowQueryParams{}
	if v, ok := d.GetOk("friendly_name"); ok {
		t.FriendlyName = util.String(v.(string))
	}

	return t
}

func flattenWorkflow(w *types.TaskRouterWorkflow) interface{} {
	workflow := map[string]interface{}{}
	workflow["assignment_callback_url"] = w.AssignmentCallbackURL
	workflow["account_sid"] = w.AccountSID
	workflow["configuration"] = w.Configuration
	workflow["date_created"] = w.DateCreated
	workflow["date_updated"] = w.DateUpdated
	workflow["document_content_type"] = w.DocumentContentType
	workflow["fallback_assignment_callback_url"] = w.FallbackAssignmentCallbackURL
	workflow["friendly_name"] = w.FriendlyName
	workflow["task_reservation_timeout"] = w.TaskReservationTimeout
	workflow["workspace_sid"] = w.WorkspaceSID
	workflow["url"] = w.URL
	workflow["links"] = w.Links

	return workflow
}
