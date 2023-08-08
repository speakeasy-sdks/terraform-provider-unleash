// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PlaygroundStrategySchema struct {
	// The strategy's constraints and their evaluation results.
	Constraints []PlaygroundConstraintSchema `json:"constraints"`
	// The strategy's status. Disabled strategies are not evaluated
	Disabled bool `json:"disabled"`
	// The strategy's id.
	ID string `json:"id"`
	// A set of links to actions you can perform on this strategy
	Links PlaygroundStrategySchemaLinks `json:"links"`
	// The strategy's name.
	Name string `json:"name"`
	// A list of parameters for a strategy
	Parameters map[string]string `json:"parameters"`
	// The strategy's evaluation result. If the strategy is a custom strategy that Unleash can't evaluate, `evaluationStatus` will be `unknown`. Otherwise, it will be `true` or `false`
	Result PlaygroundStrategySchemaResult `json:"result"`
	// The strategy's segments and their evaluation results.
	Segments []PlaygroundSegmentSchema `json:"segments"`
	// Description of the feature's purpose.
	Title *string `json:"title,omitempty"`
}
