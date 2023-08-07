// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// StrategyVariantSchemaPayloadType - The type of the value. Commonly used types are string, json and csv.
type StrategyVariantSchemaPayloadType string

const (
	StrategyVariantSchemaPayloadTypeJSON   StrategyVariantSchemaPayloadType = "json"
	StrategyVariantSchemaPayloadTypeCsv    StrategyVariantSchemaPayloadType = "csv"
	StrategyVariantSchemaPayloadTypeString StrategyVariantSchemaPayloadType = "string"
)

func (e StrategyVariantSchemaPayloadType) ToPointer() *StrategyVariantSchemaPayloadType {
	return &e
}

func (e *StrategyVariantSchemaPayloadType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "csv":
		fallthrough
	case "string":
		*e = StrategyVariantSchemaPayloadType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StrategyVariantSchemaPayloadType: %v", v)
	}
}

// StrategyVariantSchemaPayload - Extra data configured for this variant
type StrategyVariantSchemaPayload struct {
	// The type of the value. Commonly used types are string, json and csv.
	Type StrategyVariantSchemaPayloadType `json:"type"`
	// The actual value of payload
	Value string `json:"value"`
}