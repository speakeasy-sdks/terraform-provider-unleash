// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"terraform/internal/sdk/pkg/utils"
)

// Links - A set of links to actions you can perform on this strategy
type Links struct {
	Edit string `json:"edit"`
}

func (o *Links) GetEdit() string {
	if o == nil {
		return ""
	}
	return o.Edit
}

// PlaygroundStrategySchemaEvaluationStatus - Signals that this strategy was evaluated successfully.
type PlaygroundStrategySchemaEvaluationStatus string

const (
	PlaygroundStrategySchemaEvaluationStatusComplete PlaygroundStrategySchemaEvaluationStatus = "complete"
)

func (e PlaygroundStrategySchemaEvaluationStatus) ToPointer() *PlaygroundStrategySchemaEvaluationStatus {
	return &e
}

func (e *PlaygroundStrategySchemaEvaluationStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "complete":
		*e = PlaygroundStrategySchemaEvaluationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaygroundStrategySchemaEvaluationStatus: %v", v)
	}
}

// PlaygroundStrategySchemaType - The format of the payload.
type PlaygroundStrategySchemaType string

const (
	PlaygroundStrategySchemaTypeJSON   PlaygroundStrategySchemaType = "json"
	PlaygroundStrategySchemaTypeCsv    PlaygroundStrategySchemaType = "csv"
	PlaygroundStrategySchemaTypeString PlaygroundStrategySchemaType = "string"
)

func (e PlaygroundStrategySchemaType) ToPointer() *PlaygroundStrategySchemaType {
	return &e
}

func (e *PlaygroundStrategySchemaType) UnmarshalJSON(data []byte) error {
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
		*e = PlaygroundStrategySchemaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaygroundStrategySchemaType: %v", v)
	}
}

// PlaygroundStrategySchemaPayload - An optional payload attached to the variant.
type PlaygroundStrategySchemaPayload struct {
	// The format of the payload.
	Type PlaygroundStrategySchemaType `json:"type"`
	// The payload value stringified.
	Value string `json:"value"`
}

func (o *PlaygroundStrategySchemaPayload) GetType() PlaygroundStrategySchemaType {
	if o == nil {
		return PlaygroundStrategySchemaType("")
	}
	return o.Type
}

func (o *PlaygroundStrategySchemaPayload) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// PlaygroundStrategySchemaVariant - The feature variant you receive based on the provided context or the _disabled
//
//	variant_. If a feature is disabled or doesn't have any
//	variants, you would get the _disabled variant_.
//	Otherwise, you'll get one of the feature's defined variants.
type PlaygroundStrategySchemaVariant struct {
	// Whether the variant is enabled or not. If the feature is disabled or if it doesn't have variants, this property will be `false`
	Enabled bool `json:"enabled"`
	// The variant's name. If there is no variant or if the toggle is disabled, this will be `disabled`
	Name string `json:"name"`
	// An optional payload attached to the variant.
	Payload *PlaygroundStrategySchemaPayload `json:"payload,omitempty"`
}

func (o *PlaygroundStrategySchemaVariant) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *PlaygroundStrategySchemaVariant) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PlaygroundStrategySchemaVariant) GetPayload() *PlaygroundStrategySchemaPayload {
	if o == nil {
		return nil
	}
	return o.Payload
}

// PlaygroundStrategySchema2 - The strategy's evaluation result. If the strategy is a custom strategy that Unleash can't evaluate, `evaluationStatus` will be `unknown`. Otherwise, it will be `true` or `false`
type PlaygroundStrategySchema2 struct {
	// Whether this strategy evaluates to true or not.
	Enabled bool `json:"enabled"`
	// Signals that this strategy was evaluated successfully.
	EvaluationStatus PlaygroundStrategySchemaEvaluationStatus `json:"evaluationStatus"`
	// The feature variant you receive based on the provided context or the _disabled
	//                           variant_. If a feature is disabled or doesn't have any
	//                           variants, you would get the _disabled variant_.
	//                           Otherwise, you'll get one of the feature's defined variants.
	Variant *PlaygroundStrategySchemaVariant `json:"variant,omitempty"`
	// The feature variants.
	Variants []VariantSchema `json:"variants,omitempty"`
}

