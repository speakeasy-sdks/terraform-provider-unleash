// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://app.unleash-hosted.com/hosted",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

type SDK struct {
	// Create, update, and delete [Unleash addons](https://docs.getunleash.io/addons).
	Addons *Addons
	// Create, update, and delete [Unleash API tokens](https://docs.getunleash.io/reference/api-tokens-and-client-keys).
	APITokens *APITokens
	// Revive or permanently delete [archived feature toggles](https://docs.getunleash.io/advanced/archived_toggles).
	Archive *Archive
	// Manage logins, passwords, etc.
	Auth *Auth
	// Register, read, or delete metrics recorded by Unleash.
	Metrics *Metrics
	// Create, update, and delete [features toggles](https://docs.getunleash.io/reference/feature-toggles).
	Features *Features
	// Create, update, and delete [context fields](https://docs.getunleash.io/reference/unleash-context) that Unleash is aware of.
	UContext *UContext
	// Create, update, delete, manage [custom strategies](https://docs.getunleash.io/reference/custom-activation-strategies).
	Strategies *Strategies
	// Create, update, delete, enable or disable [environments](https://docs.getunleash.io/reference/environments) for this Unleash instance.
	Environments *Environments
	// Read events from this Unleash instance.
	Events *Events
	// Experimental endpoints that may change or disappear at any time.
	Unstable *Unstable
	// [Import and export](https://docs.getunleash.io/deploy/import_export) the state of your Unleash instance.
	ImportExport *ImportExport
	// Configuration for the Unleash Admin UI. These endpoints should not be relied upon and can change at any point without prior notice.
	AdminUI *AdminUI
	// Manage users and passwords.
	Users *Users
	// Instance admin endpoints used to manage the Unleash instance itself.
	InstanceAdmin *InstanceAdmin
	// Create, update, and delete [Unleash Public Signup tokens](https://docs.getunleash.io/reference/public-signup-tokens).
	PublicSignupTokens *PublicSignupTokens
	// Enable/disable the maintenance mode of Unleash.
	Maintenance *Maintenance
	// Evaluate an Unleash context against your feature toggles.
	Playground *Playground
	// Create, update, and delete [Unleash projects](https://docs.getunleash.io/reference/projects).
	Projects *Projects
	// Create, update, and delete [tags and tag types](https://docs.getunleash.io/reference/tags).
	Tags *Tags
	// Create, update, delete, and manage [segments](https://docs.getunleash.io/reference/segments).
	Segments *Segments
	// Endpoints for managing [Service Accounts](https://docs.getunleash.io/reference/service-accounts), which enable programmatic access to the Unleash API.
	ServiceAccounts *ServiceAccounts
	// API for information about telemetry collection
	Telemetry *Telemetry
	// Create, update, and delete [Personal access tokens](https://docs.getunleash.io/reference/api-tokens-and-client-keys#personal-access-tokens).
	PersonalAccessTokens *PersonalAccessTokens
	// Endpoints for [Unleash server-side clients](https://docs.getunleash.io/reference/sdks).
	Client *Client
	// API for connecting client-side (frontend) applications to Unleash.
	FrontendAPI *FrontendAPI
	// Endpoints related to Unleash on the Edge.
	Edge *Edge
	// Endpoints related to the operational status of this Unleash instance.
	Operational *Operational

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "5.3.3",
			SDKVersion:        "0.15.0",
			GenVersion:        "2.228.1",
			UserAgent:         "speakeasy-sdk/go 0.15.0 2.228.1 5.3.3 terraform",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Addons = newAddons(sdk.sdkConfiguration)

	sdk.APITokens = newAPITokens(sdk.sdkConfiguration)

	sdk.Archive = newArchive(sdk.sdkConfiguration)

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.Metrics = newMetrics(sdk.sdkConfiguration)

	sdk.Features = newFeatures(sdk.sdkConfiguration)

	sdk.UContext = newUContext(sdk.sdkConfiguration)

	sdk.Strategies = newStrategies(sdk.sdkConfiguration)

	sdk.Environments = newEnvironments(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.Unstable = newUnstable(sdk.sdkConfiguration)

	sdk.ImportExport = newImportExport(sdk.sdkConfiguration)

	sdk.AdminUI = newAdminUI(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.InstanceAdmin = newInstanceAdmin(sdk.sdkConfiguration)

	sdk.PublicSignupTokens = newPublicSignupTokens(sdk.sdkConfiguration)

	sdk.Maintenance = newMaintenance(sdk.sdkConfiguration)

	sdk.Playground = newPlayground(sdk.sdkConfiguration)

	sdk.Projects = newProjects(sdk.sdkConfiguration)

	sdk.Tags = newTags(sdk.sdkConfiguration)

	sdk.Segments = newSegments(sdk.sdkConfiguration)

	sdk.ServiceAccounts = newServiceAccounts(sdk.sdkConfiguration)

	sdk.Telemetry = newTelemetry(sdk.sdkConfiguration)

	sdk.PersonalAccessTokens = newPersonalAccessTokens(sdk.sdkConfiguration)

	sdk.Client = newClient(sdk.sdkConfiguration)

	sdk.FrontendAPI = newFrontendAPI(sdk.sdkConfiguration)

	sdk.Edge = newEdge(sdk.sdkConfiguration)

	sdk.Operational = newOperational(sdk.sdkConfiguration)

	return sdk
}
