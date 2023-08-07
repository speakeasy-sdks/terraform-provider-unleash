// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteTagRequest struct {
	Type  string `pathParam:"style=simple,explode=false,name=type"`
	Value string `pathParam:"style=simple,explode=false,name=value"`
}

type DeleteTagResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}