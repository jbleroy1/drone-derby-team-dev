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
	"os"

	"cloud.google.com/go/firestore"
	openapi_command "drone-server-api/pkg/openapi-command"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
)

var client *firestore.Client
var errorClient error
var  collectionMapping string =os.Getenv("DRONE_COLLECTION")

//Initialize Firestore connection
func init() {
	log.Printf("Initializing Firestore client")

	ctx := context.Background()
	client, errorClient = firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if errorClient != nil {
		log.Fatalf("Cannot init Firestore client",errorClient)
	}
	log.Printf("Firestore client connected")
}

func getAllMapping(ctx context.Context) []openapi_command.SignMapping {
	var signs  []openapi_command.SignMapping
	documentIterator := client.Collection(collectionMapping).Documents(ctx)
	defer documentIterator.Stop()
	for {
		doc, err := documentIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {

		}
		var s openapi_command.SignMapping
		if err := doc.DataTo(&s); err != nil {
		}
		signs = append(signs, s)
	}
	return signs
}

func getSignMappingById(ctx context.Context, id string) (*openapi_command.SignMapping,error) {
	get, err := client.Collection(collectionMapping).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var s *openapi_command.SignMapping
	if err := get.DataTo(&s); err != nil {
		return nil, err
	}
	return s,nil

}

func createSignMapping(ctx context.Context,sign *openapi_command.SignMapping) (*openapi_command.SignMapping,error) {
	docRef, _, err := client.Collection(collectionMapping).Add(ctx,sign)
	if err != nil {
		log.Fatalf("error during writing",err)
	}
	sign.Id=docRef.ID
	docRef.Set(ctx,sign)
	return sign,nil

}

func updateSignMapping(ctx context.Context,sign *openapi_command.SignMapping, id string) (*openapi_command.SignMapping,error) {

	_, err := client.Collection(collectionMapping).Doc(id).Set(ctx,sign)
	if err != nil {
		return nil, err
	}

	return sign,nil

}

