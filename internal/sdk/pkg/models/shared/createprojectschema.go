// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Mode - A mode of the project affecting what actions are possible in this project
type Mode string

const (
	ModeOpen      Mode = "open"
	ModeProtected Mode = "protected"
)

func (e Mode) ToPointer() *Mode {
	return &e
}

func (e *Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "open":
		fallthrough
	case "protected":
		*e = Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mode: %v", v)
	}
}

type CreateProjectSchema struct {
	// A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy
	DefaultStickiness *string `json:"defaultStickiness,omitempty"`
	// The project's description.
	Description *string `json:"description,omitempty"`
	// A limit on the number of features allowed in the project. `null` if no limit.
	FeatureLimit *int64 `json:"featureLimit,omitempty"`
	// The project's identifier.
	ID string `json:"id"`
	// A mode of the project affecting what actions are possible in this project
	Mode *Mode `json:"mode,omitempty"`
	// The project's name.
	Name string `json:"name"`
}

func (o *CreateProjectSchema) GetDefaultStickiness() *string {
	if o == nil {
		return nil
	}
	return o.DefaultStickiness
}

func (o *CreateProjectSchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateProjectSchema) GetFeatureLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.FeatureLimit
}

func (o *CreateProjectSchema) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateProjectSchema) GetMode() *Mode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *CreateProjectSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
