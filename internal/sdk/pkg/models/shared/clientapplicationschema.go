// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

type ClientApplicationSchemaStartedType string

const (
	ClientApplicationSchemaStartedTypeDateTime ClientApplicationSchemaStartedType = "date-time"
	ClientApplicationSchemaStartedTypeNumber   ClientApplicationSchemaStartedType = "number"
)

type ClientApplicationSchemaStarted struct {
	DateTime *time.Time
	Number   *float64

	Type ClientApplicationSchemaStartedType
}

func CreateClientApplicationSchemaStartedDateTime(dateTime time.Time) ClientApplicationSchemaStarted {
	typ := ClientApplicationSchemaStartedTypeDateTime

	return ClientApplicationSchemaStarted{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateClientApplicationSchemaStartedNumber(number float64) ClientApplicationSchemaStarted {
	typ := ClientApplicationSchemaStartedTypeNumber

	return ClientApplicationSchemaStarted{
		Number: &number,
		Type:   typ,
	}
}

func (u *ClientApplicationSchemaStarted) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	dateTime := new(time.Time)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&dateTime); err == nil {
		u.DateTime = dateTime
		u.Type = ClientApplicationSchemaStartedTypeDateTime
		return nil
	}

	number := new(float64)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&number); err == nil {
		u.Number = number
		u.Type = ClientApplicationSchemaStartedTypeNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ClientApplicationSchemaStarted) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return json.Marshal(u.DateTime)
	}

	if u.Number != nil {
		return json.Marshal(u.Number)
	}

	return nil, nil
}

// ClientApplicationSchema - A client application is an instance of one of our SDKs
type ClientApplicationSchema struct {
	// An identifier for the app that uses the sdk, should be static across SDK restarts
	AppName string `json:"appName"`
	// The SDK's configured 'environment' property. Deprecated. This property  does **not** control which Unleash environment the SDK gets toggles for. To control Unleash environments, use the SDKs API key.
	//
	// @deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Environment *string `json:"environment,omitempty"`
	// A unique identifier identifying the instance of the application running the SDK. Often changes based on execution environment. For instance: two pods in Kubernetes will have two different instanceIds
	InstanceID *string `json:"instanceId,omitempty"`
	// How often (in seconds) does the client refresh its toggles
	Interval float64 `json:"interval"`
	// An SDK version identifier. Usually formatted as "unleash-client-<language>:<version>"
	SDKVersion *string `json:"sdkVersion,omitempty"`
	// Either an RFC-3339 timestamp or a unix timestamp in seconds
	Started ClientApplicationSchemaStarted `json:"started"`
	// Which strategies the SDKs runtime knows about
	Strategies []string `json:"strategies"`
}
