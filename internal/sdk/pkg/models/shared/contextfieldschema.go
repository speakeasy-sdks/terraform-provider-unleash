// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"time"
)

// ContextFieldSchema - A representation of a [context field](https://docs.getunleash.io/reference/unleash-context).
type ContextFieldSchema struct {
	// When this context field was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The description of the context field.
	Description *string `json:"description,omitempty"`
	// Allowed values for this context field schema. Can be used to narrow down accepted input
	LegalValues []LegalValueSchema `json:"legalValues,omitempty"`
	// The name of the context field
	Name string `json:"name"`
	// Used when sorting a list of context fields. Is also used as a tiebreaker if a list of context fields is sorted alphabetically.
	SortOrder *int64 `json:"sortOrder,omitempty"`
	// Does this context field support being used for [stickiness](https://docs.getunleash.io/reference/stickiness) calculations
	Stickiness *bool `json:"stickiness,omitempty"`
	// Number of projects where this context field is used in
	UsedInFeatures *int64 `json:"usedInFeatures,omitempty"`
	// Number of projects where this context field is used in
	UsedInProjects *int64 `json:"usedInProjects,omitempty"`
}

func (c ContextFieldSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ContextFieldSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ContextFieldSchema) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ContextFieldSchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ContextFieldSchema) GetLegalValues() []LegalValueSchema {
	if o == nil {
		return nil
	}
	return o.LegalValues
}

func (o *ContextFieldSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ContextFieldSchema) GetSortOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.SortOrder
}

func (o *ContextFieldSchema) GetStickiness() *bool {
	if o == nil {
		return nil
	}
	return o.Stickiness
}

func (o *ContextFieldSchema) GetUsedInFeatures() *int64 {
	if o == nil {
		return nil
	}
	return o.UsedInFeatures
}

func (o *ContextFieldSchema) GetUsedInProjects() *int64 {
	if o == nil {
		return nil
	}
	return o.UsedInProjects
}
