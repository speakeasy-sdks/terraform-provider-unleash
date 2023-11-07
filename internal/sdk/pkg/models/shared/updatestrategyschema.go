// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UpdateStrategySchemaType - The [type of the parameter](https://docs.getunleash.io/reference/custom-activation-strategies#parameter-types)
type UpdateStrategySchemaType string

const (
	UpdateStrategySchemaTypeString     UpdateStrategySchemaType = "string"
	UpdateStrategySchemaTypePercentage UpdateStrategySchemaType = "percentage"
	UpdateStrategySchemaTypeList       UpdateStrategySchemaType = "list"
	UpdateStrategySchemaTypeNumber     UpdateStrategySchemaType = "number"
	UpdateStrategySchemaTypeBoolean    UpdateStrategySchemaType = "boolean"
)

func (e UpdateStrategySchemaType) ToPointer() *UpdateStrategySchemaType {
	return &e
}

func (e *UpdateStrategySchemaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "string":
		fallthrough
	case "percentage":
		fallthrough
	case "list":
		fallthrough
	case "number":
		fallthrough
	case "boolean":
		*e = UpdateStrategySchemaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateStrategySchemaType: %v", v)
	}
}

type UpdateStrategySchemaParameters struct {
	// A description of this strategy parameter. Use this to indicate to the users what the parameter does.
	Description *string `json:"description,omitempty"`
	// The name of the parameter
	Name string `json:"name"`
	// Whether this parameter must be configured when using the strategy. Defaults to `false`
	Required *bool `json:"required,omitempty"`
	// The [type of the parameter](https://docs.getunleash.io/reference/custom-activation-strategies#parameter-types)
	Type UpdateStrategySchemaType `json:"type"`
}

func (o *UpdateStrategySchemaParameters) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateStrategySchemaParameters) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateStrategySchemaParameters) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *UpdateStrategySchemaParameters) GetType() UpdateStrategySchemaType {
	if o == nil {
		return UpdateStrategySchemaType("")
	}
	return o.Type
}

// UpdateStrategySchema - The data required to update a strategy type.
type UpdateStrategySchema struct {
	// A description of the strategy type.
	Description *string `json:"description,omitempty"`
	// The parameter list lets you pass arguments to your custom activation strategy. These will be made available to your custom strategy implementation.
	Parameters []UpdateStrategySchemaParameters `json:"parameters"`
}

func (o *UpdateStrategySchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateStrategySchema) GetParameters() []UpdateStrategySchemaParameters {
	if o == nil {
		return []UpdateStrategySchemaParameters{}
	}
	return o.Parameters
}
