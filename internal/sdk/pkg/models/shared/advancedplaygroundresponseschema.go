// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// AdvancedPlaygroundResponseSchema - The state of all features given the provided input.
type AdvancedPlaygroundResponseSchema struct {
	// The list of features that have been evaluated.
	Features []AdvancedPlaygroundFeatureSchema `json:"features"`
	// Data for the playground API to evaluate toggles in advanced mode with environment and context multi selection
	Input AdvancedPlaygroundRequestSchema `json:"input"`

	AdditionalProperties interface{} `json:"-"`
}
type _AdvancedPlaygroundResponseSchema AdvancedPlaygroundResponseSchema

func (c *AdvancedPlaygroundResponseSchema) UnmarshalJSON(bs []byte) error {
	data := _AdvancedPlaygroundResponseSchema{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = AdvancedPlaygroundResponseSchema(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "features")
	delete(additionalFields, "input")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c AdvancedPlaygroundResponseSchema) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_AdvancedPlaygroundResponseSchema(c))
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
