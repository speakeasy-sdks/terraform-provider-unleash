// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ChangeRoleForGroupRequest struct {
	GroupID   string `pathParam:"style=simple,explode=false,name=groupId"`
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
	RoleID    string `pathParam:"style=simple,explode=false,name=roleId"`
}

func (o *ChangeRoleForGroupRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

func (o *ChangeRoleForGroupRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *ChangeRoleForGroupRequest) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

type ChangeRoleForGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ChangeRoleForGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ChangeRoleForGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ChangeRoleForGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
