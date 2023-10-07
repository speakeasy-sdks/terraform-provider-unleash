// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// TagsBulkAddSchema - Represents tag changes to be applied to a list of features.
type TagsBulkAddSchema struct {
	// The list of features that will be affected by the tag changes.
	Features []string `json:"features"`
	// Represents a set of changes to a feature's tags, such as adding or removing tags.
	Tags UpdateTagsSchema `json:"tags"`

	AdditionalProperties interface{} `json:"-"`
}
type _TagsBulkAddSchema TagsBulkAddSchema

func (c *TagsBulkAddSchema) UnmarshalJSON(bs []byte) error {
	data := _TagsBulkAddSchema{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = TagsBulkAddSchema(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "features")
	delete(additionalFields, "tags")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c TagsBulkAddSchema) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_TagsBulkAddSchema(c))
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
