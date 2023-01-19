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
	"net/http"

	openapi_command "drone-server-api/pkg/openapi-command"
	"team.git.corp.google.com/drone-derby-eng-team/external/gopkgs/dcuevents"
)

type DataStoreAPIService struct {
	openapi_command.DefaultApiServicer
}

func NewDataStoreAPIService() openapi_command.DefaultApiServicer {

	return &DataStoreAPIService{}
}

func (s *DataStoreAPIService) GetSignId(ctx context.Context, id string ) (openapi_command.ImplResponse, error) {

	signMappingById, err := getSignMappingById(ctx, id)
	if err != nil {
		response := openapi_command.ImplResponse{
			Code: http.StatusNotFound,
			Body: nil,
		}
		return response, err
	}
	response := openapi_command.ImplResponse{
		Code: http.StatusOK,
		Body: signMappingById,
	}


	return response, nil
}

func (s *DataStoreAPIService) GetSign(ctx context.Context) (openapi_command.ImplResponse, error) {

	response := openapi_command.ImplResponse{
		Code: http.StatusOK,
		Body: getAllMapping(ctx),
	}

	return response, nil
}

func (s *DataStoreAPIService) PostSignId(ctx context.Context, id string , sign openapi_command.SignMapping) (openapi_command.ImplResponse, error){
	signMapping, err := createSignMapping(ctx, &sign)
	if err != nil {
		response := openapi_command.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: nil,
		}
		return response, err
	}
	response := openapi_command.ImplResponse{
		Code: http.StatusOK,
		Body: signMapping,
	}
	return response, nil
}
func (s *DataStoreAPIService) PutSignId(ctx context.Context,id  string, sign openapi_command.SignMapping) (openapi_command.ImplResponse, error){
	signMapping, err := updateSignMapping(ctx, &sign,id)
	if err != nil {
		response := openapi_command.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: nil,
		}
		return response, err
	}
	response := openapi_command.ImplResponse{
		Code: http.StatusOK,
		Body: signMapping,
	}
	return response, nil
}
// DeleteSignId - DELETE SignMapping
func (s *DataStoreAPIService) DeleteSignId(ctx context.Context, id string) (openapi_command.ImplResponse, error) {
	signMapping, err := s.DeleteSignId(ctx,id)
	if err != nil {
		response := openapi_command.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: nil,
		}
		return response, err
	}
	response := openapi_command.ImplResponse{
		Code: http.StatusOK,
		Body: signMapping,
	}
	return response, nil
}

// PostControl - POST Control order to send to the coral board
func (s *DataStoreAPIService) PostControl(ctx context.Context, flightControl openapi_command.FlightControl) (openapi_command.ImplResponse, error) {

	//forward the order to the CDU via pubsub
	var events=dcuevents.EventSendCommandPayload{
		Action: flightControl.Action,
		Angle: int(flightControl.Direction.Angle),
	}



	err := publishMessage(dcuevents.EncodeEvent(events))
	if err != nil {
		response := openapi_command.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: nil,
		}
		return response, err
	}
	response := openapi_command.ImplResponse{
		Code: http.StatusOK,
		Body: nil,
	}
	return response, nil
}