func (o *PlaygroundStrategySchema2) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *PlaygroundStrategySchema2) GetEvaluationStatus() PlaygroundStrategySchemaEvaluationStatus {
	if o == nil {
		return PlaygroundStrategySchemaEvaluationStatus("")
	}
	return o.EvaluationStatus
}

func (o *PlaygroundStrategySchema2) GetVariant() *PlaygroundStrategySchemaVariant {
	if o == nil {
		return nil
	}
	return o.Variant
}

func (o *PlaygroundStrategySchema2) GetVariants() []VariantSchema {
	if o == nil {
		return nil
	}
	return o.Variants
}

// PlaygroundStrategySchemaSchemas2 - Whether this strategy resolves to `false` or if it might resolve to `true`. Because Unleash can't evaluate the strategy, it can't say for certain whether it will be `true`, but if you have failing constraints or segments, it _can_ determine that your strategy would be `false`.
type PlaygroundStrategySchemaSchemas2 string

const (
	PlaygroundStrategySchemaSchemas2Unknown PlaygroundStrategySchemaSchemas2 = "unknown"
)

func (e PlaygroundStrategySchemaSchemas2) ToPointer() *PlaygroundStrategySchemaSchemas2 {
	return &e
}

func (e *PlaygroundStrategySchemaSchemas2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unknown":
		*e = PlaygroundStrategySchemaSchemas2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlaygroundStrategySchemaSchemas2: %v", v)
	}
}

type EnabledType string

const (
	EnabledTypeBoolean                          EnabledType = "boolean"
	EnabledTypePlaygroundStrategySchemaSchemas2 EnabledType = "playgroundStrategySchema_Schemas_2"
)

type Enabled struct {
	Boolean                          *bool
	PlaygroundStrategySchemaSchemas2 *PlaygroundStrategySchemaSchemas2

	Type EnabledType
}

