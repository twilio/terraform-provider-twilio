package twilio

import (
	"context"
	"fmt"
	"testing"

	openapi "github.com/twilio/twilio-go/rest/serverless/v1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	. "github.com/twilio/terraform-provider-twilio/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccServerlessSetup_basic(t *testing.T) {
	serverlessSvcResourceName := "twilio_serverless_services_v1.service"
	serverlessSvcFuncResourceName := "twilio_serverless_services_functions_v1.function"
	serverlessSvcEnvResourceName := "twilio_serverless_services_environments_v1.environment"

	serviceName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	var serviceBefore, serviceAfter openapi.ServerlessV1Environment

	resource.Test(t, resource.TestCase{
		IsUnitTest:                false,
		PreCheck:                  func() { testAccPreCheck(t) },
		ProviderFactories:         testAccProviderFactories,
		PreventPostDestroyRefresh: false,
		CheckDestroy:              testAccServerlessServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioServerlessConfig(serviceName, "Serverless func", "environment-dummy"),
				Check: resource.ComposeTestCheckFunc(
					testAccServerlessServiceEnvironmentExists(serverlessSvcEnvResourceName, &serviceBefore),
					resource.TestCheckResourceAttr(serverlessSvcResourceName, "friendly_name", "Terraform service"),
					resource.TestCheckResourceAttr(serverlessSvcResourceName, "unique_name", serviceName),
					resource.TestCheckResourceAttr(serverlessSvcFuncResourceName, "friendly_name", "Serverless func"),
					resource.TestCheckResourceAttr(serverlessSvcEnvResourceName, "unique_name", "environment-dummy"),
				),
			},
			{
				Config: testAccTwilioServerlessConfig(serviceName, "Serverless func 2", "environment-dummy-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccServerlessServiceEnvironmentExists(serverlessSvcEnvResourceName, &serviceAfter),
					resource.TestCheckResourceAttr(serverlessSvcResourceName, "friendly_name", "Terraform service"),
					resource.TestCheckResourceAttr(serverlessSvcResourceName, "unique_name", serviceName),
					resource.TestCheckResourceAttr(serverlessSvcFuncResourceName, "friendly_name", "Serverless func 2"),
					resource.TestCheckResourceAttr(serverlessSvcEnvResourceName, "unique_name", "environment-dummy-updated"),
					testAccEnvResourceWasRecreated(&serviceBefore, &serviceAfter),
				),
			},
			{
				ResourceName:      serverlessSvcResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      serverlessSvcFuncResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      serverlessSvcEnvResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
		IDRefreshName:   "",
		IDRefreshIgnore: nil,
	})
}

func testAccEnvResourceWasRecreated(before, after *openapi.ServerlessV1Environment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *before.UniqueName == *after.UniqueName {
			return fmt.Errorf("serverless Service Environment was not recreated, "+
				"before name: %s, after name: %s", *before.UniqueName, *after.UniqueName)
		}

		conn, _ := testAccProviderFactories[ProviderName]()
		conn.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))

		client := conn.Meta().(*Config)
		_, err := client.Client.ServerlessV1.FetchEnvironment(*before.ServiceSid, *before.Sid)
		if err == nil {
			return fmt.Errorf("serverless service environment with sid %s exists but should not; err: %s", *before.Sid, err)
		}

		_, err = client.Client.ServerlessV1.FetchEnvironment(*after.ServiceSid, *after.Sid)
		if err != nil {
			return fmt.Errorf("serverless service environment was not recreated becauase service %s does not exist; err: %s", *after.Sid, err)
		}

		return nil
	}
}

func testAccServerlessServiceEnvironmentExists(resourceName string, env *openapi.ServerlessV1Environment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		conn, _ := testAccProviderFactories[ProviderName]()
		conn.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))

		client := conn.Meta().(*Config)
		output, err := client.Client.ServerlessV1.FetchEnvironment(rs.Primary.Attributes["service_sid"], rs.Primary.Attributes["sid"])
		if err != nil {
			return fmt.Errorf("serverless service environment does not exist; service_sid: %s, primary id: %s",
				rs.Primary.Attributes["service_sid"], rs.Primary.ID)
		}

		*env = *output
		return nil
	}
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

func testAccTwilioServerlessConfig(serviceName string, funcName string, environmentName string) string {
	return fmt.Sprintf(`
	resource "twilio_serverless_services_v1" "service" {
		friendly_name = "Terraform service"
		unique_name   = "%s"
	}

	resource "twilio_serverless_services_functions_v1" "function" {
		service_sid   = twilio_serverless_services_v1.service.id
		friendly_name = "%s"
	}

	resource "twilio_serverless_services_environments_v1" "environment" {
  		service_sid   = twilio_serverless_services_v1.service.sid
  		unique_name   = "%s"
  		domain_suffix = "com"
	}
`, serviceName, funcName, environmentName)
}
