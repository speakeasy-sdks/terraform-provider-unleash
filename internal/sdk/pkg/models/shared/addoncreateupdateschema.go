// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AddonCreateUpdateSchema - Data required to create or update an [Unleash addon](https://docs.getunleash.io/reference/addons) instance.
type AddonCreateUpdateSchema struct {
	// A description of the addon.
	Description *string `json:"description,omitempty"`
	// Whether the addon should be enabled or not.
	Enabled bool `json:"enabled"`
	// The list of environments that this addon will listen to events from. An empty list means it will listen to events from **all** environments.
	Environments []string `json:"environments,omitempty"`
	// The event types that will trigger this specific addon.
	Events []string `json:"events"`
	// Parameters for the addon provider. This object has different required and optional properties depending on the provider you choose. Consult the documentation for details.
	Parameters map[string]interface{} `json:"parameters"`
	// The projects that this addon will listen to events from. An empty list means it will listen to events from **all** projects.
	Projects []string `json:"projects,omitempty"`
	// The addon provider, such as "webhook" or "slack". This string is **case sensitive** and maps to the provider's `name` property.
	//
	// The list of all supported providers and their parameters for a specific Unleash instance can be found by making a GET request to the `api/admin/addons` endpoint: the `providers` property of that response will contain all available providers.
	//
	// The default set of providers can be found in the [addons reference documentation](https://docs.getunleash.io/reference/addons). The default supported options are:
	// - `datadog` for [Datadog](https://docs.getunleash.io/reference/addons/datadog)
	// - `slack` for [Slack](https://docs.getunleash.io/reference/addons/slack)
	// - `teams` for [Microsoft Teams](https://docs.getunleash.io/reference/addons/teams)
	// - `webhook` for [webhooks](https://docs.getunleash.io/reference/addons/webhook)
	//
	// The provider you choose for your addon dictates what properties the `parameters` object needs. Refer to the documentation for each provider for more information.
	//
	Provider string `json:"provider"`
}

func (o *AddonCreateUpdateSchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AddonCreateUpdateSchema) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *AddonCreateUpdateSchema) GetEnvironments() []string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *AddonCreateUpdateSchema) GetEvents() []string {
	if o == nil {
		return []string{}
	}
	return o.Events
}

func (o *AddonCreateUpdateSchema) GetParameters() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Parameters
}

func (o *AddonCreateUpdateSchema) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *AddonCreateUpdateSchema) GetProvider() string {
	if o == nil {
		return ""
	}
	return o.Provider
}
