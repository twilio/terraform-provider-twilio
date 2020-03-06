package twilio

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testNamePrefix = "tf-acc-test-"

var (
	testAccProviders map[string]terraform.ResourceProvider
	testAccProvider  *schema.Provider
)

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"twilio": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	fmt.Printf("\n STARTING: %s, %s \n", os.Getenv("ACCOUNT_SID"), os.Getenv("AUTH_TOKEN"))
	if v := os.Getenv("ACCOUNT_SID"); v == "" {
		t.Fatal("DIGITALOCEAN_TOKEN must be set for acceptance tests")
	}

	if v := os.Getenv("AUTH_TOKEN"); v == "" {
		t.Fatal("DIGITALOCEAN_TOKEN must be set for acceptance tests")
	}
}

func TestURLOverride(t *testing.T) {
	customEndpoint := "https://mock-api.internal.example.com/"

	rawProvider := Provider()
	raw := map[string]interface{}{
		"token":        "12345",
		"api_endpoint": customEndpoint,
	}

	err := rawProvider.(*schema.Provider).Configure(terraform.NewResourceConfigRaw(raw))
	meta := rawProvider.(*schema.Provider).Meta()
	if meta == nil {
		t.Fatalf("Expected metadata, got nil: err: %s", err)
	}
	// client := meta.(*Config)
	// if client.BaseURL.String() != customEndpoint {
	// 	t.Fatalf("Expected %s, got %s", customEndpoint, client.BaseURL.String())
	// }
}

func TestURLDefault(t *testing.T) {
	rawProvider := Provider()
	raw := map[string]interface{}{
		"token": "12345",
	}

	err := rawProvider.(*schema.Provider).Configure(terraform.NewResourceConfigRaw(raw))
	meta := rawProvider.(*schema.Provider).Meta()
	if meta == nil {
		t.Fatalf("Expected metadata, got nil: err: %s", err)
	}
	// client := meta.(*Config)
	// if client.BaseURL.String() != "https://api.digitalocean.com" {
	// 	t.Fatalf("Expected %s, got %s", "https://api.digitalocean.com", client.BaseURL.String())
	// }
}

func randomTestName() string {
	return randomName(testNamePrefix, 10)
}

func randomName(prefix string, length int) string {
	return fmt.Sprintf("%s%s", prefix, acctest.RandString(length))
}
