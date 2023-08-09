// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetAddonRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetAddonResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// addonSchema
	AddonSchema *shared.AddonSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Login401Response *shared.Login401Response
}
