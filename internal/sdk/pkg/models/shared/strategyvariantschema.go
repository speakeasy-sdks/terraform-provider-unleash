// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// StrategyVariantSchemaWeightType - Set to `fix` if this variant must have exactly the weight allocated to it. If the type is `variable`, the weight will adjust so that the total weight of all variants adds up to 1000. Refer to the [variant weight documentation](https://docs.getunleash.io/reference/feature-toggle-variants#variant-weight).
type StrategyVariantSchemaWeightType string

const (
	StrategyVariantSchemaWeightTypeVariable StrategyVariantSchemaWeightType = "variable"
	StrategyVariantSchemaWeightTypeFix      StrategyVariantSchemaWeightType = "fix"
)

func (e StrategyVariantSchemaWeightType) ToPointer() *StrategyVariantSchemaWeightType {
	return &e
}

func (e *StrategyVariantSchemaWeightType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "variable":
		fallthrough
	case "fix":
		*e = StrategyVariantSchemaWeightType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StrategyVariantSchemaWeightType: %v", v)
	}
}

// StrategyVariantSchema - This is an experimental property. It may change or be removed as we work on it. Please don't depend on it yet. A strategy variant allows you to attach any data to strategies instead of only returning `true`/`false`. Strategy variants take precedence over feature variants.
type StrategyVariantSchema struct {
	// The variant name. Must be unique for this feature toggle
	Name string `json:"name"`
	// Extra data configured for this variant
	Payload *StrategyVariantSchemaPayload `json:"payload,omitempty"`
	// The [stickiness](https://docs.getunleash.io/reference/feature-toggle-variants#variant-stickiness) to use for distribution of this variant. Stickiness is how Unleash guarantees that the same user gets the same variant every time
	Stickiness string `json:"stickiness"`
	// The weight is the likelihood of any one user getting this variant. It is an integer between 0 and 1000. See the section on [variant weights](https://docs.getunleash.io/reference/feature-toggle-variants#variant-weight) for more information
	Weight int64 `json:"weight"`
	// Set to `fix` if this variant must have exactly the weight allocated to it. If the type is `variable`, the weight will adjust so that the total weight of all variants adds up to 1000. Refer to the [variant weight documentation](https://docs.getunleash.io/reference/feature-toggle-variants#variant-weight).
	WeightType StrategyVariantSchemaWeightType `json:"weightType"`
}