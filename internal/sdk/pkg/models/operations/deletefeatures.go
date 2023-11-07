// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type DeleteFeaturesRequest struct {
	// batchFeaturesSchema
	BatchFeaturesSchema shared.BatchFeaturesSchema `request:"mediaType=application/json"`
	ProjectID           string                     `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *DeleteFeaturesRequest) GetBatchFeaturesSchema() shared.BatchFeaturesSchema {
	if o == nil {
		return shared.BatchFeaturesSchema{}
	}
	return o.BatchFeaturesSchema
}

func (o *DeleteFeaturesRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// DeleteFeaturesArchiveResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type DeleteFeaturesArchiveResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteFeaturesArchiveResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteFeaturesArchiveResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteFeaturesArchiveResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteFeaturesArchiveResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type DeleteFeaturesArchiveResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteFeaturesArchiveResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteFeaturesArchiveResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteFeaturesArchiveResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteFeaturesResponseBody - The request data does not match what we expect.
type DeleteFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type DeleteFeaturesResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *DeleteFeaturesResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *DeleteFeaturesArchiveResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *DeleteFeaturesArchiveResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteFeaturesResponse) GetFourHundredApplicationJSONObject() *DeleteFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *DeleteFeaturesResponse) GetFourHundredAndOneApplicationJSONObject() *DeleteFeaturesArchiveResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *DeleteFeaturesResponse) GetFourHundredAndThreeApplicationJSONObject() *DeleteFeaturesArchiveResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *DeleteFeaturesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteFeaturesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteFeaturesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
