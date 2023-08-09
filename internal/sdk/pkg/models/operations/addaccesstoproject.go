// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AddAccessToProjectRequest struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
	RoleID    string `pathParam:"style=simple,explode=false,name=roleId"`
}

type AddAccessToProjectResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
