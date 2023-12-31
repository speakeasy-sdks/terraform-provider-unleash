// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// ImportTogglesResponseBody - The requested resource was not found.
type ImportTogglesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ImportTogglesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ImportTogglesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ImportTogglesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ImportTogglesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The requested resource was not found.
	Object *ImportTogglesResponseBody
}

func (o *ImportTogglesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ImportTogglesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ImportTogglesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ImportTogglesResponse) GetObject() *ImportTogglesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
