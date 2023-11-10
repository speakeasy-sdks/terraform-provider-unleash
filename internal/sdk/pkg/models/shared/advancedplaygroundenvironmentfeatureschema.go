// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
)

// AdvancedPlaygroundEnvironmentFeatureSchema2 - The cumulative results of all the feature's strategies. Can be `true`,
//
//	`false`, or `unknown`.
//	This property will only be `unknown`
//	if one or more of the strategies can't be fully evaluated and the rest of the strategies
//	all resolve to `false`.
type AdvancedPlaygroundEnvironmentFeatureSchema2 string

const (
	AdvancedPlaygroundEnvironmentFeatureSchema2Unknown AdvancedPlaygroundEnvironmentFeatureSchema2 = "unknown"
)

func (e AdvancedPlaygroundEnvironmentFeatureSchema2) ToPointer() *AdvancedPlaygroundEnvironmentFeatureSchema2 {
	return &e
}

func (e *AdvancedPlaygroundEnvironmentFeatureSchema2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unknown":
		*e = AdvancedPlaygroundEnvironmentFeatureSchema2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AdvancedPlaygroundEnvironmentFeatureSchema2: %v", v)
	}
}

type AdvancedPlaygroundEnvironmentFeatureSchemaResultType string

const (
	AdvancedPlaygroundEnvironmentFeatureSchemaResultTypeBoolean                                     AdvancedPlaygroundEnvironmentFeatureSchemaResultType = "boolean"
	AdvancedPlaygroundEnvironmentFeatureSchemaResultTypeAdvancedPlaygroundEnvironmentFeatureSchema2 AdvancedPlaygroundEnvironmentFeatureSchemaResultType = "advancedPlaygroundEnvironmentFeatureSchema_2"
)

type AdvancedPlaygroundEnvironmentFeatureSchemaResult struct {
	Boolean                                     *bool
	AdvancedPlaygroundEnvironmentFeatureSchema2 *AdvancedPlaygroundEnvironmentFeatureSchema2

	Type AdvancedPlaygroundEnvironmentFeatureSchemaResultType
}

