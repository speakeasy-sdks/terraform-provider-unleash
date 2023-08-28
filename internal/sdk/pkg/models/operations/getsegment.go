// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetSegmentRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

// GetSegment404ApplicationJSON - The requested resource was not found.
type GetSegment404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type GetSegmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// adminSegmentSchema
	AdminSegmentSchema *shared.AdminSegmentSchema
	// The requested resource was not found.
	GetSegment404ApplicationJSONObject *GetSegment404ApplicationJSON
}
