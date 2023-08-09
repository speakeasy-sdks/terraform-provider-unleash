// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OidcSettingsSchemaDefaultRootRole - [Default role](https://docs.getunleash.io/reference/rbac#standard-roles) granted to users auto-created from email. Only relevant if autoCreate is `true`
type OidcSettingsSchemaDefaultRootRole string

const (
	OidcSettingsSchemaDefaultRootRoleViewer OidcSettingsSchemaDefaultRootRole = "Viewer"
	OidcSettingsSchemaDefaultRootRoleEditor OidcSettingsSchemaDefaultRootRole = "Editor"
	OidcSettingsSchemaDefaultRootRoleAdmin  OidcSettingsSchemaDefaultRootRole = "Admin"
)

func (e OidcSettingsSchemaDefaultRootRole) ToPointer() *OidcSettingsSchemaDefaultRootRole {
	return &e
}

func (e *OidcSettingsSchemaDefaultRootRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Viewer":
		fallthrough
	case "Editor":
		fallthrough
	case "Admin":
		*e = OidcSettingsSchemaDefaultRootRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OidcSettingsSchemaDefaultRootRole: %v", v)
	}
}

// OidcSettingsSchemaIDTokenSigningAlgorithm - The signing algorithm used to sign our token. Refer to the [JWT signatures](https://jwt.io/introduction) documentation for more information.
type OidcSettingsSchemaIDTokenSigningAlgorithm string

const (
	OidcSettingsSchemaIDTokenSigningAlgorithmRs256 OidcSettingsSchemaIDTokenSigningAlgorithm = "RS256"
	OidcSettingsSchemaIDTokenSigningAlgorithmRs384 OidcSettingsSchemaIDTokenSigningAlgorithm = "RS384"
	OidcSettingsSchemaIDTokenSigningAlgorithmRs512 OidcSettingsSchemaIDTokenSigningAlgorithm = "RS512"
)

func (e OidcSettingsSchemaIDTokenSigningAlgorithm) ToPointer() *OidcSettingsSchemaIDTokenSigningAlgorithm {
	return &e
}

func (e *OidcSettingsSchemaIDTokenSigningAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RS256":
		fallthrough
	case "RS384":
		fallthrough
	case "RS512":
		*e = OidcSettingsSchemaIDTokenSigningAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OidcSettingsSchemaIDTokenSigningAlgorithm: %v", v)
	}
}

// OidcSettingsSchema - Settings for configuring OpenID Connect as a login provider for Unleash
type OidcSettingsSchema struct {
	// Authentication Context Class Reference, used to request extra values in the acr claim returned from the server. If multiple values are required, they should be space separated.
	//  Consult [the OIDC reference](https://openid.net/specs/openid-connect-core-1_0.html#AuthorizationEndpoint) for more information
	//
	AcrValues *string `json:"acrValues,omitempty"`
	// Auto create users based on email addresses from login tokens
	AutoCreate *bool `json:"autoCreate,omitempty"`
	// The OIDC client ID of this application.
	ClientID string `json:"clientId"`
	// [Default role](https://docs.getunleash.io/reference/rbac#standard-roles) granted to users auto-created from email. Only relevant if autoCreate is `true`
	DefaultRootRole *OidcSettingsSchemaDefaultRootRole `json:"defaultRootRole,omitempty"`
	// The [.well-known OpenID discover URL](https://swagger.io/docs/specification/authentication/openid-connect-discovery/)
	DiscoverURL *string `json:"discoverUrl,omitempty"`
	// Comma separated list of email domains that are automatically approved for an account in the server. Only relevant if autoCreate is `true`
	EmailDomains *string `json:"emailDomains,omitempty"`
	// Support Single sign out when user clicks logout in Unleash. If `true` user is signed out of all OpenID Connect sessions against the clientId they may have active
	EnableSingleSignOut *bool `json:"enableSingleSignOut,omitempty"`
	// `true` if OpenID connect is turned on for this instance, otherwise `false`
	Enabled *bool `json:"enabled,omitempty"`
	// The signing algorithm used to sign our token. Refer to the [JWT signatures](https://jwt.io/introduction) documentation for more information.
	IDTokenSigningAlgorithm *OidcSettingsSchemaIDTokenSigningAlgorithm `json:"idTokenSigningAlgorithm,omitempty"`
	// Shared secret from OpenID server. Used to authenticate login requests
	Secret string `json:"secret"`
}
