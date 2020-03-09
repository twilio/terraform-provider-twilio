package twilio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/twilio/terraform-provider-twilio/util"
	types "github.com/twilio/twilio-go"
)

func dataSourceTaskRouterTaskQueues() *schema.Resource { // nolint:golint,funlen
	return &schema.Resource{
		Read: dataSourceTaskRouterTaskQueueRead,
		Schema: map[string]*schema.Schema{
			"workspace_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_queues": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: TaskRouterTaskQueueSchema(),
				},
			},
		},
	}
}

func dataSourceTaskRouterTaskQueueRead(d *schema.ResourceData, m interface{}) error {
	workspaceSID := d.Get("workspace_sid").(string)
	t, err := m.(*Config).Client.TaskRouter.TaskQueues.Read(workspaceSID, dataSourceTaskRouterTaskQueueQueryParams(d))
	if err != nil {
		return err
	}

	taskqueues := []interface{}{}
	for _, v := range t.TaskQueues {
		taskqueues = append(taskqueues, flattenTaskQueue(v))
	}

	d.Set("task_queues", taskqueues)

	return nil
}

func dataSourceTaskRouterTaskQueueQueryParams(d *schema.ResourceData) *types.TaskRouterTaskQueueQueryParams {
	t := &types.TaskRouterTaskQueueQueryParams{}
	if v, ok := d.GetOk("friendly_name"); ok {
		t.FriendlyName = util.String(v.(string))
	}

	return t
}

func flattenTaskQueue(t *types.TaskRouterTaskQueue) interface{} {
	task := map[string]interface{}{}
	task["account_sid"] = t.AccountSID
	task["assignment_activity_sid"] = t.AssignmentActivitySID
	task["assignment_activity_name"] = t.AssignmentActivityName
	task["date_created"] = t.DateCreated.String
	task["date_updated"] = t.DateUpdated.String
	task["friendly_name"] = t.FriendlyName
	task["max_reserved_workers"] = t.MaxReservedWorkers
	task["reservation_activity_sid"] = t.ReservationActivitySID
	task["reservation_activity_name"] = t.ReservationActivityName
	task["target_workers"] = t.TargetWorkers
	task["task_order"] = t.TaskOrder
	task["url"] = t.URL
	task["workspace_sid"] = t.WorkspaceSID
	task["links"] = t.Links

	return task
}
