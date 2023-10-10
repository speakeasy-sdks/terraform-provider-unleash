// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ProxyFeatureSchemaVariantPayloadType - The format of the payload.
type ProxyFeatureSchemaVariantPayloadType string

const (
	ProxyFeatureSchemaVariantPayloadTypeJSON   ProxyFeatureSchemaVariantPayloadType = "json"
	ProxyFeatureSchemaVariantPayloadTypeCsv    ProxyFeatureSchemaVariantPayloadType = "csv"
	ProxyFeatureSchemaVariantPayloadTypeString ProxyFeatureSchemaVariantPayloadType = "string"
)

func (e ProxyFeatureSchemaVariantPayloadType) ToPointer() *ProxyFeatureSchemaVariantPayloadType {
	return &e
}

func (e *ProxyFeatureSchemaVariantPayloadType) UnmarshalJSON(data []byte) error {
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
		*e = ProxyFeatureSchemaVariantPayloadType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProxyFeatureSchemaVariantPayloadType: %v", v)
	}
}

// ProxyFeatureSchemaVariantPayload - Extra data configured for this variant
type ProxyFeatureSchemaVariantPayload struct {
	// The format of the payload.
	Type ProxyFeatureSchemaVariantPayloadType `json:"type"`
	// The payload value stringified.
	Value string `json:"value"`
}

// ProxyFeatureSchemaVariant - Variant details
type ProxyFeatureSchemaVariant struct {
	// Whether the variant is enabled or not.
	Enabled bool `json:"enabled"`
	// The variants name. Is unique for this feature toggle
	Name string `json:"name"`
	// Extra data configured for this variant
	Payload *ProxyFeatureSchemaVariantPayload `json:"payload,omitempty"`
}

// ProxyFeatureSchema - Frontend API feature
type ProxyFeatureSchema struct {
	// Always set to `true`.
	Enabled bool `json:"enabled"`
	// `true` if the impression data collection is enabled for the feature, otherwise `false`.
	ImpressionData bool `json:"impressionData"`
	// Unique feature name.
	Name string `json:"name"`
	// Variant details
	Variant *ProxyFeatureSchemaVariant `json:"variant,omitempty"`
}
