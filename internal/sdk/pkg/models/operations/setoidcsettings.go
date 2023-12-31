// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

// SetOidcSettingsAuthResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type SetOidcSettingsAuthResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetOidcSettingsAuthResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetOidcSettingsAuthResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetOidcSettingsAuthResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetOidcSettingsAuthResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type SetOidcSettingsAuthResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetOidcSettingsAuthResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetOidcSettingsAuthResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetOidcSettingsAuthResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetOidcSettingsAuthResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type SetOidcSettingsAuthResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetOidcSettingsAuthResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetOidcSettingsAuthResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetOidcSettingsAuthResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetOidcSettingsResponseBody - The request data does not match what we expect.
type SetOidcSettingsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetOidcSettingsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetOidcSettingsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetOidcSettingsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type SetOidcSettingsResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *SetOidcSettingsResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *SetOidcSettingsAuthResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *SetOidcSettingsAuthResponseResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *SetOidcSettingsAuthResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// oidcSettingsSchema
	OidcSettingsSchema *shared.OidcSettingsSchema
}

func (o *SetOidcSettingsResponse) GetFourHundredApplicationJSONObject() *SetOidcSettingsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *SetOidcSettingsResponse) GetFourHundredAndOneApplicationJSONObject() *SetOidcSettingsAuthResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *SetOidcSettingsResponse) GetFourHundredAndThreeApplicationJSONObject() *SetOidcSettingsAuthResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *SetOidcSettingsResponse) GetFourHundredAndFifteenApplicationJSONObject() *SetOidcSettingsAuthResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *SetOidcSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetOidcSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetOidcSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SetOidcSettingsResponse) GetOidcSettingsSchema() *shared.OidcSettingsSchema {
	if o == nil {
		return nil
	}
	return o.OidcSettingsSchema
}
