// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ImportTogglesValidateItemSchema - A description of an error or warning pertaining to a feature toggle import job.
type ImportTogglesValidateItemSchema struct {
	// The items affected by this error message
	AffectedItems []string `json:"affectedItems"`
	// The validation error message
	Message string `json:"message"`
}

func (o *ImportTogglesValidateItemSchema) GetAffectedItems() []string {
	if o == nil {
		return []string{}
	}
	return o.AffectedItems
}

func (o *ImportTogglesValidateItemSchema) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}
