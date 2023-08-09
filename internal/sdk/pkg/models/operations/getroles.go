// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unleash/internal/sdk/pkg/models/shared"
)

type GetRolesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// rolesWithVersionSchema
	RolesWithVersionSchema *shared.RolesWithVersionSchema
}