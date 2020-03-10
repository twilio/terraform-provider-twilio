package twilio

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	client "github.com/twilio/twilio-go"
)

func TestAccTwilioTaskRouterWorkspace_basic(t *testing.T) {
	var workspace client.TaskRouterWorkspace

	rn := "twilio_taskrouter_workspace.foo"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTwilioTaskRouterWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioTaskRouterWorkspaceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceExists(rn, &workspace),
					testAccCheckTwilioTaskRouterWorkspaceAttributes(&workspace, &testAccTwilioTaskRouterWorkspaceExpectedAttributes{
						FriendlyName:         "alex",
						EventCallbackURL:     "https://workspace-example.free.beeceptor.com",
						EventsFilter:         "task.created",
						MultitaskEnabled:     true,
						PrioritizeQueueOrder: "LIFO",
						Template:             "FIFO",
					}),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkspaceUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceExists(rn, &workspace),
					testAccCheckTwilioTaskRouterWorkspaceAttributes(&workspace, &testAccTwilioTaskRouterWorkspaceExpectedAttributes{
						FriendlyName:         "will",
						EventCallbackURL:     "https://workspace-example.free.beeceptor.com",
						EventsFilter:         "task.created",
						MultitaskEnabled:     true,
						PrioritizeQueueOrder: "LIFO",
						Template:             "FIFO",
					}),
				),
			},
		},
	})
}

func testAccCheckTwilioTaskRouterWorkspaceExists(n string, workspace *client.TaskRouterWorkspace) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("workspace not found in state")
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("workspace id has not been set")
		}

		client := testAccProvider.Meta().(*Config)

		ws, err := client.Client.TaskRouter.Workspaces.Fetch(rs.Primary.ID)

		if err != nil {
			return err
		}

		if *ws.SID != rs.Primary.ID {
			return fmt.Errorf("workspace not found on fetch")
		}

		*workspace = *ws

		return nil
	}
}

type testAccTwilioTaskRouterWorkspaceExpectedAttributes struct {
	FriendlyName         string
	EventCallbackURL     string
	EventsFilter         string
	MultitaskEnabled     bool
	PrioritizeQueueOrder string
	Template             string
}

func testAccCheckTwilioTaskRouterWorkspaceAttributes(workspace *client.TaskRouterWorkspace, want *testAccTwilioTaskRouterWorkspaceExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if *workspace.FriendlyName != want.FriendlyName {
			return fmt.Errorf("got workspace friendly_name %s; want %s", *workspace.FriendlyName, want.FriendlyName)
		}

		if *workspace.EventCallbackURL != want.EventCallbackURL {
			return fmt.Errorf("got workspace event_callback_url %s; want %s", *workspace.FriendlyName, want.FriendlyName)
		}

		if *workspace.EventsFilter != want.EventsFilter {
			return fmt.Errorf("got workspace events_filter %s; want %s", *workspace.FriendlyName, want.FriendlyName)
		}

		if *workspace.MultitaskEnabled != want.MultitaskEnabled {
			return fmt.Errorf("got workspace multi_task_enabled %s; want %s", *workspace.FriendlyName, want.FriendlyName)
		}

		if *workspace.PrioritizeQueueOrder != want.PrioritizeQueueOrder {
			return fmt.Errorf("got workspace prioritize_queue_order %s; want %s", *workspace.FriendlyName, want.FriendlyName)
		}

		return nil
	}
}

func testAccCheckTwilioTaskRouterWorkspaceDestroy(s *terraform.State) error {

	client := testAccProvider.Meta().(*Config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "twilio_taskrouter_workspace" {
			continue
		}

		_, err := client.Client.TaskRouter.Workspaces.Fetch(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("taskrouter workspace still exists")
		}
	}

	return nil
}

const testAccTwilioTaskRouterWorkspaceConfig = `
resource "twilio_taskrouter_workspace" "foo" {
	friendly_name = "alex"
	event_callback_url = "https://workspace-example.free.beeceptor.com"
	events_filter = "task.created"
	multi_task_enabled = true
	prioritize_queue_order = "LIFO"
	template = "FIFO"
}
`

const testAccTwilioTaskRouterWorkspaceUpdateConfig = `
resource "twilio_taskrouter_workspace" "foo" {
	friendly_name = "will"
	event_callback_url = "https://workspace-example.free.beeceptor.com"
	events_filter = "task.created"
	multi_task_enabled = true
	prioritize_queue_order = "LIFO"
	template = "FIFO"
}
`
