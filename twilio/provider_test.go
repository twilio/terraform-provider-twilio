package twilio

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ProviderName = "twilio"

var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		ProviderName: func() (*schema.Provider, error) { return Provider(), nil },
	}
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TWILIO_ACCOUNT_SID"); v == "" {
		t.Fatal("ACCOUNT_SID must be set for acceptance tests")
	}
	if v := os.Getenv("TWILIO_AUTH_TOKEN"); v == "" {
		t.Fatal("AUTH_TOKEN must be set for acceptance tests")
	}
}
