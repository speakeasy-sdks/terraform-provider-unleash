// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unleash/internal/sdk/pkg/models/shared"
)

type AddFavoriteFeatureRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
}

type AddFavoriteFeatureResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Login401Response *shared.Login401Response
	// The requested resource was not found.
	SendResetPasswordEmail404Response *shared.SendResetPasswordEmail404Response
}
