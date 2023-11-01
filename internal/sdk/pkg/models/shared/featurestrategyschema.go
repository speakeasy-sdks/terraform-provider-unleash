// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// FeatureStrategySchema - A single activation strategy configuration schema for a feature
type FeatureStrategySchema struct {
	// A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints
	Constraints []ConstraintSchema `json:"constraints,omitempty"`
	// A toggle to disable the strategy. defaults to false. Disabled strategies are not evaluated or returned to the SDKs
	Disabled *bool `json:"disabled,omitempty"`
	// The name or feature the strategy is attached to
	FeatureName *string `json:"featureName,omitempty"`
	// A uuid for the feature strategy
	ID *string `json:"id,omitempty"`
	// The name or type of strategy
	Name string `json:"name"`
	// A list of parameters for a strategy
	Parameters map[string]string `json:"parameters,omitempty"`
	// A list of segment ids attached to the strategy
	Segments []float64 `json:"segments,omitempty"`
	// The order of the strategy in the list
	SortOrder *float64 `json:"sortOrder,omitempty"`
	// A descriptive title for the strategy
	Title *string `json:"title,omitempty"`
	// Strategy level variants
	Variants []StrategyVariantSchema `json:"variants,omitempty"`
}

func (o *FeatureStrategySchema) GetConstraints() []ConstraintSchema {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *FeatureStrategySchema) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *FeatureStrategySchema) GetFeatureName() *string {
	if o == nil {
		return nil
	}
	return o.FeatureName
}

func (o *FeatureStrategySchema) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FeatureStrategySchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *FeatureStrategySchema) GetParameters() map[string]string {
	if o == nil {
		return nil
	}
	return o.Parameters
}

func (o *FeatureStrategySchema) GetSegments() []float64 {
	if o == nil {
		return nil
	}
	return o.Segments
}

func (o *FeatureStrategySchema) GetSortOrder() *float64 {
	if o == nil {
		return nil
	}
	return o.SortOrder
}

func (o *FeatureStrategySchema) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *FeatureStrategySchema) GetVariants() []StrategyVariantSchema {
	if o == nil {
		return nil
	}
	return o.Variants
}
