// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AddonParameterSchema - An addon parameter definition.
type AddonParameterSchema struct {
	// A description of the parameter. This should explain to the end user what the parameter is used for.
	Description *string `json:"description,omitempty"`
	// The name of the parameter as it is shown to the end user in the Admin UI.
	DisplayName string `json:"displayName"`
	// The name of the parameter as it is used in code. References to this parameter should use this value.
	Name string `json:"name"`
	// The default value for this parameter. This value is used if no other value is provided.
	Placeholder *string `json:"placeholder,omitempty"`
	// Whether this parameter is required or not. If a parameter is required, you must give it a value when you create the addon. If it is not required it can be left out. It may receive a default value in those cases.
	Required bool `json:"required"`
	// Indicates whether this parameter is **sensitive** or not. Unleash will not return sensitive parameters to API requests. It will instead use a number of asterisks to indicate that a value is set, e.g. "******". The number of asterisks does not correlate to the parameter's value.
	Sensitive bool `json:"sensitive"`
	// The type of the parameter. Corresponds roughly to [HTML `input` field types](https://developer.mozilla.org/docs/Web/HTML/Element/Input#input_types). Multi-line inut fields are indicated as `textfield` (equivalent to the HTML `textarea` tag).
	Type string `json:"type"`
}

func (o *AddonParameterSchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AddonParameterSchema) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *AddonParameterSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AddonParameterSchema) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *AddonParameterSchema) GetRequired() bool {
	if o == nil {
		return false
	}
	return o.Required
}

func (o *AddonParameterSchema) GetSensitive() bool {
	if o == nil {
		return false
	}
	return o.Sensitive
}

func (o *AddonParameterSchema) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
