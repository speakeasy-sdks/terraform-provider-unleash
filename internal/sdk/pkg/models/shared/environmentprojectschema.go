// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EnvironmentProjectSchema - Describes a project's configuration in a given environment.
type EnvironmentProjectSchema struct {
	// Create a strategy configuration in a feature
	DefaultStrategy *CreateFeatureStrategySchema `json:"defaultStrategy,omitempty"`
	// `true` if the environment is enabled for the project, otherwise `false`
	Enabled bool `json:"enabled"`
	// The name of the environment
	Name string `json:"name"`
	// The number of client and front-end API tokens that have access to this project
	ProjectAPITokenCount *int64 `json:"projectApiTokenCount,omitempty"`
	// The number of features enabled in this environment for this project
	ProjectEnabledToggleCount *int64 `json:"projectEnabledToggleCount,omitempty"`
	// `true` if the environment is protected, otherwise `false`. A *protected* environment can not be deleted.
	Protected bool `json:"protected"`
	// Priority of the environment in a list of environments, the lower the value, the higher up in the list the environment will appear
	SortOrder int64 `json:"sortOrder"`
	// The [type of environment](https://docs.getunleash.io/reference/environments#environment-types).
	Type string `json:"type"`
}

func (o *EnvironmentProjectSchema) GetDefaultStrategy() *CreateFeatureStrategySchema {
	if o == nil {
		return nil
	}
	return o.DefaultStrategy
}

func (o *EnvironmentProjectSchema) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *EnvironmentProjectSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *EnvironmentProjectSchema) GetProjectAPITokenCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ProjectAPITokenCount
}

func (o *EnvironmentProjectSchema) GetProjectEnabledToggleCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ProjectEnabledToggleCount
}

func (o *EnvironmentProjectSchema) GetProtected() bool {
	if o == nil {
		return false
	}
	return o.Protected
}

func (o *EnvironmentProjectSchema) GetSortOrder() int64 {
	if o == nil {
		return 0
	}
	return o.SortOrder
}

func (o *EnvironmentProjectSchema) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
