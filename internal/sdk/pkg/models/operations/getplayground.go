// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

// GetPlaygroundPlaygroundResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetPlaygroundPlaygroundResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetPlaygroundPlaygroundResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetPlaygroundPlaygroundResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPlaygroundPlaygroundResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetPlaygroundResponseBody - The request data does not match what we expect.
type GetPlaygroundResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetPlaygroundResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetPlaygroundResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPlaygroundResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetPlaygroundResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *GetPlaygroundResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetPlaygroundPlaygroundResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// playgroundResponseSchema
	PlaygroundResponseSchema *shared.PlaygroundResponseSchema
}

func (o *GetPlaygroundResponse) GetFourHundredApplicationJSONObject() *GetPlaygroundResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *GetPlaygroundResponse) GetFourHundredAndOneApplicationJSONObject() *GetPlaygroundPlaygroundResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetPlaygroundResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPlaygroundResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPlaygroundResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPlaygroundResponse) GetPlaygroundResponseSchema() *shared.PlaygroundResponseSchema {
	if o == nil {
		return nil
	}
	return o.PlaygroundResponseSchema
}
