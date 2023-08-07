// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"fmt"
	"net/http"
	"time"
	"unleash/internal/sdk/pkg/models/shared"
	"unleash/internal/sdk/pkg/utils"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"/",
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
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// UnleashServerAPI
type UnleashServerAPI struct {
	// APITokens - Create, update, and delete [Unleash API tokens](https://docs.getunleash.io/reference/api-tokens-and-client-keys).
	APITokens *apiTokens
	// Addons - Create, update, and delete [Unleash addons](https://docs.getunleash.io/addons).
	Addons *addons
	// AdminUI - Configuration for the Unleash Admin UI. These endpoints should not be relied upon and can change at any point without prior notice.
	AdminUI *adminUI
	// Archive - Revive or permanently delete [archived feature toggles](https://docs.getunleash.io/advanced/archived_toggles).
	Archive *archive
	// Auth - Manage logins, passwords, etc.
	Auth *auth
	// Client - Endpoints for [Unleash server-side clients](https://docs.getunleash.io/reference/sdks).
	Client *client
	// Edge - Endpoints related to Unleash on the Edge.
	Edge *edge
	// Environments - Create, update, delete, enable or disable [environments](https://docs.getunleash.io/reference/environments) for this Unleash instance.
	Environments *environments
	// Events - Read events from this Unleash instance.
	Events *events
	// Features - Create, update, and delete [features toggles](https://docs.getunleash.io/reference/feature-toggles).
	Features *features
	// FrontendAPI - API for connecting client-side (frontend) applications to Unleash.
	FrontendAPI *frontendAPI
	// ImportExport - [Import and export](https://docs.getunleash.io/deploy/import_export) the state of your Unleash instance.
	ImportExport *importExport
	// InstanceAdmin - Instance admin endpoints used to manage the Unleash instance itself.
	InstanceAdmin *instanceAdmin
	// Maintenance - Enable/disable the maintenance mode of Unleash.
	Maintenance *maintenance
	// Metrics - Register, read, or delete metrics recorded by Unleash.
	Metrics *metrics
	// Operational - Endpoints related to the operational status of this Unleash instance.
	Operational *operational
	// PersonalAccessTokens - Create, update, and delete [Personal access tokens](https://docs.getunleash.io/reference/api-tokens-and-client-keys#personal-access-tokens).
	PersonalAccessTokens *personalAccessTokens
	// Playground - Evaluate an Unleash context against your feature toggles.
	Playground *playground
	// Projects - Create, update, and delete [Unleash projects](https://docs.getunleash.io/reference/projects).
	Projects *projects
	// PublicSignupTokens - Create, update, and delete [Unleash Public Signup tokens](https://docs.getunleash.io/reference/public-signup-tokens).
	PublicSignupTokens *publicSignupTokens
	// Strategies - Create, update, delete, manage [custom strategies](https://docs.getunleash.io/reference/custom-activation-strategies).
	Strategies *strategies
	// Tags - Create, update, and delete [tags and tag types](https://docs.getunleash.io/reference/tags).
	Tags *tags
	// Telemetry - API for information about telemetry collection
	Telemetry *telemetry
	// Unstable - Experimental endpoints that may change or disappear at any time.
	Unstable *unstable
	// Users - Manage users and passwords.
	Users *users
	// UContext - Create, update, and delete [context fields](https://docs.getunleash.io/reference/unleash-context) that Unleash is aware of.
	UContext *uContext

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*UnleashServerAPI)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *UnleashServerAPI) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *UnleashServerAPI) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *UnleashServerAPI) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *UnleashServerAPI) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *UnleashServerAPI) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *UnleashServerAPI {
	sdk := &UnleashServerAPI{
		sdkConfiguration: sdkConfiguration{
			Language:          "terraform",
			OpenAPIDocVersion: "5.3.3",
			SDKVersion:        "0.0.1",
			GenVersion:        "2.81.1",
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

	sdk.APITokens = newAPITokens(sdk.sdkConfiguration)

	sdk.Addons = newAddons(sdk.sdkConfiguration)

	sdk.AdminUI = newAdminUI(sdk.sdkConfiguration)

	sdk.Archive = newArchive(sdk.sdkConfiguration)

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.Client = newClient(sdk.sdkConfiguration)

	sdk.Edge = newEdge(sdk.sdkConfiguration)

	sdk.Environments = newEnvironments(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.Features = newFeatures(sdk.sdkConfiguration)

	sdk.FrontendAPI = newFrontendAPI(sdk.sdkConfiguration)

	sdk.ImportExport = newImportExport(sdk.sdkConfiguration)

	sdk.InstanceAdmin = newInstanceAdmin(sdk.sdkConfiguration)

	sdk.Maintenance = newMaintenance(sdk.sdkConfiguration)

	sdk.Metrics = newMetrics(sdk.sdkConfiguration)

	sdk.Operational = newOperational(sdk.sdkConfiguration)

	sdk.PersonalAccessTokens = newPersonalAccessTokens(sdk.sdkConfiguration)

	sdk.Playground = newPlayground(sdk.sdkConfiguration)

	sdk.Projects = newProjects(sdk.sdkConfiguration)

	sdk.PublicSignupTokens = newPublicSignupTokens(sdk.sdkConfiguration)

	sdk.Strategies = newStrategies(sdk.sdkConfiguration)

	sdk.Tags = newTags(sdk.sdkConfiguration)

	sdk.Telemetry = newTelemetry(sdk.sdkConfiguration)

	sdk.Unstable = newUnstable(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.UContext = newUContext(sdk.sdkConfiguration)

	return sdk
}
