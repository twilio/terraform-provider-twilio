package twilio

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	. "github.com/twilio/terraform-provider-twilio/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccServerlessSetup_basic(t *testing.T) {
	serverlessSvcResourceName := "twilio_serverless_services_v1.service"
	serverlessSvcFuncResourceName := "twilio_serverless_services_functions_v1.function"

	serviceName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		IsUnitTest:                false,
		PreCheck:                  func() { testAccPreCheck(t) },
		ProviderFactories:         testAccProviderFactories,
		PreventPostDestroyRefresh: false,
		CheckDestroy:              testAccServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessConfig(serviceName, "Serverless func"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(serverlessSvcResourceName, "friendly_name", "Terraform service"),
					resource.TestCheckResourceAttr(serverlessSvcResourceName, "unique_name", serviceName),
					resource.TestCheckResourceAttr(serverlessSvcFuncResourceName, "friendly_name", "Serverless func"),
				),
			},
			{
				Config: testAccTwilioServerlessConfig(serviceName, "Serverless func 2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(serverlessSvcResourceName, "friendly_name", "Terraform service"),
					resource.TestCheckResourceAttr(serverlessSvcResourceName, "unique_name", serviceName),
					resource.TestCheckResourceAttr(serverlessSvcFuncResourceName, "friendly_name", "Serverless func 2"),
				),
			},
		},
		IDRefreshName:   "",
		IDRefreshIgnore: nil,
	})
}

func testAccServerlessServiceDestroy(state *terraform.State) error {
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "twilio_serverless_services_v1" {
			continue
		}

		conn, _ := testAccProviderFactories[ProviderName]()
		conn.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))

		client := conn.Meta().(*Config)
		_, err := client.Client.ServerlessV1.FetchService(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("serverless service still exists")
		}
	}

	return nil
}

func testAccTwilioServerlessConfig(serviceName string, funcName string) string {
	return fmt.Sprintf(`
	resource "twilio_serverless_services_v1" "service" {
		friendly_name = "Terraform service"
		unique_name   = "%s"
	}

	resource "twilio_serverless_services_functions_v1" "function" {
		service_sid   = twilio_serverless_services_v1.service.id 
		friendly_name = "%s"
	}
`, serviceName, funcName)
}
