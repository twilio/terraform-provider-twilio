package twilio

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	. "github.com/twilio/terraform-provider-twilio/client"
)

func TestAccFlexSetup_basic(t *testing.T) {
	flexResourceName := "twilio_flex_flex_flows_v1.flows"
	chatResourceName := "twilio_chat_services_v2.chat_service"
	studioResourceName := "twilio_studio_flows_v2.studio_flow"

	chatServiceName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	studioFlowName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	flexName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		IsUnitTest:                false,
		PreCheck:                  func() { testAccPreCheck(t) },
		ProviderFactories:         testAccProviderFactories,
		PreventPostDestroyRefresh: false,
		CheckDestroy:              testAccFlexInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioFlexConfig(chatServiceName, studioFlowName, flexName, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(chatResourceName, "friendly_name", chatServiceName),
					resource.TestCheckResourceAttr(chatResourceName, "reachability_enabled", "false"),
					resource.TestCheckResourceAttr(studioResourceName, "friendly_name", studioFlowName),
					resource.TestCheckResourceAttr(studioResourceName, "commit_message", "first draft"),
					resource.TestCheckResourceAttr(studioResourceName, "status", "draft"),
					resource.TestCheckResourceAttr(flexResourceName, "friendly_name", flexName),
					resource.TestCheckResourceAttr(flexResourceName, "channel_type", "sms"),
					resource.TestCheckResourceAttr(flexResourceName, "contact_identity", "true"),
					resource.TestCheckResourceAttr(flexResourceName, "enabled", "true"),
				),
			},
			{
				Config: testAccTwilioFlexConfig(chatServiceName, studioFlowName, flexName, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(chatResourceName, "friendly_name", chatServiceName),
					resource.TestCheckResourceAttr(studioResourceName, "friendly_name", studioFlowName),
					resource.TestCheckResourceAttr(flexResourceName, "friendly_name", flexName),
					resource.TestCheckResourceAttr(flexResourceName, "channel_type", "sms"),
					resource.TestCheckResourceAttr(flexResourceName, "contact_identity", "false"),
					resource.TestCheckResourceAttr(flexResourceName, "enabled", "false"),
				),
			},
		},
		IDRefreshName:   "",
		IDRefreshIgnore: nil,
	})
}

func testAccFlexInstanceDestroy(state *terraform.State) error {

	for _, rs := range state.RootModule().Resources {

		conn, _ := testAccProviderFactories[ProviderName]()
		conn.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
		client := conn.Meta().(*Config)

		if rs.Type == "twilio_chat_services_v2" {
			//Check if the chat service has destroyed
			_, err := client.Client.ChatV2.FetchService(rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("flex chat service still exists")
			}
		} else if rs.Type == "twilio_studio_flows_v2" {
			//Check if the studio flow has destroyed
			_, err := client.Client.StudioV2.FetchFlow(rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("flex studio flow still exists")
			}
		} else if rs.Type == "twilio_flex_flex_flows_v1" {
			//Check if the flex instance has destroyed
			_, err := client.Client.FlexV1.FetchFlexFlow(rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("flex flow still exists")
			}
		} else {
			continue
		}
	}

	return nil
}

func testAccTwilioFlexConfig(chatResourceName string, studioResourceName string, rName string, enabled string) string {
	return fmt.Sprintf(`

	resource "twilio_chat_services_v2" "chat_service" {
		friendly_name = "%s"
		webhook_filters = ["onMessageSent",
		"onChannelDestroyed",
		"onChannelUpdated"]
	}

	resource "twilio_studio_flows_v2" "studio_flow" {
		commit_message = "first draft"
		friendly_name  = "%s"
		status         = "draft"
		definition = jsonencode({
		description = "A New Flow",
		states = [{
		  name        = "Trigger"
		  type        = "trigger"
		  transitions = []
		  properties = {
			offset = {
			  x = 0
			  y = 0
			}
		  }
		}]
		initial_state = "Trigger"
		flags = {
		  allow_concurrent_calls = true
		}
		})
	}

	resource "twilio_flex_flex_flows_v1" "flows" {
	  friendly_name        = "%s"
	  chat_service_sid     = twilio_chat_services_v2.chat_service.id
	  channel_type         = "sms"
	  integration_flow_sid = twilio_studio_flows_v2.studio_flow.id
	  contact_identity     = "true"
	  enabled              = %s
}
`, chatResourceName, studioResourceName, rName, enabled)
}
