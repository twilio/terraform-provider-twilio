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

func testAccCheckTwilioTaskRouterWorkspaceExists(n string, hook *client.TaskRouterWorkspace) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		hookID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			// return unconvertibleIdErr(rs.Primary.ID, err)
		}
		if hookID == 0 {
			return fmt.Errorf("No repository name is set")
		}

		// org := testAccProvider.Meta().(*Organization)
		// conn := org.client
		// getHook, _, err := conn.Organizations.GetHook(context.TODO(), org.name, hookID)
		if err != nil {
			return err
		}
		// *hook = *getHook
		return nil
	}
}

type testAccTwilioTaskRouterWorkspaceExpectedAttributes struct {
	FriendlyName string
}

func testAccCheckTwilioTaskRouterWorkspaceAttributes(hook *client.TaskRouterWorkspace, want *testAccTwilioTaskRouterWorkspaceExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// if *hook.Active != want.Active {
		// 	return fmt.Errorf("got hook %t; want %t", *hook.Active, want.Active)
		// }
		// if !strings.HasPrefix(*hook.URL, "https://") {
		// 	return fmt.Errorf("got http URL %q; want to start with 'https://'", *hook.URL)
		// }
		// if !reflect.DeepEqual(hook.Events, want.Events) {
		// 	return fmt.Errorf("got hook events %q; want %q", hook.Events, want.Events)
		// }
		// if !reflect.DeepEqual(hook.Config, want.Configuration) {
		// 	return fmt.Errorf("got hook configuration %q; want %q", hook.Config, want.Configuration)
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

		// gotHook, resp, err := conn.Organizations.GetHook(context.TODO(), orgName, id)
		if err == nil {
			// if gotHook != nil && *gotHook.ID == id {
			return fmt.Errorf("Webhook still exists")
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
