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

// DeleteFeatures403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type DeleteFeatures403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteFeatures403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteFeatures403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteFeatures403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteFeatures401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type DeleteFeatures401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteFeatures401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteFeatures401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteFeatures401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteFeatures400ApplicationJSON - The request data does not match what we expect.
type DeleteFeatures400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteFeatures400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteFeatures400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteFeatures400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type DeleteFeaturesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request data does not match what we expect.
	DeleteFeatures400ApplicationJSONObject *DeleteFeatures400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	DeleteFeatures401ApplicationJSONObject *DeleteFeatures401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	DeleteFeatures403ApplicationJSONObject *DeleteFeatures403ApplicationJSON
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

func (o *DeleteFeaturesResponse) GetDeleteFeatures400ApplicationJSONObject() *DeleteFeatures400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteFeatures400ApplicationJSONObject
}

func (o *DeleteFeaturesResponse) GetDeleteFeatures401ApplicationJSONObject() *DeleteFeatures401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteFeatures401ApplicationJSONObject
}

func (o *DeleteFeaturesResponse) GetDeleteFeatures403ApplicationJSONObject() *DeleteFeatures403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteFeatures403ApplicationJSONObject
}
