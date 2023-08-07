// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateStrategySchemaParametersInnerType - The [type of the parameter](https://docs.getunleash.io/reference/custom-activation-strategies#parameter-types)
type CreateStrategySchemaParametersInnerType string

const (
	CreateStrategySchemaParametersInnerTypeString     CreateStrategySchemaParametersInnerType = "string"
	CreateStrategySchemaParametersInnerTypePercentage CreateStrategySchemaParametersInnerType = "percentage"
	CreateStrategySchemaParametersInnerTypeList       CreateStrategySchemaParametersInnerType = "list"
	CreateStrategySchemaParametersInnerTypeNumber     CreateStrategySchemaParametersInnerType = "number"
	CreateStrategySchemaParametersInnerTypeBoolean    CreateStrategySchemaParametersInnerType = "boolean"
)

func (e CreateStrategySchemaParametersInnerType) ToPointer() *CreateStrategySchemaParametersInnerType {
	return &e
}

func (e *CreateStrategySchemaParametersInnerType) UnmarshalJSON(data []byte) error {
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
		*e = CreateStrategySchemaParametersInnerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStrategySchemaParametersInnerType: %v", v)
	}
}

type CreateStrategySchemaParametersInner struct {
	// A description of this strategy parameter. Use this to indicate to the users what the parameter does.
	Description *string `json:"description,omitempty"`
	// The name of the parameter
	Name string `json:"name"`
	// Whether this parameter must be configured when using the strategy. Defaults to `false`
	Required *bool `json:"required,omitempty"`
	// The [type of the parameter](https://docs.getunleash.io/reference/custom-activation-strategies#parameter-types)
	Type CreateStrategySchemaParametersInnerType `json:"type"`
}