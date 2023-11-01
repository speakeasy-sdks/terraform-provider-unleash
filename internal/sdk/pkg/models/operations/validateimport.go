// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// ValidateImport404ApplicationJSON - The requested resource was not found.
type ValidateImport404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateImport404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateImport404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateImport404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ValidateImportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// importTogglesValidateSchema
	ImportTogglesValidateSchema *shared.ImportTogglesValidateSchema
	// The requested resource was not found.
	ValidateImport404ApplicationJSONObject *ValidateImport404ApplicationJSON
}

func (o *ValidateImportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ValidateImportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ValidateImportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ValidateImportResponse) GetImportTogglesValidateSchema() *shared.ImportTogglesValidateSchema {
	if o == nil {
		return nil
	}
	return o.ImportTogglesValidateSchema
}

func (o *ValidateImportResponse) GetValidateImport404ApplicationJSONObject() *ValidateImport404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ValidateImport404ApplicationJSONObject
}
