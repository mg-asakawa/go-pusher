package pusher

import (
	"fmt"
	"github.com/tidwall/gjson"
)

// EventStub contains just the "type" of event.
// Knowing the type, we can then unmarshal again, into appropriate type.
type EventStub struct {
	Event string `json:"event"`
}

// Event is a pusher event
type Event struct {
	Event   string        `json:"event"`
	Data    gjson.Result  `json:"data"`
	Channel string        `json:"channel"`
}

// EventError contains a structured error in its Data field.
// It implements error.
type EventError struct {
	Event string `json:"event"`
	Data  struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"data"`
}

func (ewe EventError) Error() string {
	return fmt.Sprintf("Pusher return error: code %d, message %q", ewe.Data.Code, ewe.Data.Message)
}
