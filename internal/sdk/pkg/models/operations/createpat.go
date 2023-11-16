// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

// CreatePatPersonalAccessTokensResponseResponseBody - The requested resource was not found.
type CreatePatPersonalAccessTokensResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreatePatPersonalAccessTokensResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreatePatPersonalAccessTokensResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreatePatPersonalAccessTokensResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreatePatPersonalAccessTokensResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type CreatePatPersonalAccessTokensResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreatePatPersonalAccessTokensResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreatePatPersonalAccessTokensResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreatePatPersonalAccessTokensResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreatePatResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type CreatePatResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreatePatResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreatePatResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreatePatResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type CreatePatResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *CreatePatResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *CreatePatPersonalAccessTokensResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *CreatePatPersonalAccessTokensResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The resource was successfully created.
	PatSchema *shared.PatSchema
}

func (o *CreatePatResponse) GetFourHundredAndOneApplicationJSONObject() *CreatePatResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *CreatePatResponse) GetFourHundredAndThreeApplicationJSONObject() *CreatePatPersonalAccessTokensResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *CreatePatResponse) GetFourHundredAndFourApplicationJSONObject() *CreatePatPersonalAccessTokensResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *CreatePatResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePatResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreatePatResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePatResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreatePatResponse) GetPatSchema() *shared.PatSchema {
	if o == nil {
		return nil
	}
	return o.PatSchema
}
