// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// PlaygroundFeatureSchemaStrategiesResult2 - The cumulative results of all the feature's strategies. Can be `true`,
//
//	`false`, or `unknown`.
//	This property will only be `unknown`
//	if one or more of the strategies can't be fully evaluated and the rest of the strategies
//	all resolve to `false`.
type PlaygroundFeatureSchemaStrategiesResult2 string

const (
	PlaygroundFeatureSchemaStrategiesResult2Unknown PlaygroundFeatureSchemaStrategiesResult2 = "unknown"
)

func (e PlaygroundFeatureSchemaStrategiesResult2) ToPointer() *PlaygroundFeatureSchemaStrategiesResult2 {
	return &e
}

func (e *PlaygroundFeatureSchemaStrategiesResult2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unknown":
		*e = PlaygroundFeatureSchemaStrategiesResult2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaygroundFeatureSchemaStrategiesResult2: %v", v)
	}
}

type PlaygroundFeatureSchemaStrategiesResultType string

const (
	PlaygroundFeatureSchemaStrategiesResultTypeBoolean                                  PlaygroundFeatureSchemaStrategiesResultType = "boolean"
	PlaygroundFeatureSchemaStrategiesResultTypePlaygroundFeatureSchemaStrategiesResult2 PlaygroundFeatureSchemaStrategiesResultType = "playgroundFeatureSchema_strategies_result_2"
)

type PlaygroundFeatureSchemaStrategiesResult struct {
	Boolean                                  *bool
	PlaygroundFeatureSchemaStrategiesResult2 *PlaygroundFeatureSchemaStrategiesResult2

	Type PlaygroundFeatureSchemaStrategiesResultType
}

func CreatePlaygroundFeatureSchemaStrategiesResultBoolean(boolean bool) PlaygroundFeatureSchemaStrategiesResult {
	typ := PlaygroundFeatureSchemaStrategiesResultTypeBoolean

	return PlaygroundFeatureSchemaStrategiesResult{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreatePlaygroundFeatureSchemaStrategiesResultPlaygroundFeatureSchemaStrategiesResult2(playgroundFeatureSchemaStrategiesResult2 PlaygroundFeatureSchemaStrategiesResult2) PlaygroundFeatureSchemaStrategiesResult {
	typ := PlaygroundFeatureSchemaStrategiesResultTypePlaygroundFeatureSchemaStrategiesResult2

	return PlaygroundFeatureSchemaStrategiesResult{
		PlaygroundFeatureSchemaStrategiesResult2: &playgroundFeatureSchemaStrategiesResult2,
		Type:                                     typ,
	}
}

func (u *PlaygroundFeatureSchemaStrategiesResult) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	boolean := new(bool)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&boolean); err == nil {
		u.Boolean = boolean
		u.Type = PlaygroundFeatureSchemaStrategiesResultTypeBoolean
		return nil
	}

	playgroundFeatureSchemaStrategiesResult2 := new(PlaygroundFeatureSchemaStrategiesResult2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&playgroundFeatureSchemaStrategiesResult2); err == nil {
		u.PlaygroundFeatureSchemaStrategiesResult2 = playgroundFeatureSchemaStrategiesResult2
		u.Type = PlaygroundFeatureSchemaStrategiesResultTypePlaygroundFeatureSchemaStrategiesResult2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PlaygroundFeatureSchemaStrategiesResult) MarshalJSON() ([]byte, error) {
	if u.Boolean != nil {
		return json.Marshal(u.Boolean)
	}

	if u.PlaygroundFeatureSchemaStrategiesResult2 != nil {
		return json.Marshal(u.PlaygroundFeatureSchemaStrategiesResult2)
	}

	return nil, nil
}

// PlaygroundFeatureSchemaStrategies - The feature's applicable strategies and cumulative results of the strategies
type PlaygroundFeatureSchemaStrategies struct {
	// The strategies that apply to this feature.
	Data []PlaygroundStrategySchema `json:"data"`
	// The cumulative results of all the feature's strategies. Can be `true`,
	//                                   `false`, or `unknown`.
	//                                   This property will only be `unknown`
	//                                   if one or more of the strategies can't be fully evaluated and the rest of the strategies
	//                                   all resolve to `false`.
	Result PlaygroundFeatureSchemaStrategiesResult `json:"result"`
}

// PlaygroundFeatureSchemaVariantPayload - An optional payload attached to the variant.
type PlaygroundFeatureSchemaVariantPayload struct {
	// The format of the payload.
	Type string `json:"type"`
	// The payload value stringified.
	Value string `json:"value"`
}

// PlaygroundFeatureSchemaVariant - The feature variant you receive based on the provided context or the _disabled
//
//	variant_. If a feature is disabled or doesn't have any
//	variants, you would get the _disabled variant_.
//	Otherwise, you'll get one of thefeature's defined variants.
type PlaygroundFeatureSchemaVariant struct {
	// Whether the variant is enabled or not. If the feature is disabled or if it doesn't have variants, this property will be `false`
	Enabled bool `json:"enabled"`
	// The variant's name. If there is no variant or if the toggle is disabled, this will be `disabled`
	Name string `json:"name"`
	// An optional payload attached to the variant.
	Payload *PlaygroundFeatureSchemaVariantPayload `json:"payload,omitempty"`
}

// PlaygroundFeatureSchema - A simplified feature toggle model intended for the Unleash playground.
type PlaygroundFeatureSchema struct {
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
	// The feature's applicable strategies and cumulative results of the strategies
	Strategies PlaygroundFeatureSchemaStrategies `json:"strategies"`
	// The feature variant you receive based on the provided context or the _disabled
	//                           variant_. If a feature is disabled or doesn't have any
	//                           variants, you would get the _disabled variant_.
	//                           Otherwise, you'll get one of thefeature's defined variants.
	Variant PlaygroundFeatureSchemaVariant `json:"variant"`
	// The feature variants.
	Variants []VariantSchema `json:"variants"`
}
