// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ClientMetricsSchema - Client usage metrics, accumulated in buckets of hour by hour by default
type ClientMetricsSchema struct {
	// The name of the application that is evaluating toggles
	AppName string `json:"appName"`
	// Holds all metrics gathered over a window of time. Typically 1 hour wide
	Bucket ClientMetricsSchemaBucket `json:"bucket"`
	// Which environment the application is running in
	Environment *string `json:"environment,omitempty"`
	// A [(somewhat) unique identifier](https://docs.getunleash.io/reference/sdks/node#advanced-usage) for the application
	InstanceID *string `json:"instanceId,omitempty"`
}
