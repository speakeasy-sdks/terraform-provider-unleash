// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AddonSchema - An [addon](https://docs.getunleash.io/reference/addons) instance description. Contains data about what kind of provider it uses, whether it's enabled or not, what events it listens for, and more.
type AddonSchema struct {
	// A description of the addon. `null` if no description exists.
	Description *string `json:"description"`
	// Whether the addon is enabled or not.
	Enabled bool `json:"enabled"`
	// The list of environments that this addon listens to events from. An empty list means it listens to events from **all** environments.
	Environments []string `json:"environments,omitempty"`
	// The event types that trigger this specific addon.
	Events []string `json:"events"`
	// The addon's unique identifier.
	ID int64 `json:"id"`
	// Parameters for the addon provider. This object has different required and optional properties depending on the provider you choose.
	Parameters map[string]interface{} `json:"parameters"`
	// The projects that this addon listens to events from. An empty list means it listens to events from **all** projects.
	Projects []string `json:"projects,omitempty"`
	// The addon provider, such as "webhook" or "slack".
	Provider string `json:"provider"`
}

func (o *AddonSchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AddonSchema) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *AddonSchema) GetEnvironments() []string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *AddonSchema) GetEvents() []string {
	if o == nil {
		return []string{}
	}
	return o.Events
}

func (o *AddonSchema) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *AddonSchema) GetParameters() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Parameters
}

func (o *AddonSchema) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *AddonSchema) GetProvider() string {
	if o == nil {
		return ""
	}
	return o.Provider
}
