/*
 * Drone-Server-API
 *
 * Drone Server API
 *
 * API version: 1.0
 * Contact: drone-derby-eng-team@google.com
 */

package command

import (
	"context"
	"errors"
	"net/http"

	openapi_telemtry "drone-server-api/pkg/openapi-telemetry"

)

type TelemetryAPIService struct {
	openapi_telemtry.DefaultApiServicer
}

func NewTelemetryAPIService() openapi_telemtry.DefaultApiServicer {

	return &TelemetryAPIService{}
}
// GetDiagCoral - GET Coral board status
func (s *TelemetryAPIService) GetDiagCoral(ctx context.Context) (openapi_telemtry.ImplResponse, error) {
	// TODO - update GetDiagCoral with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, CoralStatus{}) or use other options such as http.Ok ...
	//return Response(200, CoralStatus{}), nil

	return openapi_telemtry.Response(http.StatusNotImplemented, nil), errors.New("GetDiagCoral method not implemented")
}

// GetDiagDrone - GET Drone Status
func (s *TelemetryAPIService) GetDiagDrone(ctx context.Context) (openapi_telemtry.ImplResponse, error) {
	// TODO - update GetDiagDrone with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DroneStatus{}) or use other options such as http.Ok ...
	//return Response(200, DroneStatus{}), nil

	return openapi_telemtry.Response(http.StatusNotImplemented, nil), errors.New("GetDiagDrone method not implemented")
}

// GetDiagOperations - GET Operations Details
func (s *TelemetryAPIService) GetDiagOperations(ctx context.Context) (openapi_telemtry.ImplResponse, error) {
	// TODO - update GetDiagOperations with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Operations{}) or use other options such as http.Ok ...
	//return Response(200, Operations{}), nil

	return openapi_telemtry.Response(http.StatusNotImplemented, nil), errors.New("GetDiagOperations method not implemented")
}

