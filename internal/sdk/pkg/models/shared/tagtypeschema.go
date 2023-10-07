// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// TagTypeSchema - A tag type.
type TagTypeSchema struct {
	// The description of the tag type.
	Description *string `json:"description,omitempty"`
	// The icon of the tag type.
	Icon *string `json:"icon,omitempty"`
	// The name of the tag type.
	Name string `json:"name"`

	AdditionalProperties interface{} `json:"-"`
}
type _TagTypeSchema TagTypeSchema

func (c *TagTypeSchema) UnmarshalJSON(bs []byte) error {
	data := _TagTypeSchema{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = TagTypeSchema(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "description")
	delete(additionalFields, "icon")
	delete(additionalFields, "name")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c TagTypeSchema) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_TagTypeSchema(c))
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
