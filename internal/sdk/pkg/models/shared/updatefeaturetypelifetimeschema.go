// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateFeatureTypeLifetimeSchema - Data used when updating the lifetime of a [feature toggle type](https://docs.getunleash.io/reference/feature-toggle-types).
type UpdateFeatureTypeLifetimeSchema struct {
	// The new lifetime (in days) that you want to assign to the feature toggle type. If the value is `null` or `0`, then the feature toggles of that type will never be marked as potentially stale. Otherwise, they will be considered potentially stale after the number of days indicated by this property.
	LifetimeDays *int64 `json:"lifetimeDays"`
}

func (o *UpdateFeatureTypeLifetimeSchema) GetLifetimeDays() *int64 {
	if o == nil {
		return nil
	}
	return o.LifetimeDays
}
