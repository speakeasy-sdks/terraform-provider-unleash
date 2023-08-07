// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AddPublicSignupTokenUser409Response - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type AddPublicSignupTokenUser409Response struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}