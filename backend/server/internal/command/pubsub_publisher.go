package command

import (
	"context"
	"os"
	"strconv"

	"cloud.google.com/go/pubsub"
	log "github.com/sirupsen/logrus"
)

var  PROJECT_ID string =os.Getenv("PUBSUB_PROJECT_ID")
var  TOPIC_ID string =os.Getenv("PUBSUB_TOPIC_ID")


// Create a global client variable
var pubsub_client *pubsub.Client

func init() {
	 ctx :=  context.Background()
	// Create a client
	log.Info("Build Pub Sub client for PROJECT "+PROJECT_ID+" and TOPIC "+TOPIC_ID)
	var err error
	pubsub_client, err = pubsub.NewClient(ctx, PROJECT_ID)
	createTopic, err := strconv.ParseBool(os.Getenv("CREATE_TOPIC_ID"))
	if createTopic {
		pubsub_client.CreateTopic(ctx,TOPIC_ID)
	}


	if err != nil {
		panic(err)
	}
}



func publishMessage(message []byte, err error) error {

	// Get a reference to the topic
	topic := pubsub_client.Topic(TOPIC_ID)
	ctx :=  context.Background()


	// Publish the message
	result := topic.Publish(ctx, &pubsub.Message{
		Data: message,
	})

	//Handle pubsub response
	_, err = result.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
