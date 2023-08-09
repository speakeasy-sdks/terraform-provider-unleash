// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unleash/internal/sdk/pkg/models/shared"
)

type ValidatePublicSignupTokenRequest struct {
	Token string `pathParam:"style=simple,explode=false,name=token"`
}

type ValidatePublicSignupTokenResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The request data does not match what we expect.
	GetGoogleSettings400Response *shared.GetGoogleSettings400Response
}
