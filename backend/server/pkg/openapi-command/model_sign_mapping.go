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

type SignMapping struct {

	Id string `json:"id,omitempty"`

	Direction Direction `json:"direction"`

	Altitude Altitude `json:"altitude"`

	Speed string `json:"speed"`

	NextAction string `json:"nextAction"`

	ImageUrl string `json:"imageUrl"`

	Description string `json:"description,omitempty"`

	CreationDate string `json:"creationDate,omitempty"`

	UpdateTime string `json:"updateTime,omitempty"`
}

// AssertSignMappingRequired checks if the required fields are not zero-ed
func AssertSignMappingRequired(obj SignMapping) error {
	elements := map[string]interface{}{
		"direction": obj.Direction,
		"altitude": obj.Altitude,
		"speed": obj.Speed,
		"nextAction": obj.NextAction,
		"imageUrl": obj.ImageUrl,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertDirectionRequired(obj.Direction); err != nil {
		return err
	}
	if err := AssertAltitudeRequired(obj.Altitude); err != nil {
		return err
	}
	return nil
}

// AssertRecurseSignMappingRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SignMapping (e.g. [][]SignMapping), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSignMappingRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSignMapping, ok := obj.(SignMapping)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSignMappingRequired(aSignMapping)
	})
}
