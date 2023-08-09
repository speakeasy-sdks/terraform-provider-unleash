// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unleash/internal/sdk/pkg/models/shared"
)

type ArchiveFeatureRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
}

type ArchiveFeatureResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetGoogleSettings403Response *shared.GetGoogleSettings403Response
	// The requested resource was not found.
	GetGroup404Response *shared.GetGroup404Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Login401Response *shared.Login401Response
}
