// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type OverwriteEnvironmentFeatureVariantsRequest struct {
	// variantsSchema
	RequestBody []shared.VariantSchema `request:"mediaType=application/json"`
	Environment string                 `pathParam:"style=simple,explode=false,name=environment"`
	FeatureName string                 `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string                 `pathParam:"style=simple,explode=false,name=projectId"`
}

type OverwriteEnvironmentFeatureVariantsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// featureVariantsSchema
	FeatureVariantsSchema *shared.FeatureVariantsSchema
	// The request data does not match what we expect.
	GetGoogleSettings400Response *shared.GetGoogleSettings400Response
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetGoogleSettings403Response *shared.GetGoogleSettings403Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Login401Response *shared.Login401Response
}
