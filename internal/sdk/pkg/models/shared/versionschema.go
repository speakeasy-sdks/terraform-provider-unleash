// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// VersionSchemaCurrent - The current version of Unleash.
type VersionSchemaCurrent struct {
	// The Enterpris version of Unleash used to build this instance, represented as a git revision belonging to the [Unleash Enterprise](https://github.com/ivarconr/unleash-enterprise) repository. Will be an empty string if no enterprise version was used,
	Enterprise *string `json:"enterprise,omitempty"`
	// The OSS version used when building this Unleash instance, represented as a git revision belonging to the [main Unleash git repo](https://github.com/Unleash/unleash/)
	Oss *string `json:"oss,omitempty"`
}

// VersionSchemaLatest - Information about the latest available Unleash releases. Will be an empty object if no data is available.
type VersionSchemaLatest struct {
	// The latest available Enterprise version of Unleash
	Enterprise *string `json:"enterprise,omitempty"`
	// The latest available OSS version of Unleash
	Oss *string `json:"oss,omitempty"`
}

// VersionSchema - Detailed information about an Unleash version
type VersionSchema struct {
	// The current version of Unleash.
	Current VersionSchemaCurrent `json:"current"`
	// The instance identifier of the Unleash instance
	InstanceID *string `json:"instanceId,omitempty"`
	// Whether the Unleash server is running the latest release (`true`) or if there are updates available (`false`)
	IsLatest bool `json:"isLatest"`
	// Information about the latest available Unleash releases. Will be an empty object if no data is available.
	Latest VersionSchemaLatest `json:"latest"`
}
