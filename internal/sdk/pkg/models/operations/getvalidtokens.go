// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetValidTokensResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The request body is larger than what we accept. By default we only accept bodies of 100kB or less
	CreateAddon413Response *shared.CreateAddon413Response
	// The request data does not match what we expect.
	GetGoogleSettings400Response *shared.GetGoogleSettings400Response
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	SetGoogleSettings415Response *shared.SetGoogleSettings415Response
	// validatedEdgeTokensSchema
	ValidatedEdgeTokensSchema *shared.ValidatedEdgeTokensSchema
}
