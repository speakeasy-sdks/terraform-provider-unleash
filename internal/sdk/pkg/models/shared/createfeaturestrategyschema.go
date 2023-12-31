// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateFeatureStrategySchema - Create a strategy configuration in a feature
type CreateFeatureStrategySchema struct {
	// A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints
	Constraints []ConstraintSchema `json:"constraints,omitempty"`
	// A toggle to disable the strategy. defaults to false. Disabled strategies are not evaluated or returned to the SDKs
	Disabled *bool `json:"disabled,omitempty"`
	// The name of the strategy type
	Name string `json:"name"`
	// A list of parameters for a strategy
	Parameters map[string]string `json:"parameters,omitempty"`
	// Ids of segments to use for this strategy
	Segments []float64 `json:"segments,omitempty"`
	// The order of the strategy in the list
	SortOrder *float64 `json:"sortOrder,omitempty"`
	// A descriptive title for the strategy
	Title *string `json:"title,omitempty"`
	// Strategy level variants
	Variants []CreateStrategyVariantSchema `json:"variants,omitempty"`
}

func (o *CreateFeatureStrategySchema) GetConstraints() []ConstraintSchema {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *CreateFeatureStrategySchema) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *CreateFeatureStrategySchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateFeatureStrategySchema) GetParameters() map[string]string {
	if o == nil {
		return nil
	}
	return o.Parameters
}

func (o *CreateFeatureStrategySchema) GetSegments() []float64 {
	if o == nil {
		return nil
	}
	return o.Segments
}

func (o *CreateFeatureStrategySchema) GetSortOrder() *float64 {
	if o == nil {
		return nil
	}
	return o.SortOrder
}

func (o *CreateFeatureStrategySchema) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *CreateFeatureStrategySchema) GetVariants() []CreateStrategyVariantSchema {
	if o == nil {
		return nil
	}
	return o.Variants
}
