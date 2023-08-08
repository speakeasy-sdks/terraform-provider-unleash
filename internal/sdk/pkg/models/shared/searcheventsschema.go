// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SearchEventsSchemaType - Find events by event type (case-sensitive).
type SearchEventsSchemaType string

const (
	SearchEventsSchemaTypeApplicationCreated                SearchEventsSchemaType = "application-created"
	SearchEventsSchemaTypeFeatureCreated                    SearchEventsSchemaType = "feature-created"
	SearchEventsSchemaTypeFeatureDeleted                    SearchEventsSchemaType = "feature-deleted"
	SearchEventsSchemaTypeFeatureUpdated                    SearchEventsSchemaType = "feature-updated"
	SearchEventsSchemaTypeFeatureMetadataUpdated            SearchEventsSchemaType = "feature-metadata-updated"
	SearchEventsSchemaTypeFeatureVariantsUpdated            SearchEventsSchemaType = "feature-variants-updated"
	SearchEventsSchemaTypeFeatureEnvironmentVariantsUpdated SearchEventsSchemaType = "feature-environment-variants-updated"
	SearchEventsSchemaTypeFeatureProjectChange              SearchEventsSchemaType = "feature-project-change"
	SearchEventsSchemaTypeFeatureArchived                   SearchEventsSchemaType = "feature-archived"
	SearchEventsSchemaTypeFeatureRevived                    SearchEventsSchemaType = "feature-revived"
	SearchEventsSchemaTypeFeatureImport                     SearchEventsSchemaType = "feature-import"
	SearchEventsSchemaTypeFeatureTagged                     SearchEventsSchemaType = "feature-tagged"
	SearchEventsSchemaTypeFeatureTagImport                  SearchEventsSchemaType = "feature-tag-import"
	SearchEventsSchemaTypeFeatureStrategyUpdate             SearchEventsSchemaType = "feature-strategy-update"
	SearchEventsSchemaTypeFeatureStrategyAdd                SearchEventsSchemaType = "feature-strategy-add"
	SearchEventsSchemaTypeFeatureStrategyRemove             SearchEventsSchemaType = "feature-strategy-remove"
	SearchEventsSchemaTypeStrategyOrderChanged              SearchEventsSchemaType = "strategy-order-changed"
	SearchEventsSchemaTypeDropFeatureTags                   SearchEventsSchemaType = "drop-feature-tags"
	SearchEventsSchemaTypeFeatureUntagged                   SearchEventsSchemaType = "feature-untagged"
	SearchEventsSchemaTypeFeatureStaleOn                    SearchEventsSchemaType = "feature-stale-on"
	SearchEventsSchemaTypeFeatureStaleOff                   SearchEventsSchemaType = "feature-stale-off"
	SearchEventsSchemaTypeDropFeatures                      SearchEventsSchemaType = "drop-features"
	SearchEventsSchemaTypeFeatureEnvironmentEnabled         SearchEventsSchemaType = "feature-environment-enabled"
	SearchEventsSchemaTypeFeatureEnvironmentDisabled        SearchEventsSchemaType = "feature-environment-disabled"
	SearchEventsSchemaTypeStrategyCreated                   SearchEventsSchemaType = "strategy-created"
	SearchEventsSchemaTypeStrategyDeleted                   SearchEventsSchemaType = "strategy-deleted"
	SearchEventsSchemaTypeStrategyDeprecated                SearchEventsSchemaType = "strategy-deprecated"
	SearchEventsSchemaTypeStrategyReactivated               SearchEventsSchemaType = "strategy-reactivated"
	SearchEventsSchemaTypeStrategyUpdated                   SearchEventsSchemaType = "strategy-updated"
	SearchEventsSchemaTypeStrategyImport                    SearchEventsSchemaType = "strategy-import"
	SearchEventsSchemaTypeDropStrategies                    SearchEventsSchemaType = "drop-strategies"
	SearchEventsSchemaTypeContextFieldCreated               SearchEventsSchemaType = "context-field-created"
	SearchEventsSchemaTypeContextFieldUpdated               SearchEventsSchemaType = "context-field-updated"
	SearchEventsSchemaTypeContextFieldDeleted               SearchEventsSchemaType = "context-field-deleted"
	SearchEventsSchemaTypeProjectAccessAdded                SearchEventsSchemaType = "project-access-added"
	SearchEventsSchemaTypeProjectCreated                    SearchEventsSchemaType = "project-created"
	SearchEventsSchemaTypeProjectUpdated                    SearchEventsSchemaType = "project-updated"
	SearchEventsSchemaTypeProjectDeleted                    SearchEventsSchemaType = "project-deleted"
	SearchEventsSchemaTypeProjectImport                     SearchEventsSchemaType = "project-import"
	SearchEventsSchemaTypeProjectUserAdded                  SearchEventsSchemaType = "project-user-added"
	SearchEventsSchemaTypeProjectUserRemoved                SearchEventsSchemaType = "project-user-removed"
	SearchEventsSchemaTypeProjectUserRoleChanged            SearchEventsSchemaType = "project-user-role-changed"
	SearchEventsSchemaTypeProjectGroupRoleChanged           SearchEventsSchemaType = "project-group-role-changed"
	SearchEventsSchemaTypeProjectGroupAdded                 SearchEventsSchemaType = "project-group-added"
	SearchEventsSchemaTypeProjectGroupRemoved               SearchEventsSchemaType = "project-group-removed"
	SearchEventsSchemaTypeDropProjects                      SearchEventsSchemaType = "drop-projects"
	SearchEventsSchemaTypeTagCreated                        SearchEventsSchemaType = "tag-created"
	SearchEventsSchemaTypeTagDeleted                        SearchEventsSchemaType = "tag-deleted"
	SearchEventsSchemaTypeTagImport                         SearchEventsSchemaType = "tag-import"
	SearchEventsSchemaTypeDropTags                          SearchEventsSchemaType = "drop-tags"
	SearchEventsSchemaTypeTagTypeCreated                    SearchEventsSchemaType = "tag-type-created"
	SearchEventsSchemaTypeTagTypeDeleted                    SearchEventsSchemaType = "tag-type-deleted"
	SearchEventsSchemaTypeTagTypeUpdated                    SearchEventsSchemaType = "tag-type-updated"
	SearchEventsSchemaTypeTagTypeImport                     SearchEventsSchemaType = "tag-type-import"
	SearchEventsSchemaTypeDropTagTypes                      SearchEventsSchemaType = "drop-tag-types"
	SearchEventsSchemaTypeAddonConfigCreated                SearchEventsSchemaType = "addon-config-created"
	SearchEventsSchemaTypeAddonConfigUpdated                SearchEventsSchemaType = "addon-config-updated"
	SearchEventsSchemaTypeAddonConfigDeleted                SearchEventsSchemaType = "addon-config-deleted"
	SearchEventsSchemaTypeDbPoolUpdate                      SearchEventsSchemaType = "db-pool-update"
	SearchEventsSchemaTypeUserCreated                       SearchEventsSchemaType = "user-created"
	SearchEventsSchemaTypeUserUpdated                       SearchEventsSchemaType = "user-updated"
	SearchEventsSchemaTypeUserDeleted                       SearchEventsSchemaType = "user-deleted"
	SearchEventsSchemaTypeDropEnvironments                  SearchEventsSchemaType = "drop-environments"
	SearchEventsSchemaTypeEnvironmentImport                 SearchEventsSchemaType = "environment-import"
	SearchEventsSchemaTypeSegmentCreated                    SearchEventsSchemaType = "segment-created"
	SearchEventsSchemaTypeSegmentUpdated                    SearchEventsSchemaType = "segment-updated"
	SearchEventsSchemaTypeSegmentDeleted                    SearchEventsSchemaType = "segment-deleted"
	SearchEventsSchemaTypeGroupCreated                      SearchEventsSchemaType = "group-created"
	SearchEventsSchemaTypeGroupUpdated                      SearchEventsSchemaType = "group-updated"
	SearchEventsSchemaTypeSettingCreated                    SearchEventsSchemaType = "setting-created"
	SearchEventsSchemaTypeSettingUpdated                    SearchEventsSchemaType = "setting-updated"
	SearchEventsSchemaTypeSettingDeleted                    SearchEventsSchemaType = "setting-deleted"
	SearchEventsSchemaTypeClientMetrics                     SearchEventsSchemaType = "client-metrics"
	SearchEventsSchemaTypeClientRegister                    SearchEventsSchemaType = "client-register"
	SearchEventsSchemaTypePatCreated                        SearchEventsSchemaType = "pat-created"
	SearchEventsSchemaTypePatDeleted                        SearchEventsSchemaType = "pat-deleted"
	SearchEventsSchemaTypePublicSignupTokenCreated          SearchEventsSchemaType = "public-signup-token-created"
	SearchEventsSchemaTypePublicSignupTokenUserAdded        SearchEventsSchemaType = "public-signup-token-user-added"
	SearchEventsSchemaTypePublicSignupTokenUpdated          SearchEventsSchemaType = "public-signup-token-updated"
	SearchEventsSchemaTypeChangeRequestCreated              SearchEventsSchemaType = "change-request-created"
	SearchEventsSchemaTypeChangeRequestDiscarded            SearchEventsSchemaType = "change-request-discarded"
	SearchEventsSchemaTypeChangeAdded                       SearchEventsSchemaType = "change-added"
	SearchEventsSchemaTypeChangeDiscarded                   SearchEventsSchemaType = "change-discarded"
	SearchEventsSchemaTypeChangeEdited                      SearchEventsSchemaType = "change-edited"
	SearchEventsSchemaTypeChangeRequestApproved             SearchEventsSchemaType = "change-request-approved"
	SearchEventsSchemaTypeChangeRequestApprovalAdded        SearchEventsSchemaType = "change-request-approval-added"
	SearchEventsSchemaTypeChangeRequestCancelled            SearchEventsSchemaType = "change-request-cancelled"
	SearchEventsSchemaTypeChangeRequestSentToReview         SearchEventsSchemaType = "change-request-sent-to-review"
	SearchEventsSchemaTypeChangeRequestApplied              SearchEventsSchemaType = "change-request-applied"
	SearchEventsSchemaTypeAPITokenCreated                   SearchEventsSchemaType = "api-token-created"
	SearchEventsSchemaTypeAPITokenUpdated                   SearchEventsSchemaType = "api-token-updated"
	SearchEventsSchemaTypeAPITokenDeleted                   SearchEventsSchemaType = "api-token-deleted"
	SearchEventsSchemaTypeFeatureFavorited                  SearchEventsSchemaType = "feature-favorited"
	SearchEventsSchemaTypeFeatureUnfavorited                SearchEventsSchemaType = "feature-unfavorited"
	SearchEventsSchemaTypeProjectFavorited                  SearchEventsSchemaType = "project-favorited"
	SearchEventsSchemaTypeProjectUnfavorited                SearchEventsSchemaType = "project-unfavorited"
	SearchEventsSchemaTypeFeaturesExported                  SearchEventsSchemaType = "features-exported"
	SearchEventsSchemaTypeFeaturesImported                  SearchEventsSchemaType = "features-imported"
	SearchEventsSchemaTypeServiceAccountCreated             SearchEventsSchemaType = "service-account-created"
	SearchEventsSchemaTypeServiceAccountDeleted             SearchEventsSchemaType = "service-account-deleted"
	SearchEventsSchemaTypeServiceAccountUpdated             SearchEventsSchemaType = "service-account-updated"
	SearchEventsSchemaTypeFeaturePotentiallyStaleOn         SearchEventsSchemaType = "feature-potentially-stale-on"
)

