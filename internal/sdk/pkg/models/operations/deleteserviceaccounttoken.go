// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteServiceAccountTokenRequest struct {
	ID      string `pathParam:"style=simple,explode=false,name=id"`
	TokenID string `pathParam:"style=simple,explode=false,name=tokenId"`
}

func (o *DeleteServiceAccountTokenRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteServiceAccountTokenRequest) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

// DeleteServiceAccountTokenServiceAccountsResponseResponseBody - The requested resource was not found.
type DeleteServiceAccountTokenServiceAccountsResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteServiceAccountTokenServiceAccountsResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteServiceAccountTokenServiceAccountsResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteServiceAccountTokenServiceAccountsResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteServiceAccountTokenServiceAccountsResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type DeleteServiceAccountTokenServiceAccountsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteServiceAccountTokenServiceAccountsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteServiceAccountTokenServiceAccountsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteServiceAccountTokenServiceAccountsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteServiceAccountTokenResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type DeleteServiceAccountTokenResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteServiceAccountTokenResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteServiceAccountTokenResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteServiceAccountTokenResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type DeleteServiceAccountTokenResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *DeleteServiceAccountTokenResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *DeleteServiceAccountTokenServiceAccountsResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *DeleteServiceAccountTokenServiceAccountsResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteServiceAccountTokenResponse) GetFourHundredAndOneApplicationJSONObject() *DeleteServiceAccountTokenResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *DeleteServiceAccountTokenResponse) GetFourHundredAndThreeApplicationJSONObject() *DeleteServiceAccountTokenServiceAccountsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *DeleteServiceAccountTokenResponse) GetFourHundredAndFourApplicationJSONObject() *DeleteServiceAccountTokenServiceAccountsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *DeleteServiceAccountTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteServiceAccountTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteServiceAccountTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
