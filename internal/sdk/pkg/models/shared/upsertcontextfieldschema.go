// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpsertContextFieldSchema - upsertContextFieldSchema
type UpsertContextFieldSchema struct {
	Description *string            `json:"description,omitempty"`
	LegalValues []LegalValueSchema `json:"legalValues,omitempty"`
	Name        string             `json:"name"`
	SortOrder   *float64           `json:"sortOrder,omitempty"`
	Stickiness  *bool              `json:"stickiness,omitempty"`
}