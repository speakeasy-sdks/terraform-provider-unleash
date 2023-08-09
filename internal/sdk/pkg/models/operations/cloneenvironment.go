// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type CloneEnvironmentRequest struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type CloneEnvironmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The request data does not match what we expect.
	GetGoogleSettings400Response *shared.GetGoogleSettings400Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Login401Response *shared.Login401Response
}
