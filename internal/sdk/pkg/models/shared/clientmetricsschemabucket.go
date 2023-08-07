// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ClientMetricsSchemaBucket - Holds all metrics gathered over a window of time. Typically 1 hour wide
type ClientMetricsSchemaBucket struct {
	Start DateSchema `json:"start"`
	Stop  DateSchema `json:"stop"`
	// an object containing feature names with yes/no plus variant usage
	Toggles map[string]ClientMetricsSchemaBucketTogglesValue `json:"toggles"`
}