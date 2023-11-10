// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"time"
)

type ProxyClientSchemaStartedType string

const (
	ProxyClientSchemaStartedTypeDateTime ProxyClientSchemaStartedType = "date-time"
	ProxyClientSchemaStartedTypeNumber   ProxyClientSchemaStartedType = "number"
)

type ProxyClientSchemaStarted struct {
	DateTime *time.Time
	Number   *float64

	Type ProxyClientSchemaStartedType
}

func CreateProxyClientSchemaStartedDateTime(dateTime time.Time) ProxyClientSchemaStarted {
	typ := ProxyClientSchemaStartedTypeDateTime

	return ProxyClientSchemaStarted{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateProxyClientSchemaStartedNumber(number float64) ProxyClientSchemaStarted {
	typ := ProxyClientSchemaStartedTypeNumber

	return ProxyClientSchemaStarted{
		Number: &number,
		Type:   typ,
	}
}

func (u *ProxyClientSchemaStarted) UnmarshalJSON(data []byte) error {

	dateTime := new(time.Time)
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = dateTime
		u.Type = ProxyClientSchemaStartedTypeDateTime
		return nil
	}

	number := new(float64)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = number
		u.Type = ProxyClientSchemaStartedTypeNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ProxyClientSchemaStarted) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// ProxyClientSchema - Frontend SDK client registration information
type ProxyClientSchema struct {
	// Name of the application using Unleash
	AppName string `json:"appName"`
	// deprecated
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Environment *string `json:"environment,omitempty"`
	// Instance id for this application (typically hostname, podId or similar)
	InstanceID *string `json:"instanceId,omitempty"`
	// At which interval, in milliseconds, will this client be expected to send metrics
	Interval float64 `json:"interval"`
	// Optional field that describes the sdk version (name:version)
	SDKVersion *string `json:"sdkVersion,omitempty"`
	// When this client started. Should be reported as ISO8601 time.
	Started ProxyClientSchemaStarted `json:"started"`
	// List of strategies implemented by this application
	Strategies []string `json:"strategies"`
}

func (o *ProxyClientSchema) GetAppName() string {
	if o == nil {
		return ""
	}
	return o.AppName
}

func (o *ProxyClientSchema) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *ProxyClientSchema) GetInstanceID() *string {
	if o == nil {
		return nil
	}
	return o.InstanceID
}

func (o *ProxyClientSchema) GetInterval() float64 {
	if o == nil {
		return 0.0
	}
	return o.Interval
}

func (o *ProxyClientSchema) GetSDKVersion() *string {
	if o == nil {
		return nil
	}
	return o.SDKVersion
}

func (o *ProxyClientSchema) GetStarted() ProxyClientSchemaStarted {
	if o == nil {
		return ProxyClientSchemaStarted{}
	}
	return o.Started
}

func (o *ProxyClientSchema) GetStrategies() []string {
	if o == nil {
		return []string{}
	}
	return o.Strategies
}
