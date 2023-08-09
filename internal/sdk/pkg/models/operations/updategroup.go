// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type UpdateGroupRequest struct {
	// createGroupSchema
	CreateGroupSchema shared.CreateGroupSchema `request:"mediaType=application/json"`
	GroupID           string                   `pathParam:"style=simple,explode=false,name=groupId"`
}

type UpdateGroupResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	CreateGroup409Response *shared.CreateGroup409Response
	// The request data does not match what we expect.
	GetGoogleSettings400Response *shared.GetGoogleSettings400Response
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetGoogleSettings403Response *shared.GetGoogleSettings403Response
	// The requested resource was not found.
	GetGroup404Response *shared.GetGroup404Response
	// groupSchema
	GroupSchema *shared.GroupSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Login401Response *shared.Login401Response
}
