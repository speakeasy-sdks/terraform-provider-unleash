// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContextFieldStrategiesSchemaStrategies struct {
	// The ID of the environment where this strategy is in.
	Environment string `json:"environment"`
	// The name of the feature that contains this strategy.
	FeatureName string `json:"featureName"`
	// The ID of the strategy.
	ID string `json:"id"`
	// The ID of the project that contains this feature.
	ProjectID string `json:"projectId"`
	// The name of the strategy.
	StrategyName string `json:"strategyName"`
}

func (o *ContextFieldStrategiesSchemaStrategies) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *ContextFieldStrategiesSchemaStrategies) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *ContextFieldStrategiesSchemaStrategies) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ContextFieldStrategiesSchemaStrategies) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *ContextFieldStrategiesSchemaStrategies) GetStrategyName() string {
	if o == nil {
		return ""
	}
	return o.StrategyName
}

// ContextFieldStrategiesSchema - A wrapper object containing all strategies that use a specific context field
type ContextFieldStrategiesSchema struct {
	// List of strategies using the context field
	Strategies []ContextFieldStrategiesSchemaStrategies `json:"strategies"`
}

func (o *ContextFieldStrategiesSchema) GetStrategies() []ContextFieldStrategiesSchemaStrategies {
	if o == nil {
		return []ContextFieldStrategiesSchemaStrategies{}
	}
	return o.Strategies
}
