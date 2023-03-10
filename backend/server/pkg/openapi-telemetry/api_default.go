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

import (

	"net/http"
	"strings"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"GetDiagCoral",
			strings.ToUpper("Get"),
			"/diag/coral",
			c.GetDiagCoral,
		},
		{
			"GetDiagDrone",
			strings.ToUpper("Get"),
			"/diag/drone",
			c.GetDiagDrone,
		},
		{
			"GetDiagOperations",
			strings.ToUpper("Get"),
			"/diag/operations",
			c.GetDiagOperations,
		},
	}
}

// GetDiagCoral - GET Coral board status
func (c *DefaultApiController) GetDiagCoral(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetDiagCoral(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetDiagDrone - GET Drone Status
func (c *DefaultApiController) GetDiagDrone(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetDiagDrone(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetDiagOperations - GET Operations Details
func (c *DefaultApiController) GetDiagOperations(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetDiagOperations(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
