// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type ArchiveFeaturesRequest struct {
	// batchFeaturesSchema
	BatchFeaturesSchema shared.BatchFeaturesSchema `request:"mediaType=application/json"`
	ProjectID           string                     `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *ArchiveFeaturesRequest) GetBatchFeaturesSchema() shared.BatchFeaturesSchema {
	if o == nil {
		return shared.BatchFeaturesSchema{}
	}
	return o.BatchFeaturesSchema
}

func (o *ArchiveFeaturesRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// ArchiveFeatures415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type ArchiveFeatures415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ArchiveFeatures415ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ArchiveFeatures415ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ArchiveFeatures415ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ArchiveFeatures403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type ArchiveFeatures403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ArchiveFeatures403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ArchiveFeatures403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ArchiveFeatures403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ArchiveFeatures401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ArchiveFeatures401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ArchiveFeatures401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ArchiveFeatures401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ArchiveFeatures401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ArchiveFeatures400ApplicationJSON - The request data does not match what we expect.
type ArchiveFeatures400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ArchiveFeatures400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ArchiveFeatures400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ArchiveFeatures400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ArchiveFeaturesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request data does not match what we expect.
	ArchiveFeatures400ApplicationJSONObject *ArchiveFeatures400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	ArchiveFeatures401ApplicationJSONObject *ArchiveFeatures401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	ArchiveFeatures403ApplicationJSONObject *ArchiveFeatures403ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	ArchiveFeatures415ApplicationJSONObject *ArchiveFeatures415ApplicationJSON
}

func (o *ArchiveFeaturesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ArchiveFeaturesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ArchiveFeaturesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ArchiveFeaturesResponse) GetArchiveFeatures400ApplicationJSONObject() *ArchiveFeatures400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ArchiveFeatures400ApplicationJSONObject
}

func (o *ArchiveFeaturesResponse) GetArchiveFeatures401ApplicationJSONObject() *ArchiveFeatures401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ArchiveFeatures401ApplicationJSONObject
}

func (o *ArchiveFeaturesResponse) GetArchiveFeatures403ApplicationJSONObject() *ArchiveFeatures403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ArchiveFeatures403ApplicationJSONObject
}

func (o *ArchiveFeaturesResponse) GetArchiveFeatures415ApplicationJSONObject() *ArchiveFeatures415ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ArchiveFeatures415ApplicationJSONObject
}
