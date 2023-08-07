// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type PlaygroundStrategySchemaResultType string

const (
	PlaygroundStrategySchemaResultTypePlaygroundStrategySchemaResultAnyOf  PlaygroundStrategySchemaResultType = "playgroundStrategySchema_result_anyOf"
	PlaygroundStrategySchemaResultTypePlaygroundStrategySchemaResultAnyOf1 PlaygroundStrategySchemaResultType = "playgroundStrategySchema_result_anyOf_1"
)

type PlaygroundStrategySchemaResult struct {
	PlaygroundStrategySchemaResultAnyOf  *PlaygroundStrategySchemaResultAnyOf
	PlaygroundStrategySchemaResultAnyOf1 *PlaygroundStrategySchemaResultAnyOf1

	Type PlaygroundStrategySchemaResultType
}

func CreatePlaygroundStrategySchemaResultPlaygroundStrategySchemaResultAnyOf(playgroundStrategySchemaResultAnyOf PlaygroundStrategySchemaResultAnyOf) PlaygroundStrategySchemaResult {
	typ := PlaygroundStrategySchemaResultTypePlaygroundStrategySchemaResultAnyOf

	return PlaygroundStrategySchemaResult{
		PlaygroundStrategySchemaResultAnyOf: &playgroundStrategySchemaResultAnyOf,
		Type:                                typ,
	}
}

func CreatePlaygroundStrategySchemaResultPlaygroundStrategySchemaResultAnyOf1(playgroundStrategySchemaResultAnyOf1 PlaygroundStrategySchemaResultAnyOf1) PlaygroundStrategySchemaResult {
	typ := PlaygroundStrategySchemaResultTypePlaygroundStrategySchemaResultAnyOf1

	return PlaygroundStrategySchemaResult{
		PlaygroundStrategySchemaResultAnyOf1: &playgroundStrategySchemaResultAnyOf1,
		Type:                                 typ,
	}
}

func (u *PlaygroundStrategySchemaResult) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	playgroundStrategySchemaResultAnyOf := new(PlaygroundStrategySchemaResultAnyOf)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&playgroundStrategySchemaResultAnyOf); err == nil {
		u.PlaygroundStrategySchemaResultAnyOf = playgroundStrategySchemaResultAnyOf
		u.Type = PlaygroundStrategySchemaResultTypePlaygroundStrategySchemaResultAnyOf
		return nil
	}

	playgroundStrategySchemaResultAnyOf1 := new(PlaygroundStrategySchemaResultAnyOf1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&playgroundStrategySchemaResultAnyOf1); err == nil {
		u.PlaygroundStrategySchemaResultAnyOf1 = playgroundStrategySchemaResultAnyOf1
		u.Type = PlaygroundStrategySchemaResultTypePlaygroundStrategySchemaResultAnyOf1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PlaygroundStrategySchemaResult) MarshalJSON() ([]byte, error) {
	if u.PlaygroundStrategySchemaResultAnyOf != nil {
		return json.Marshal(u.PlaygroundStrategySchemaResultAnyOf)
	}

	if u.PlaygroundStrategySchemaResultAnyOf1 != nil {
		return json.Marshal(u.PlaygroundStrategySchemaResultAnyOf1)
	}

	return nil, nil
}