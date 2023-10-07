// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// OverrideSchema - An override for deciding which variant should be assigned to a user based on the context name
type OverrideSchema struct {
	// The name of the context field used to determine overrides
	ContextName string `json:"contextName"`
	// Which values that should be overriden
	Values []string `json:"values"`

	AdditionalProperties interface{} `json:"-"`
}
type _OverrideSchema OverrideSchema

func (c *OverrideSchema) UnmarshalJSON(bs []byte) error {
	data := _OverrideSchema{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = OverrideSchema(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "contextName")
	delete(additionalFields, "values")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c OverrideSchema) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_OverrideSchema(c))
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
