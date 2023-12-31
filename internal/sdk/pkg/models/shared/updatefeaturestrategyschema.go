// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateFeatureStrategySchema - Update a strategy configuration in a feature
type UpdateFeatureStrategySchema struct {
	// A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints
	Constraints []ConstraintSchema `json:"constraints,omitempty"`
	// A toggle to disable the strategy. defaults to true. Disabled strategies are not evaluated or returned to the SDKs
	Disabled *bool `json:"disabled,omitempty"`
	// The name of the strategy type
	Name *string `json:"name,omitempty"`
	// A list of parameters for a strategy
	Parameters map[string]string `json:"parameters,omitempty"`
	// The order of the strategy in the list in feature environment configuration
	SortOrder *float64 `json:"sortOrder,omitempty"`
	// A descriptive title for the strategy
	Title *string `json:"title,omitempty"`
}

func (o *UpdateFeatureStrategySchema) GetConstraints() []ConstraintSchema {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *UpdateFeatureStrategySchema) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *UpdateFeatureStrategySchema) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateFeatureStrategySchema) GetParameters() map[string]string {
	if o == nil {
		return nil
	}
	return o.Parameters
}

func (o *UpdateFeatureStrategySchema) GetSortOrder() *float64 {
	if o == nil {
		return nil
	}
	return o.SortOrder
}

func (o *UpdateFeatureStrategySchema) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}
