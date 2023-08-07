// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unleash/internal/sdk/pkg/models/shared"
)

type GetApplicationRequest struct {
	AppName string `pathParam:"style=simple,explode=false,name=appName"`
}

type GetApplicationResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// applicationSchema
	ApplicationSchema *shared.ApplicationSchema
	// The requested resource was not found.
	SendResetPasswordEmail404Response *shared.SendResetPasswordEmail404Response
}
