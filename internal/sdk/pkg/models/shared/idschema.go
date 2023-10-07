// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// IDSchema - Email id used for password reset
type IDSchema struct {
	// User email
	ID string `json:"id"`

	AdditionalProperties interface{} `json:"-"`
}
type _IDSchema IDSchema

func (c *IDSchema) UnmarshalJSON(bs []byte) error {
	data := _IDSchema{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = IDSchema(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "id")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c IDSchema) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_IDSchema(c))
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
