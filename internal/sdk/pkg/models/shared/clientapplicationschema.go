// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"time"
)

type StartedType string

const (
	StartedTypeDateTime StartedType = "date-time"
	StartedTypeNumber   StartedType = "number"
)

// Started - Either an RFC-3339 timestamp or a unix timestamp in seconds
type Started struct {
	DateTime *time.Time
	Number   *float64

	Type StartedType
}

func CreateStartedDateTime(dateTime time.Time) Started {
	typ := StartedTypeDateTime

	return Started{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateStartedNumber(number float64) Started {
	typ := StartedTypeNumber

	return Started{
		Number: &number,
		Type:   typ,
	}
}

func (u *Started) UnmarshalJSON(data []byte) error {

	dateTime := new(time.Time)
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = dateTime
		u.Type = StartedTypeDateTime
		return nil
	}

	number := new(float64)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = number
		u.Type = StartedTypeNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Started) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// ClientApplicationSchema - A client application is an instance of one of our SDKs
type ClientApplicationSchema struct {
	// An identifier for the app that uses the sdk, should be static across SDK restarts
	AppName string `json:"appName"`
	// The SDK's configured 'environment' property. Deprecated. This property  does **not** control which Unleash environment the SDK gets toggles for. To control Unleash environments, use the SDKs API key.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Environment *string `json:"environment,omitempty"`
	// A unique identifier identifying the instance of the application running the SDK. Often changes based on execution environment. For instance: two pods in Kubernetes will have two different instanceIds
	InstanceID *string `json:"instanceId,omitempty"`
	// How often (in seconds) does the client refresh its toggles
	Interval float64 `json:"interval"`
	// An SDK version identifier. Usually formatted as "unleash-client-<language>:<version>"
	SDKVersion *string `json:"sdkVersion,omitempty"`
	// Either an RFC-3339 timestamp or a unix timestamp in seconds
	Started Started `json:"started"`
	// Which strategies the SDKs runtime knows about
	Strategies []string `json:"strategies"`
}

func (o *ClientApplicationSchema) GetAppName() string {
	if o == nil {
		return ""
	}
	return o.AppName
}

func (o *ClientApplicationSchema) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *ClientApplicationSchema) GetInstanceID() *string {
	if o == nil {
		return nil
	}
	return o.InstanceID
}

func (o *ClientApplicationSchema) GetInterval() float64 {
	if o == nil {
		return 0.0
	}
	return o.Interval
}

func (o *ClientApplicationSchema) GetSDKVersion() *string {
	if o == nil {
		return nil
	}
	return o.SDKVersion
}

func (o *ClientApplicationSchema) GetStarted() Started {
	if o == nil {
		return Started{}
	}
	return o.Started
}

func (o *ClientApplicationSchema) GetStrategies() []string {
	if o == nil {
		return []string{}
	}
	return o.Strategies
}
