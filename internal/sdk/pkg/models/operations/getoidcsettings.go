// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

// GetOidcSettingsAuthResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetOidcSettingsAuthResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetOidcSettingsAuthResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetOidcSettingsAuthResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetOidcSettingsAuthResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetOidcSettingsAuthResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetOidcSettingsAuthResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetOidcSettingsAuthResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetOidcSettingsAuthResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetOidcSettingsAuthResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetOidcSettingsResponseBody - The request data does not match what we expect.
type GetOidcSettingsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetOidcSettingsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetOidcSettingsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetOidcSettingsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetOidcSettingsResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *GetOidcSettingsResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetOidcSettingsAuthResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *GetOidcSettingsAuthResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// oidcSettingsSchema
	OidcSettingsSchema *shared.OidcSettingsSchema
}

func (o *GetOidcSettingsResponse) GetFourHundredApplicationJSONObject() *GetOidcSettingsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *GetOidcSettingsResponse) GetFourHundredAndOneApplicationJSONObject() *GetOidcSettingsAuthResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetOidcSettingsResponse) GetFourHundredAndThreeApplicationJSONObject() *GetOidcSettingsAuthResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *GetOidcSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOidcSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOidcSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetOidcSettingsResponse) GetOidcSettingsSchema() *shared.OidcSettingsSchema {
	if o == nil {
		return nil
	}
	return o.OidcSettingsSchema
}
