// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ClientFeatureSchema - Feature toggle configuration used by SDKs to evaluate state of a toggle
type ClientFeatureSchema struct {
	// A description of the toggle
	Description *string `json:"description,omitempty"`
	// Whether the feature flag is enabled for the current API key or not. This is ANDed with the evaluation results of the strategies list, so if this is false, the evaluation result will always be false
	Enabled bool `json:"enabled"`
	// Set to true if SDKs should trigger [impression events](https://docs.getunleash.io/reference/impression-data) when this toggle is evaluated
	ImpressionData *bool `json:"impressionData,omitempty"`
	// The unique name of a feature toggle. Is validated to be URL safe on creation
	Name string `json:"name"`
	// Which project this feature toggle belongs to
	Project *string `json:"project,omitempty"`
	// If this is true Unleash believes this feature toggle has been active longer than Unleash expects a toggle of this type to be active
	Stale *bool `json:"stale,omitempty"`
	// Evaluation strategies for this toggle. Each entry in this list will be evaluated and ORed together
	Strategies []FeatureStrategySchema `json:"strategies,omitempty"`
	// What kind of feature flag is this. Refer to the documentation on [feature toggle types](https://docs.getunleash.io/reference/feature-toggle-types) for more information
	Type *string `json:"type,omitempty"`
	// [Variants](https://docs.getunleash.io/reference/feature-toggle-variants#what-are-variants) configured for this toggle
	Variants []VariantSchema `json:"variants,omitempty"`
}

func (o *ClientFeatureSchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ClientFeatureSchema) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *ClientFeatureSchema) GetImpressionData() *bool {
	if o == nil {
		return nil
	}
	return o.ImpressionData
}

func (o *ClientFeatureSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ClientFeatureSchema) GetProject() *string {
	if o == nil {
		return nil
	}
	return o.Project
}

func (o *ClientFeatureSchema) GetStale() *bool {
	if o == nil {
		return nil
	}
	return o.Stale
}

func (o *ClientFeatureSchema) GetStrategies() []FeatureStrategySchema {
	if o == nil {
		return nil
	}
	return o.Strategies
}

func (o *ClientFeatureSchema) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ClientFeatureSchema) GetVariants() []VariantSchema {
	if o == nil {
		return nil
	}
	return o.Variants
}
