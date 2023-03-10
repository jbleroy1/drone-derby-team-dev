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

type Command struct {

	Id string `json:"id,omitempty"`
}

// AssertCommandRequired checks if the required fields are not zero-ed
func AssertCommandRequired(obj Command) error {
	return nil
}

// AssertRecurseCommandRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Command (e.g. [][]Command), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCommandRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCommand, ok := obj.(Command)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCommandRequired(aCommand)
	})
}
