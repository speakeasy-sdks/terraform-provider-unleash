// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateStrategySchemaType - The [type of the parameter](https://docs.getunleash.io/reference/custom-activation-strategies#parameter-types)
type CreateStrategySchemaType string

const (
	CreateStrategySchemaTypeString     CreateStrategySchemaType = "string"
	CreateStrategySchemaTypePercentage CreateStrategySchemaType = "percentage"
	CreateStrategySchemaTypeList       CreateStrategySchemaType = "list"
	CreateStrategySchemaTypeNumber     CreateStrategySchemaType = "number"
	CreateStrategySchemaTypeBoolean    CreateStrategySchemaType = "boolean"
)

func (e CreateStrategySchemaType) ToPointer() *CreateStrategySchemaType {
	return &e
}

func (e *CreateStrategySchemaType) UnmarshalJSON(data []byte) error {
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
		*e = CreateStrategySchemaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStrategySchemaType: %v", v)
	}
}

type Parameters struct {
	// A description of this strategy parameter. Use this to indicate to the users what the parameter does.
	Description *string `json:"description,omitempty"`
	// The name of the parameter
	Name string `json:"name"`
	// Whether this parameter must be configured when using the strategy. Defaults to `false`
	Required *bool `json:"required,omitempty"`
	// The [type of the parameter](https://docs.getunleash.io/reference/custom-activation-strategies#parameter-types)
	Type CreateStrategySchemaType `json:"type"`
}

func (o *Parameters) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Parameters) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Parameters) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *Parameters) GetType() CreateStrategySchemaType {
	if o == nil {
		return CreateStrategySchemaType("")
	}
	return o.Type
}

// CreateStrategySchema - The data required to create a strategy type. Refer to the docs on [custom strategy types](https://docs.getunleash.io/reference/custom-activation-strategies) for more information.
type CreateStrategySchema struct {
	// A description of the strategy type.
	Description *string `json:"description,omitempty"`
	// The name of the strategy type. Must be unique.
	Name string `json:"name"`
	// The parameter list lets you pass arguments to your custom activation strategy. These will be made available to your custom strategy implementation.
	Parameters []Parameters `json:"parameters"`
}

func (o *CreateStrategySchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateStrategySchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateStrategySchema) GetParameters() []Parameters {
	if o == nil {
		return []Parameters{}
	}
	return o.Parameters
}
