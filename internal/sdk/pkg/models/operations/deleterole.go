// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteRoleRequest struct {
	RoleID string `pathParam:"style=simple,explode=false,name=roleId"`
}

// DeleteRole409ApplicationJSON - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type DeleteRole409ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type DeleteRoleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	DeleteRole409ApplicationJSONObject *DeleteRole409ApplicationJSON
}
