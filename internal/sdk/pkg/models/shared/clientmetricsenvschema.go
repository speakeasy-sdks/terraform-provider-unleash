// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// ClientMetricsEnvSchema - Used for reporting feature evaluation results from SDKs
type ClientMetricsEnvSchema struct {
	// The name of the application the SDK is being used in
	AppName string `json:"appName"`
	// Which environment the SDK is being used in
	Environment string `json:"environment"`
	// Name of the feature checked by the SDK
	FeatureName string `json:"featureName"`
	// How many times the toggle evaluated to false
	No        *int64      `json:"no,omitempty"`
	Timestamp *DateSchema `json:"timestamp,omitempty"`
	// How many times each variant was returned
	Variants map[string]int64 `json:"variants,omitempty"`
	// How many times the toggle evaluated to true
	Yes *int64 `json:"yes,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _ClientMetricsEnvSchema ClientMetricsEnvSchema

func (c *ClientMetricsEnvSchema) UnmarshalJSON(bs []byte) error {
	data := _ClientMetricsEnvSchema{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = ClientMetricsEnvSchema(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "appName")
	delete(additionalFields, "environment")
	delete(additionalFields, "featureName")
	delete(additionalFields, "no")
	delete(additionalFields, "timestamp")
	delete(additionalFields, "variants")
	delete(additionalFields, "yes")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c ClientMetricsEnvSchema) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_ClientMetricsEnvSchema(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}