func (e SearchEventsSchemaType) ToPointer() *SearchEventsSchemaType {
	return &e
}

func (e *SearchEventsSchemaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "application-created":
		fallthrough
	case "feature-created":
		fallthrough
	case "feature-deleted":
		fallthrough
	case "feature-updated":
		fallthrough
	case "feature-metadata-updated":
		fallthrough
	case "feature-variants-updated":
		fallthrough
	case "feature-environment-variants-updated":
		fallthrough
	case "feature-project-change":
		fallthrough
	case "feature-archived":
		fallthrough
	case "feature-revived":
		fallthrough
	case "feature-import":
		fallthrough
	case "feature-tagged":
		fallthrough
	case "feature-tag-import":
		fallthrough
	case "feature-strategy-update":
		fallthrough
	case "feature-strategy-add":
		fallthrough
	case "feature-strategy-remove":
		fallthrough
	case "strategy-order-changed":
		fallthrough
	case "drop-feature-tags":
		fallthrough
	case "feature-untagged":
		fallthrough
	case "feature-stale-on":
		fallthrough
	case "feature-stale-off":
		fallthrough
	case "drop-features":
		fallthrough
	case "feature-environment-enabled":
		fallthrough
	case "feature-environment-disabled":
		fallthrough
	case "strategy-created":
		fallthrough
	case "strategy-deleted":
		fallthrough
	case "strategy-deprecated":
		fallthrough
	case "strategy-reactivated":
		fallthrough
	case "strategy-updated":
		fallthrough
	case "strategy-import":
		fallthrough
	case "drop-strategies":
		fallthrough
	case "context-field-created":
		fallthrough
	case "context-field-updated":
		fallthrough
	case "context-field-deleted":
		fallthrough
	case "project-access-added":
		fallthrough
	case "project-created":
		fallthrough
	case "project-updated":
		fallthrough
	case "project-deleted":
		fallthrough
	case "project-import":
		fallthrough
	case "project-user-added":
		fallthrough
	case "project-user-removed":
		fallthrough
	case "project-user-role-changed":
		fallthrough
	case "project-group-role-changed":
		fallthrough
	case "project-group-added":
		fallthrough
	case "project-group-removed":
		fallthrough
	case "drop-projects":
		fallthrough
	case "tag-created":
		fallthrough
	case "tag-deleted":
		fallthrough
	case "tag-import":
		fallthrough
	case "drop-tags":
		fallthrough
	case "tag-type-created":
		fallthrough
	case "tag-type-deleted":
		fallthrough
	case "tag-type-updated":
		fallthrough
	case "tag-type-import":
		fallthrough
	case "drop-tag-types":
		fallthrough
	case "addon-config-created":
		fallthrough
	case "addon-config-updated":
		fallthrough
	case "addon-config-deleted":
		fallthrough
	case "db-pool-update":
		fallthrough
	case "user-created":
		fallthrough
	case "user-updated":
		fallthrough
	case "user-deleted":
		fallthrough
	case "drop-environments":
		fallthrough
	case "environment-import":
		fallthrough
	case "segment-created":
		fallthrough
	case "segment-updated":
		fallthrough
	case "segment-deleted":
		fallthrough
	case "group-created":
		fallthrough
	case "group-updated":
		fallthrough
	case "setting-created":
		fallthrough
	case "setting-updated":
		fallthrough
	case "setting-deleted":
		fallthrough
	case "client-metrics":
		fallthrough
	case "client-register":
		fallthrough
	case "pat-created":
		fallthrough
	case "pat-deleted":
		fallthrough
	case "public-signup-token-created":
		fallthrough
	case "public-signup-token-user-added":
		fallthrough
	case "public-signup-token-updated":
		fallthrough
	case "change-request-created":
		fallthrough
	case "change-request-discarded":
		fallthrough
	case "change-added":
		fallthrough
	case "change-discarded":
		fallthrough
	case "change-edited":
		fallthrough
	case "change-request-approved":
		fallthrough
	case "change-request-approval-added":
		fallthrough
	case "change-request-cancelled":
		fallthrough
	case "change-request-sent-to-review":
		fallthrough
	case "change-request-applied":
		fallthrough
	case "api-token-created":
		fallthrough
	case "api-token-updated":
		fallthrough
	case "api-token-deleted":
		fallthrough
	case "feature-favorited":
		fallthrough
	case "feature-unfavorited":
		fallthrough
	case "project-favorited":
		fallthrough
	case "project-unfavorited":
		fallthrough
	case "features-exported":
		fallthrough
	case "features-imported":
		fallthrough
	case "service-account-created":
		fallthrough
	case "service-account-deleted":
		fallthrough
	case "service-account-updated":
		fallthrough
	case "feature-potentially-stale-on":
		*e = SearchEventsSchemaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SearchEventsSchemaType: %v", v)
	}
}

// SearchEventsSchema -
//
//	Search for events by type, project, feature, free-text query,
//	or a combination thereof. Pass an empty object to fetch all events.
type SearchEventsSchema struct {
	// Find events by feature toggle name (case-sensitive).
	Feature *string `json:"feature,omitempty"`
	// The maximum amount of events to return in the search result
	Limit *int64 `json:"limit,omitempty"`
	// Which event id to start listing from
	Offset *int64 `json:"offset,omitempty"`
	// Find events by project ID (case-sensitive).
	Project *string `json:"project,omitempty"`
	//                 Find events by a free-text search query.
	//                 The query will be matched against the event type,
	//                 the username or email that created the event (if any),
	//                 and the event data payload (if any).
	//
	Query *string `json:"query,omitempty"`
	// Find events by event type (case-sensitive).
	Type *SearchEventsSchemaType `json:"type,omitempty"`
}
