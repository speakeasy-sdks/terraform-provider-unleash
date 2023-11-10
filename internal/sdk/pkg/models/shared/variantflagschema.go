// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// VariantFlagSchemaType - The type of data contained.
type VariantFlagSchemaType string

const (
	VariantFlagSchemaTypeString VariantFlagSchemaType = "string"
	VariantFlagSchemaTypeJSON   VariantFlagSchemaType = "json"
	VariantFlagSchemaTypeCsv    VariantFlagSchemaType = "csv"
)

func (e VariantFlagSchemaType) ToPointer() *VariantFlagSchemaType {
	return &e
}

func (e *VariantFlagSchemaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "string":
		fallthrough
	case "json":
		fallthrough
	case "csv":
		*e = VariantFlagSchemaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VariantFlagSchemaType: %v", v)
	}
}

// VariantFlagSchemaPayload - Additional data associated with this variant.
type VariantFlagSchemaPayload struct {
	// The type of data contained.
	Type *VariantFlagSchemaType `json:"type,omitempty"`
	// The actual associated data
	Value *string `json:"value,omitempty"`
}

func (o *VariantFlagSchemaPayload) GetType() *VariantFlagSchemaType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *VariantFlagSchemaPayload) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// VariantFlagSchema - A representation of an evaluated Unleash feature variant.
type VariantFlagSchema struct {
	// Whether the variant is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the variant. Will always be disabled if `enabled` is false.
	Name *string `json:"name,omitempty"`
	// Additional data associated with this variant.
	Payload *VariantFlagSchemaPayload `json:"payload,omitempty"`
}

func (o *VariantFlagSchema) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *VariantFlagSchema) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *VariantFlagSchema) GetPayload() *VariantFlagSchemaPayload {
	if o == nil {
		return nil
	}
	return o.Payload
}
