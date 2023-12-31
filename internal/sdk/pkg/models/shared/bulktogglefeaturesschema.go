// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// BulkToggleFeaturesSchema - The feature list used for bulk toggle operations
type BulkToggleFeaturesSchema struct {
	// The features that we want to bulk toggle
	Features []string `json:"features"`
}

func (o *BulkToggleFeaturesSchema) GetFeatures() []string {
	if o == nil {
		return []string{}
	}
	return o.Features
}
