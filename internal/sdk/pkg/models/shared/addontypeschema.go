// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AddonTypeSchemaAlertsType - The type of alert. This determines the color of the alert.
type AddonTypeSchemaAlertsType string

const (
	AddonTypeSchemaAlertsTypeSuccess AddonTypeSchemaAlertsType = "success"
	AddonTypeSchemaAlertsTypeInfo    AddonTypeSchemaAlertsType = "info"
	AddonTypeSchemaAlertsTypeWarning AddonTypeSchemaAlertsType = "warning"
	AddonTypeSchemaAlertsTypeError   AddonTypeSchemaAlertsType = "error"
)

func (e AddonTypeSchemaAlertsType) ToPointer() *AddonTypeSchemaAlertsType {
	return &e
}

func (e *AddonTypeSchemaAlertsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "info":
		fallthrough
	case "warning":
		fallthrough
	case "error":
		*e = AddonTypeSchemaAlertsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddonTypeSchemaAlertsType: %v", v)
	}
}

type AddonTypeSchemaAlerts struct {
	// The text of the alert. This is what will be displayed to the user.
	Text string `json:"text"`
	// The type of alert. This determines the color of the alert.
	Type AddonTypeSchemaAlertsType `json:"type"`
}

// AddonTypeSchemaInstallation - The installation configuration for this addon type.
type AddonTypeSchemaInstallation struct {
	// The help text of the installation configuration. This will be displayed to the user when installing addons of this type.
	HelpText *string `json:"helpText,omitempty"`
	// The title of the installation configuration. This will be displayed to the user when installing addons of this type.
	Title *string `json:"title,omitempty"`
	// A URL to where the addon configuration should redirect to install addons of this type.
	URL string `json:"url"`
}

// AddonTypeSchema - An addon provider. Defines a specific addon type and what the end user must configure when creating a new addon of that type.
type AddonTypeSchema struct {
	// A list of alerts to display to the user when installing addons of this type.
	Alerts []AddonTypeSchemaAlerts `json:"alerts,omitempty"`
	// This should be used to inform the user that this addon type is deprecated and should not be used. Deprecated addons will show a badge with this information on the UI.
	Deprecated *string `json:"deprecated,omitempty"`
	// A description of the addon type.
	Description string `json:"description"`
	// The addon type's name as it should be displayed in the admin UI.
	DisplayName string `json:"displayName"`
	// A URL to where you can find more information about using this addon type.
	DocumentationURL string `json:"documentationUrl"`
	// All the [event types](https://docs.getunleash.io/reference/api/legacy/unleash/admin/events#feature-toggle-events) that are available for this addon provider.
	Events []string `json:"events,omitempty"`
	// The installation configuration for this addon type.
	Installation *AddonTypeSchemaInstallation `json:"installation,omitempty"`
	// The name of the addon type. When creating new addons, this goes in the payload's `type` field.
	Name string `json:"name"`
	// The addon provider's parameters. Use these to configure an addon of this provider type. Items with `required: true` must be provided.
	Parameters []AddonParameterSchema `json:"parameters,omitempty"`
	// A list of [Unleash tag types](https://docs.getunleash.io/reference/tags#tag-types) that this addon uses. These tags will be added to the Unleash instance when an addon of this type is created.
	TagTypes []TagTypeSchema `json:"tagTypes,omitempty"`
}
