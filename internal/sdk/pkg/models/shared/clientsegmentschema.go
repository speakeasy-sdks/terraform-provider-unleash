// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ClientSegmentSchema - Represents a client API segment of users defined by a set of constraints.
type ClientSegmentSchema struct {
	// List of constraints that determine which users are part of the segment
	Constraints []ConstraintSchema `json:"constraints"`
	// The segment's id.
	ID float64 `json:"id"`
	// The name of the segment.
	Name *string `json:"name,omitempty"`
}