func CreateEnabledBoolean(boolean bool) Enabled {
	typ := EnabledTypeBoolean

	return Enabled{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateEnabledPlaygroundStrategySchemaSchemas2(playgroundStrategySchemaSchemas2 PlaygroundStrategySchemaSchemas2) Enabled {
	typ := EnabledTypePlaygroundStrategySchemaSchemas2

	return Enabled{
		PlaygroundStrategySchemaSchemas2: &playgroundStrategySchemaSchemas2,
		Type:                             typ,
	}
}

func (u *Enabled) UnmarshalJSON(data []byte) error {

	boolean := new(bool)
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = boolean
		u.Type = EnabledTypeBoolean
		return nil
	}

	playgroundStrategySchemaSchemas2 := new(PlaygroundStrategySchemaSchemas2)
	if err := utils.UnmarshalJSON(data, &playgroundStrategySchemaSchemas2, "", true, true); err == nil {
		u.PlaygroundStrategySchemaSchemas2 = playgroundStrategySchemaSchemas2
		u.Type = EnabledTypePlaygroundStrategySchemaSchemas2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Enabled) MarshalJSON() ([]byte, error) {
	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.PlaygroundStrategySchemaSchemas2 != nil {
		return utils.MarshalJSON(u.PlaygroundStrategySchemaSchemas2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// EvaluationStatus - Signals that this strategy could not be evaluated. This is most likely because you're using a custom strategy that Unleash doesn't know about.
type EvaluationStatus string

const (
	EvaluationStatusIncomplete EvaluationStatus = "incomplete"
)

func (e EvaluationStatus) ToPointer() *EvaluationStatus {
	return &e
}

func (e *EvaluationStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "incomplete":
		*e = EvaluationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EvaluationStatus: %v", v)
	}
}

// One - The strategy's evaluation result. If the strategy is a custom strategy that Unleash can't evaluate, `evaluationStatus` will be `unknown`. Otherwise, it will be `true` or `false`
type One struct {
	// Whether this strategy resolves to `false` or if it might resolve to `true`. Because Unleash can't evaluate the strategy, it can't say for certain whether it will be `true`, but if you have failing constraints or segments, it _can_ determine that your strategy would be `false`.
	Enabled Enabled `json:"enabled"`
	// Signals that this strategy could not be evaluated. This is most likely because you're using a custom strategy that Unleash doesn't know about.
	EvaluationStatus EvaluationStatus `json:"evaluationStatus"`
}

func (o *One) GetEnabled() Enabled {
	if o == nil {
		return Enabled{}
	}
	return o.Enabled
}

func (o *One) GetEvaluationStatus() EvaluationStatus {
	if o == nil {
		return EvaluationStatus("")
	}
	return o.EvaluationStatus
}

type ResultUnionType string

const (
	ResultUnionTypeOne                       ResultUnionType = "1"
	ResultUnionTypePlaygroundStrategySchema2 ResultUnionType = "playgroundStrategySchema_2"
)

type Result struct {
	One                       *One
	PlaygroundStrategySchema2 *PlaygroundStrategySchema2

	Type ResultUnionType
}

func CreateResultOne(one One) Result {
	typ := ResultUnionTypeOne

	return Result{
		One:  &one,
		Type: typ,
	}
}

func CreateResultPlaygroundStrategySchema2(playgroundStrategySchema2 PlaygroundStrategySchema2) Result {
	typ := ResultUnionTypePlaygroundStrategySchema2

	return Result{
		PlaygroundStrategySchema2: &playgroundStrategySchema2,
		Type:                      typ,
	}
}

func (u *Result) UnmarshalJSON(data []byte) error {

	one := new(One)
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = one
		u.Type = ResultUnionTypeOne
		return nil
	}

	playgroundStrategySchema2 := new(PlaygroundStrategySchema2)
	if err := utils.UnmarshalJSON(data, &playgroundStrategySchema2, "", true, true); err == nil {
		u.PlaygroundStrategySchema2 = playgroundStrategySchema2
		u.Type = ResultUnionTypePlaygroundStrategySchema2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Result) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.PlaygroundStrategySchema2 != nil {
		return utils.MarshalJSON(u.PlaygroundStrategySchema2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type PlaygroundStrategySchema struct {
	// The strategy's constraints and their evaluation results.
	Constraints []PlaygroundConstraintSchema `json:"constraints"`
	// The strategy's status. Disabled strategies are not evaluated
	Disabled *bool `json:"disabled"`
	// The strategy's id.
	ID string `json:"id"`
	// A set of links to actions you can perform on this strategy
	Links Links `json:"links"`
	// The strategy's name.
	Name string `json:"name"`
	// A list of parameters for a strategy
	Parameters map[string]string `json:"parameters"`
	// The strategy's evaluation result. If the strategy is a custom strategy that Unleash can't evaluate, `evaluationStatus` will be `unknown`. Otherwise, it will be `true` or `false`
	Result Result `json:"result"`
	// The strategy's segments and their evaluation results.
	Segments []PlaygroundSegmentSchema `json:"segments"`
	// Description of the feature's purpose.
	Title *string `json:"title,omitempty"`
}

func (o *PlaygroundStrategySchema) GetConstraints() []PlaygroundConstraintSchema {
	if o == nil {
		return []PlaygroundConstraintSchema{}
	}
	return o.Constraints
}

func (o *PlaygroundStrategySchema) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *PlaygroundStrategySchema) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PlaygroundStrategySchema) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *PlaygroundStrategySchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PlaygroundStrategySchema) GetParameters() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Parameters
}

func (o *PlaygroundStrategySchema) GetResult() Result {
	if o == nil {
		return Result{}
	}
	return o.Result
}

func (o *PlaygroundStrategySchema) GetSegments() []PlaygroundSegmentSchema {
	if o == nil {
		return []PlaygroundSegmentSchema{}
	}
	return o.Segments
}

func (o *PlaygroundStrategySchema) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}
