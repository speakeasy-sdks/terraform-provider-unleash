// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"time"
)

// HealthReportSchemaMode - The project's [collaboration mode](https://docs.getunleash.io/reference/project-collaboration-mode). Determines whether non-project members can submit change requests or not.
type HealthReportSchemaMode string

const (
	HealthReportSchemaModeOpen      HealthReportSchemaMode = "open"
	HealthReportSchemaModeProtected HealthReportSchemaMode = "protected"
)

func (e HealthReportSchemaMode) ToPointer() *HealthReportSchemaMode {
	return &e
}

func (e *HealthReportSchemaMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "open":
		fallthrough
	case "protected":
		*e = HealthReportSchemaMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HealthReportSchemaMode: %v", v)
	}
}

// HealthReportSchema - A report of the current health of the requested project, with datapoints like counters of currently active, stale, and potentially stale feature toggles.
type HealthReportSchema struct {
	// The number of active feature toggles.
	ActiveCount float64 `json:"activeCount"`
	// When the project was last updated.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy
	DefaultStickiness string `json:"defaultStickiness"`
	// The project's description
	Description *string `json:"description,omitempty"`
	// An array containing the names of all the environments configured for the project.
	Environments []ProjectEnvironmentSchema `json:"environments"`
	// Indicates if the project has been marked as a favorite by the current user requesting the project health overview.
	Favorite *bool `json:"favorite,omitempty"`
	// A limit on the number of features allowed in the project. Null if no limit.
	FeatureLimit *float64 `json:"featureLimit,omitempty"`
	// An array containing an overview of all the features of the project and their individual status
	Features []FeatureSchema `json:"features"`
	// The overall [health rating](https://docs.getunleash.io/reference/technical-debt#health-rating) of the project.
	Health int64 `json:"health"`
	// The number of users/members in the project.
	Members int64 `json:"members"`
	// The project's [collaboration mode](https://docs.getunleash.io/reference/project-collaboration-mode). Determines whether non-project members can submit change requests or not.
	Mode HealthReportSchemaMode `json:"mode"`
	// The project's name
	Name string `json:"name"`
	// The number of potentially stale feature toggles.
	PotentiallyStaleCount float64 `json:"potentiallyStaleCount"`
	// The number of stale feature toggles.
	StaleCount float64 `json:"staleCount"`
	// Statistics for a project, including the average time to production, number of features created, the project activity and more.
	//
	// Stats are divided into current and previous **windows**.
	// - The **current window** is the past 30 days.
	// - The **previous window** is the 30 days **before** the current window (from 60 to 30 days ago)
	Stats *ProjectStatsSchema `json:"stats,omitempty"`
	// When the project was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The project overview version.
	Version int64 `json:"version"`
}

func (h HealthReportSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HealthReportSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HealthReportSchema) GetActiveCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.ActiveCount
}

func (o *HealthReportSchema) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *HealthReportSchema) GetDefaultStickiness() string {
	if o == nil {
		return ""
	}
	return o.DefaultStickiness
}

func (o *HealthReportSchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *HealthReportSchema) GetEnvironments() []ProjectEnvironmentSchema {
	if o == nil {
		return []ProjectEnvironmentSchema{}
	}
	return o.Environments
}

func (o *HealthReportSchema) GetFavorite() *bool {
	if o == nil {
		return nil
	}
	return o.Favorite
}

func (o *HealthReportSchema) GetFeatureLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.FeatureLimit
}

func (o *HealthReportSchema) GetFeatures() []FeatureSchema {
	if o == nil {
		return []FeatureSchema{}
	}
	return o.Features
}

func (o *HealthReportSchema) GetHealth() int64 {
	if o == nil {
		return 0
	}
	return o.Health
}

func (o *HealthReportSchema) GetMembers() int64 {
	if o == nil {
		return 0
	}
	return o.Members
}

func (o *HealthReportSchema) GetMode() HealthReportSchemaMode {
	if o == nil {
		return HealthReportSchemaMode("")
	}
	return o.Mode
}

func (o *HealthReportSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *HealthReportSchema) GetPotentiallyStaleCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.PotentiallyStaleCount
}

func (o *HealthReportSchema) GetStaleCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.StaleCount
}

func (o *HealthReportSchema) GetStats() *ProjectStatsSchema {
	if o == nil {
		return nil
	}
	return o.Stats
}

func (o *HealthReportSchema) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *HealthReportSchema) GetVersion() int64 {
	if o == nil {
		return 0
	}
	return o.Version
}
