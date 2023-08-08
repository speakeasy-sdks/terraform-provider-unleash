// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ImportTogglesValidateSchema - An object containing [feature import](https://docs.getunleash.io/reference/deploy/environment-import-export) validation results.
type ImportTogglesValidateSchema struct {
	// A list of errors that prevent the provided data from being successfully imported.
	Errors []ImportTogglesValidateItemSchema `json:"errors"`
	// Any additional permissions required to import the data. If the list is empty, you require no additional permissions beyond what your user already has.
	Permissions []ImportTogglesValidateItemSchema `json:"permissions,omitempty"`
	// A list of warnings related to the provided data.
	Warnings []ImportTogglesValidateItemSchema `json:"warnings"`
}
