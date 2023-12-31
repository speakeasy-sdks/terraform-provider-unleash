// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

// ValidateImportResponseBody - The requested resource was not found.
type ValidateImportResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateImportResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateImportResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateImportResponseBody) GetName() *string {
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
	Object *ValidateImportResponseBody
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

func (o *ValidateImportResponse) GetObject() *ValidateImportResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
