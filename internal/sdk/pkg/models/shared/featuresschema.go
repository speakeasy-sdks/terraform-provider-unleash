// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// FeaturesSchema - A list of features
//
// @deprecated null: This will be removed in a future release, please migrate away from it as soon as possible.
type FeaturesSchema struct {
	// A list of features
	Features []FeatureSchema `json:"features"`
	// The version of the feature's schema
	Version int64 `json:"version"`
}