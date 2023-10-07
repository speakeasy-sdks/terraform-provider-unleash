// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// FeatureEventsSchemaVersion - An API versioning number
type FeatureEventsSchemaVersion int64

const (
	FeatureEventsSchemaVersionOne FeatureEventsSchemaVersion = 1
)

func (e FeatureEventsSchemaVersion) ToPointer() *FeatureEventsSchemaVersion {
	return &e
}

func (e *FeatureEventsSchemaVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		*e = FeatureEventsSchemaVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FeatureEventsSchemaVersion: %v", v)
	}
}

// FeatureEventsSchema - One or more events happening to a specific feature toggle
type FeatureEventsSchema struct {
	// The list of events
	Events []EventSchema `json:"events"`
	// The name of the feature toggle these events relate to
	ToggleName *string `json:"toggleName,omitempty"`
	// How many events are there for this feature toggle
	TotalEvents *int64 `json:"totalEvents,omitempty"`
	// An API versioning number
	Version *FeatureEventsSchemaVersion `json:"version,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _FeatureEventsSchema FeatureEventsSchema

func (c *FeatureEventsSchema) UnmarshalJSON(bs []byte) error {
	data := _FeatureEventsSchema{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = FeatureEventsSchema(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "events")
	delete(additionalFields, "toggleName")
	delete(additionalFields, "totalEvents")
	delete(additionalFields, "version")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c FeatureEventsSchema) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_FeatureEventsSchema(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}
