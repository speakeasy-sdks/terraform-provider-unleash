// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unleash/internal/sdk/pkg/models/shared"
)

type RemoveStrategyRequest struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type RemoveStrategyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	ChangePassword403Response *shared.ChangePassword403Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Login401Response *shared.Login401Response
	// The requested resource was not found.
	SendResetPasswordEmail404Response *shared.SendResetPasswordEmail404Response
}