func CreateAdvancedPlaygroundEnvironmentFeatureSchemaResultBoolean(boolean bool) AdvancedPlaygroundEnvironmentFeatureSchemaResult {
	typ := AdvancedPlaygroundEnvironmentFeatureSchemaResultTypeBoolean

	return AdvancedPlaygroundEnvironmentFeatureSchemaResult{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateAdvancedPlaygroundEnvironmentFeatureSchemaResultAdvancedPlaygroundEnvironmentFeatureSchema2(advancedPlaygroundEnvironmentFeatureSchema2 AdvancedPlaygroundEnvironmentFeatureSchema2) AdvancedPlaygroundEnvironmentFeatureSchemaResult {
	typ := AdvancedPlaygroundEnvironmentFeatureSchemaResultTypeAdvancedPlaygroundEnvironmentFeatureSchema2

	return AdvancedPlaygroundEnvironmentFeatureSchemaResult{
		AdvancedPlaygroundEnvironmentFeatureSchema2: &advancedPlaygroundEnvironmentFeatureSchema2,
		Type: typ,
	}
}

func (u *AdvancedPlaygroundEnvironmentFeatureSchemaResult) UnmarshalJSON(data []byte) error {

	boolean := new(bool)
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = boolean
		u.Type = AdvancedPlaygroundEnvironmentFeatureSchemaResultTypeBoolean
		return nil
	}

	advancedPlaygroundEnvironmentFeatureSchema2 := new(AdvancedPlaygroundEnvironmentFeatureSchema2)
	if err := utils.UnmarshalJSON(data, &advancedPlaygroundEnvironmentFeatureSchema2, "", true, true); err == nil {
		u.AdvancedPlaygroundEnvironmentFeatureSchema2 = advancedPlaygroundEnvironmentFeatureSchema2
		u.Type = AdvancedPlaygroundEnvironmentFeatureSchemaResultTypeAdvancedPlaygroundEnvironmentFeatureSchema2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AdvancedPlaygroundEnvironmentFeatureSchemaResult) MarshalJSON() ([]byte, error) {
	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.AdvancedPlaygroundEnvironmentFeatureSchema2 != nil {
		return utils.MarshalJSON(u.AdvancedPlaygroundEnvironmentFeatureSchema2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// AdvancedPlaygroundEnvironmentFeatureSchemaStrategies - Feature's applicable strategies and cumulative results of the strategies
type AdvancedPlaygroundEnvironmentFeatureSchemaStrategies struct {
	// The strategies that apply to this feature.
	Data []PlaygroundStrategySchema `json:"data"`
	// The cumulative results of all the feature's strategies. Can be `true`,
	//                                   `false`, or `unknown`.
	//                                   This property will only be `unknown`
	//                                   if one or more of the strategies can't be fully evaluated and the rest of the strategies
	//                                   all resolve to `false`.
	Result AdvancedPlaygroundEnvironmentFeatureSchemaResult `json:"result"`
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) GetData() []PlaygroundStrategySchema {
	if o == nil {
		return []PlaygroundStrategySchema{}
	}
	return o.Data
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) GetResult() AdvancedPlaygroundEnvironmentFeatureSchemaResult {
	if o == nil {
		return AdvancedPlaygroundEnvironmentFeatureSchemaResult{}
	}
	return o.Result
}

// AdvancedPlaygroundEnvironmentFeatureSchemaPayload - An optional payload attached to the variant.
type AdvancedPlaygroundEnvironmentFeatureSchemaPayload struct {
	// The format of the payload.
	Type string `json:"type"`
	// The payload value stringified.
	Value string `json:"value"`
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchemaPayload) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchemaPayload) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// Variant - The feature variant you receive based on the provided context or the _disabled
//
//	variant_. If a feature is disabled or doesn't have any
//	variants, you would get the _disabled variant_.
//	Otherwise, you'll get one of the feature's defined variants.
type Variant struct {
	// Whether the variant is enabled or not. If the feature is disabled or if it doesn't have variants, this property will be `false`
	Enabled bool `json:"enabled"`
	// The variant's name. If there is no variant or if the toggle is disabled, this will be `disabled`
	Name string `json:"name"`
	// An optional payload attached to the variant.
	Payload *AdvancedPlaygroundEnvironmentFeatureSchemaPayload `json:"payload,omitempty"`
}

func (o *Variant) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *Variant) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Variant) GetPayload() *AdvancedPlaygroundEnvironmentFeatureSchemaPayload {
	if o == nil {
		return nil
	}
	return o.Payload
}

// AdvancedPlaygroundEnvironmentFeatureSchema - A simplified feature toggle model intended for the Unleash playground.
type AdvancedPlaygroundEnvironmentFeatureSchema struct {
	// The Unleash context as modeled in client SDKs
	Context SDKContextSchema `json:"context"`
	// The feature's environment.
	Environment string `json:"environment"`
	// Whether this feature is enabled or not in the current environment.
	//                           If a feature can't be fully evaluated (that is, `strategies.result` is `unknown`),
	//                           this will be `false` to align with how client SDKs treat unresolved feature states.
	IsEnabled bool `json:"isEnabled"`
	// Whether the feature is active and would be evaluated in the provided environment in a normal SDK context.
	IsEnabledInCurrentEnvironment bool `json:"isEnabledInCurrentEnvironment"`
	// The feature's name.
	Name string `json:"name"`
	// The ID of the project that contains this feature.
	ProjectID string `json:"projectId"`
	// Feature's applicable strategies and cumulative results of the strategies
	Strategies AdvancedPlaygroundEnvironmentFeatureSchemaStrategies `json:"strategies"`
	// The feature variant you receive based on the provided context or the _disabled
	//                           variant_. If a feature is disabled or doesn't have any
	//                           variants, you would get the _disabled variant_.
	//                           Otherwise, you'll get one of the feature's defined variants.
	Variant *Variant `json:"variant"`
	// The feature variants.
	Variants []VariantSchema `json:"variants"`
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetContext() SDKContextSchema {
	if o == nil {
		return SDKContextSchema{}
	}
	return o.Context
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetIsEnabled() bool {
	if o == nil {
		return false
	}
	return o.IsEnabled
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetIsEnabledInCurrentEnvironment() bool {
	if o == nil {
		return false
	}
	return o.IsEnabledInCurrentEnvironment
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetStrategies() AdvancedPlaygroundEnvironmentFeatureSchemaStrategies {
	if o == nil {
		return AdvancedPlaygroundEnvironmentFeatureSchemaStrategies{}
	}
	return o.Strategies
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetVariant() *Variant {
	if o == nil {
		return nil
	}
	return o.Variant
}

func (o *AdvancedPlaygroundEnvironmentFeatureSchema) GetVariants() []VariantSchema {
	if o == nil {
		return []VariantSchema{}
	}
	return o.Variants
}
