package core

import (
	client "github.com/twilio/twilio-go"
)

// Config is provided as context to the underlying resources.
type Config struct {
	Client *client.RestClient
}
