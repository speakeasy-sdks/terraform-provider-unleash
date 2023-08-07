// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ImportTogglesValidateItemSchema - A description of an error or warning pertaining to a feature toggle import job.
type ImportTogglesValidateItemSchema struct {
	// The items affected by this error message
	AffectedItems []string `json:"affectedItems"`
	// The validation error message
	Message string `json:"message"`
}