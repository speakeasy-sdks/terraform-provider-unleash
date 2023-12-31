// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ValidatePublicSignupTokenRequest struct {
	Token string `pathParam:"style=simple,explode=false,name=token"`
}

func (o *ValidatePublicSignupTokenRequest) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

// ValidatePublicSignupTokenResponseBody - The request data does not match what we expect.
type ValidatePublicSignupTokenResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidatePublicSignupTokenResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidatePublicSignupTokenResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidatePublicSignupTokenResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ValidatePublicSignupTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request data does not match what we expect.
	Object *ValidatePublicSignupTokenResponseBody
}

func (o *ValidatePublicSignupTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ValidatePublicSignupTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ValidatePublicSignupTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ValidatePublicSignupTokenResponse) GetObject() *ValidatePublicSignupTokenResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
