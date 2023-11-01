// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type UpdateRoleRequest struct {
	// createRoleWithPermissionsSchema
	CreateRoleWithPermissionsSchema shared.CreateRoleWithPermissionsSchema `request:"mediaType=application/json"`
	RoleID                          string                                 `pathParam:"style=simple,explode=false,name=roleId"`
}

func (o *UpdateRoleRequest) GetCreateRoleWithPermissionsSchema() shared.CreateRoleWithPermissionsSchema {
	if o == nil {
		return shared.CreateRoleWithPermissionsSchema{}
	}
	return o.CreateRoleWithPermissionsSchema
}

func (o *UpdateRoleRequest) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

// UpdateRole400ApplicationJSON - The request data does not match what we expect.
type UpdateRole400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateRole400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateRole400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateRole400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type UpdateRoleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// roleWithVersionSchema
	RoleWithVersionSchema *shared.RoleWithVersionSchema
	// The request data does not match what we expect.
	UpdateRole400ApplicationJSONObject *UpdateRole400ApplicationJSON
}

func (o *UpdateRoleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRoleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRoleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRoleResponse) GetRoleWithVersionSchema() *shared.RoleWithVersionSchema {
	if o == nil {
		return nil
	}
	return o.RoleWithVersionSchema
}

func (o *UpdateRoleResponse) GetUpdateRole400ApplicationJSONObject() *UpdateRole400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateRole400ApplicationJSONObject
}
