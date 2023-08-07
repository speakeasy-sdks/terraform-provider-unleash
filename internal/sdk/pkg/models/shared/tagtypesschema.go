// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TagTypesSchema - A list of tag types with a version number representing the schema used to model the tag types.
type TagTypesSchema struct {
	// The list of tag types.
	TagTypes []TagTypeSchema `json:"tagTypes"`
	// The version of the schema used to model the tag types.
	Version int64 `json:"version"`
}