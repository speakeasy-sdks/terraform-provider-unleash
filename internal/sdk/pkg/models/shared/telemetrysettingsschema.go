// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TelemetrySettingsSchema - Contains information about which settings are configured for version info collection and feature usage collection.
type TelemetrySettingsSchema struct {
	// Whether collection of feature usage metrics is enabled/active.
	FeatureInfoCollectionEnabled bool `json:"featureInfoCollectionEnabled"`
	// Whether collection of version info is enabled/active.
	VersionInfoCollectionEnabled bool `json:"versionInfoCollectionEnabled"`
}
