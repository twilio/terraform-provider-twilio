package twilio

import (
	"fmt"
	"strconv"
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
						FriendlyName: "alex",
					}),
				),
			},
			{
				Config: testAccTwilioTaskRouterWorkspaceUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioTaskRouterWorkspaceExists(rn, &workspace),
					testAccCheckTwilioTaskRouterWorkspaceAttributes(&workspace, &testAccTwilioTaskRouterWorkspaceExpectedAttributes{
						FriendlyName: "will",
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
	FriendlyName string
}

func testAccCheckTwilioTaskRouterWorkspaceAttributes(workspace *client.TaskRouterWorkspace, want *testAccTwilioTaskRouterWorkspaceExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if *workspace.FriendlyName != want.FriendlyName {
			return fmt.Errorf("got workspace friendly name %s; want %s", *workspace.FriendlyName, want.FriendlyName)
		}
		// if !strings.HasPrefix(*workspace.URL, "https://") {
		// 	return fmt.Errorf("got http URL %q; want to start with 'https://'", *workspace.URL)
		// }
		// if !reflect.DeepEqual(workspace.Events, want.Events) {
		// 	return fmt.Errorf("got workspace events %q; want %q", workspace.Events, want.Events)
		// }
		// if !reflect.DeepEqual(workspace.Config, want.Configuration) {
		// 	return fmt.Errorf("got workspace configuration %q; want %q", workspace.Config, want.Configuration)
		// }

		return nil
	}
}

func testAccCheckTwilioTaskRouterWorkspaceDestroy(s *terraform.State) error {
	// conn := testAccProvider.Meta().(*Organization).client
	// orgName := testAccProvider.Meta().(*Organization).name

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "twilio_taskrouter_workspace" {
			continue
		}

		_, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			// return unconvertibleIdErr(rs.Primary.ID, err)
		}

		// gotworkspace, resp, err := conn.Organizations.Getworkspace(context.TODO(), orgName, id)
		if err == nil {
			// if gotworkspace != nil && *gotworkspace.ID == id {
			return fmt.Errorf("Webworkspace still exists")
			// }
		}
		// if resp.StatusCode != 404 {
		// 	return err
		// }
		return nil
	}
	return nil
}

const testAccTwilioTaskRouterWorkspaceConfig = `
resource "twilio_taskrouter_workspace" "foo" {
  friendly_name = "alex"
}
`

const testAccTwilioTaskRouterWorkspaceUpdateConfig = `
resource "twilio_taskrouter_workspace" "foo" {
  friendly_name = "will"
}
`
