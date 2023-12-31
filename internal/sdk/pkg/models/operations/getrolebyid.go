// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetRoleByIDRequest struct {
	RoleID string `pathParam:"style=simple,explode=false,name=roleId"`
}

func (o *GetRoleByIDRequest) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

type GetRoleByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// roleWithPermissionsSchema
	RoleWithPermissionsSchema *shared.RoleWithPermissionsSchema
}

func (o *GetRoleByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRoleByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRoleByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetRoleByIDResponse) GetRoleWithPermissionsSchema() *shared.RoleWithPermissionsSchema {
	if o == nil {
		return nil
	}
	return o.RoleWithPermissionsSchema
}
