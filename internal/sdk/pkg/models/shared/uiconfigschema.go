// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UIConfigSchemaAuthenticationType - The type of authentication enabled for this Unleash instance
type UIConfigSchemaAuthenticationType string

const (
	UIConfigSchemaAuthenticationTypeOpenSource UIConfigSchemaAuthenticationType = "open-source"
	UIConfigSchemaAuthenticationTypeDemo       UIConfigSchemaAuthenticationType = "demo"
	UIConfigSchemaAuthenticationTypeEnterprise UIConfigSchemaAuthenticationType = "enterprise"
	UIConfigSchemaAuthenticationTypeHosted     UIConfigSchemaAuthenticationType = "hosted"
	UIConfigSchemaAuthenticationTypeCustom     UIConfigSchemaAuthenticationType = "custom"
	UIConfigSchemaAuthenticationTypeNone       UIConfigSchemaAuthenticationType = "none"
)

func (e UIConfigSchemaAuthenticationType) ToPointer() *UIConfigSchemaAuthenticationType {
	return &e
}

func (e *UIConfigSchemaAuthenticationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "open-source":
		fallthrough
	case "demo":
		fallthrough
	case "enterprise":
		fallthrough
	case "hosted":
		fallthrough
	case "custom":
		fallthrough
	case "none":
		*e = UIConfigSchemaAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UIConfigSchemaAuthenticationType: %v", v)
	}
}

type UIConfigSchemaLinks struct {
}

// UIConfigSchema - A collection of properties used to configure the Unleash Admin UI.
type UIConfigSchema struct {
	// The type of authentication enabled for this Unleash instance
	AuthenticationType *UIConfigSchemaAuthenticationType `json:"authenticationType,omitempty"`
	// The base URI path at which this Unleash instance is listening.
	BaseURIPath string `json:"baseUriPath"`
	// Whether password authentication should be disabled or not.
	DisablePasswordAuth *bool `json:"disablePasswordAuth,omitempty"`
	// Whether this instance can send out emails or not.
	EmailEnabled *bool `json:"emailEnabled,omitempty"`
	// What kind of Unleash instance it is: Enterprise, Pro, or Open source
	Environment *string `json:"environment,omitempty"`
	// Additional (largely experimental) features that are enabled in this Unleash instance.
	Flags map[string]UIConfigSchemaFlagsValue `json:"flags,omitempty"`
	// The list of origins that the front-end API should accept requests from.
	FrontendAPIOrigins []string `json:"frontendApiOrigins,omitempty"`
	// Relevant links to use in the UI.
	Links []UIConfigSchemaLinks `json:"links,omitempty"`
	// Whether maintenance mode is currently active or not.
	MaintenanceMode *bool `json:"maintenanceMode,omitempty"`
	// The name of this Unleash instance. Used to build the text in the footer.
	Name *string `json:"name,omitempty"`
	// Whether to enable the Unleash network view or not.
	NetworkViewEnabled *bool `json:"networkViewEnabled,omitempty"`
	// The maximum number of values that can be used in a single segment.
	SegmentValuesLimit *float64 `json:"segmentValuesLimit,omitempty"`
	// The slogan to display in the UI footer.
	Slogan *string `json:"slogan,omitempty"`
	// The maximum number of segments that can be applied to a single strategy.
	StrategySegmentsLimit *float64 `json:"strategySegmentsLimit,omitempty"`
	// The URL of the Unleash instance.
	UnleashURL string `json:"unleashUrl"`
	// The current version of Unleash
	Version string `json:"version"`
	// Detailed information about an Unleash version
	VersionInfo VersionSchema `json:"versionInfo"`
}
