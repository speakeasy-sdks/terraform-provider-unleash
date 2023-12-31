// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"time"
)

// Range - A description of a time range
type Range string

const (
	RangeAllTime Range = "allTime"
	RangeThirtyd Range = "30d"
	RangeSevend  Range = "7d"
)

func (e Range) ToPointer() *Range {
	return &e
}

func (e *Range) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allTime":
		fallthrough
	case "30d":
		fallthrough
	case "7d":
		*e = Range(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Range: %v", v)
	}
}

// ClientApps - An entry describing how many client applications has been observed over the defined range
type ClientApps struct {
	// The number of client applications that have been observed in this period
	Count *float64 `json:"count,omitempty"`
	// A description of a time range
	Range *Range `json:"range,omitempty"`
}

func (o *ClientApps) GetCount() *float64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ClientApps) GetRange() *Range {
	if o == nil {
		return nil
	}
	return o.Range
}

// InstanceAdminStatsSchema - Information about an instance and statistics about usage of various features of Unleash
type InstanceAdminStatsSchema struct {
	// Whether or not OIDC authentication is enabled for this instance
	OIDCenabled *bool `json:"OIDCenabled,omitempty"`
	// Whether or not SAML authentication is enabled for this instance
	SAMLenabled *bool `json:"SAMLenabled,omitempty"`
	// A count of connected applications in the last week, last month and all time since last restart
	ClientApps []ClientApps `json:"clientApps,omitempty"`
	// The number of context fields defined in this instance.
	ContextFields *float64 `json:"contextFields,omitempty"`
	// The number of environments defined in this instance
	Environments *float64 `json:"environments,omitempty"`
	// The number of feature-toggles this instance has
	FeatureToggles *float64 `json:"featureToggles,omitempty"`
	// The number of groups defined in this instance
	Groups *float64 `json:"groups,omitempty"`
	// A unique identifier for this instance. Generated by the database migration scripts at first run. Typically a UUID.
	InstanceID string `json:"instanceId"`
	// The number of projects defined in this instance.
	Projects *float64 `json:"projects,omitempty"`
	// The number of roles defined in this instance
	Roles *float64 `json:"roles,omitempty"`
	// The number of segments defined in this instance
	Segments *float64 `json:"segments,omitempty"`
	// The number of strategies defined in this instance
	Strategies *float64 `json:"strategies,omitempty"`
	// A SHA-256 checksum of the instance statistics to be used to verify that the data in this object has not been tampered with
	Sum *string `json:"sum,omitempty"`
	// When these statistics were produced
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The number of users this instance has
	Users *float64 `json:"users,omitempty"`
	// The version of Unleash Enterprise that is bundled in this instance
	VersionEnterprise *string `json:"versionEnterprise,omitempty"`
	// The version of Unleash OSS that is bundled in this instance
	VersionOSS *string `json:"versionOSS,omitempty"`
}

func (i InstanceAdminStatsSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InstanceAdminStatsSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InstanceAdminStatsSchema) GetOIDCenabled() *bool {
	if o == nil {
		return nil
	}
	return o.OIDCenabled
}

func (o *InstanceAdminStatsSchema) GetSAMLenabled() *bool {
	if o == nil {
		return nil
	}
	return o.SAMLenabled
}

func (o *InstanceAdminStatsSchema) GetClientApps() []ClientApps {
	if o == nil {
		return nil
	}
	return o.ClientApps
}

func (o *InstanceAdminStatsSchema) GetContextFields() *float64 {
	if o == nil {
		return nil
	}
	return o.ContextFields
}

func (o *InstanceAdminStatsSchema) GetEnvironments() *float64 {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *InstanceAdminStatsSchema) GetFeatureToggles() *float64 {
	if o == nil {
		return nil
	}
	return o.FeatureToggles
}

func (o *InstanceAdminStatsSchema) GetGroups() *float64 {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *InstanceAdminStatsSchema) GetInstanceID() string {
	if o == nil {
		return ""
	}
	return o.InstanceID
}

func (o *InstanceAdminStatsSchema) GetProjects() *float64 {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *InstanceAdminStatsSchema) GetRoles() *float64 {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *InstanceAdminStatsSchema) GetSegments() *float64 {
	if o == nil {
		return nil
	}
	return o.Segments
}

func (o *InstanceAdminStatsSchema) GetStrategies() *float64 {
	if o == nil {
		return nil
	}
	return o.Strategies
}

func (o *InstanceAdminStatsSchema) GetSum() *string {
	if o == nil {
		return nil
	}
	return o.Sum
}

func (o *InstanceAdminStatsSchema) GetTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.Timestamp
}

func (o *InstanceAdminStatsSchema) GetUsers() *float64 {
	if o == nil {
		return nil
	}
	return o.Users
}

func (o *InstanceAdminStatsSchema) GetVersionEnterprise() *string {
	if o == nil {
		return nil
	}
	return o.VersionEnterprise
}

func (o *InstanceAdminStatsSchema) GetVersionOSS() *string {
	if o == nil {
		return nil
	}
	return o.VersionOSS
}
