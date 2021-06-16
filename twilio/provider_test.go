package twilio

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twilio/terraform-provider-twilio/twilio/resources"
)

func TestTwilioResourcesMap(t *testing.T) {
	twilioResources := resources.NewTwilioResources()
	assert.NotNil(t, twilioResources.Map["twilio_api_accounts_messages_v2010"])
}
