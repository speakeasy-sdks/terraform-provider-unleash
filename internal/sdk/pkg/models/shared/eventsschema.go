// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// EventsSchemaVersion - The api version of this response. A natural increasing number. Only increases if format changes
type EventsSchemaVersion int64

const (
	EventsSchemaVersionOne EventsSchemaVersion = 1
)

func (e EventsSchemaVersion) ToPointer() *EventsSchemaVersion {
	return &e
}

func (e *EventsSchemaVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		*e = EventsSchemaVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EventsSchemaVersion: %v", v)
	}
}

// EventsSchema - A list of events that has happened in the system
type EventsSchema struct {
	// The list of events
	Events []EventSchema `json:"events"`
	// The total count of events
	TotalEvents *int64 `json:"totalEvents,omitempty"`
	// The api version of this response. A natural increasing number. Only increases if format changes
	Version EventsSchemaVersion `json:"version"`
}
