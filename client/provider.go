package client

import (
	client "github.com/twilio/twilio-go/twilio"
)

// Config is provided as context to the underlying resources.
type Config struct {
	Client *client.Twilio
}
