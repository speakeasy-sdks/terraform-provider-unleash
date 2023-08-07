// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

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
