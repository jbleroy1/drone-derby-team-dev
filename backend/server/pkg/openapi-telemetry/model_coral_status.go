/*
 * Drone-Telemetry-API
 *
 * Drone Telemetry API
 *
 * API version: 1.0
 * Contact: drone-derby-eng-team@google.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_telemetry

type CoralStatus struct {

	Uptime string `json:"uptime,omitempty"`

	Version string `json:"version,omitempty"`

	Cpu int32 `json:"cpu,omitempty"`

	Memory int32 `json:"memory,omitempty"`

	AdditionalPropertiesField []CoralStatusAdditionalPropertiesInner `json:"additionalProperties,omitempty"`
}

// AssertCoralStatusRequired checks if the required fields are not zero-ed
func AssertCoralStatusRequired(obj CoralStatus) error {
	for _, el := range obj.AdditionalPropertiesField {
		if err := AssertCoralStatusAdditionalPropertiesInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseCoralStatusRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CoralStatus (e.g. [][]CoralStatus), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCoralStatusRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCoralStatus, ok := obj.(CoralStatus)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCoralStatusRequired(aCoralStatus)
	})
}
