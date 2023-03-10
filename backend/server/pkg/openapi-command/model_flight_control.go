/*
 * Drone-Command-API
 *
 * Drone Command API
 *
 * API version: 1.0
 * Contact: drone-derby-eng-team@google.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_command

import (
	"time"
)

type FlightControl struct {

	Direction Direction `json:"direction"`

	Action string `json:"action"`

	Date time.Time `json:"date,omitempty"`
}

// AssertFlightControlRequired checks if the required fields are not zero-ed
func AssertFlightControlRequired(obj FlightControl) error {
	elements := map[string]interface{}{
		"direction": obj.Direction,
		"action": obj.Action,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertDirectionRequired(obj.Direction); err != nil {
		return err
	}
	return nil
}

// AssertRecurseFlightControlRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of FlightControl (e.g. [][]FlightControl), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseFlightControlRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aFlightControl, ok := obj.(FlightControl)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertFlightControlRequired(aFlightControl)
	})
}
