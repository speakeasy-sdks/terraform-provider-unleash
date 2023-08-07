// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unleash/internal/sdk/pkg/models/shared"
)

type ValidateConstraintResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	ChangePassword403Response *shared.ChangePassword403Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Login401Response *shared.Login401Response
	// The request data does not match what we expect.
	ValidatePublicSignupToken400Response *shared.ValidatePublicSignupToken400Response
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	ValidateToken415Response *shared.ValidateToken415Response
}