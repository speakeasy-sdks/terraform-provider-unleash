// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// SplashResponseSchema - Data related to a user having seen a splash screen.
type SplashResponseSchema struct {
	// Indicates whether the user has seen the splash screen or not.
	Seen bool `json:"seen"`
	// The ID of the splash screen that was shown.
	SplashID string `json:"splashId"`
	// The ID of the user that was shown the splash screen.
	UserID int64 `json:"userId"`

	AdditionalProperties interface{} `json:"-"`
}
type _SplashResponseSchema SplashResponseSchema

func (c *SplashResponseSchema) UnmarshalJSON(bs []byte) error {
	data := _SplashResponseSchema{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SplashResponseSchema(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "seen")
	delete(additionalFields, "splashId")
	delete(additionalFields, "userId")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SplashResponseSchema) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SplashResponseSchema(c))
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
