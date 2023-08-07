// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ValidateTagTypeSchema - The result of validating a tag type.
type ValidateTagTypeSchema struct {
	// A tag type.
	TagType TagTypeSchema `json:"tagType"`
	// Whether or not the tag type is valid.
	Valid bool `json:"valid"`
}