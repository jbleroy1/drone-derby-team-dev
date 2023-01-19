```
██████╗░██████╗░░█████╗░███╗░░██╗███████╗  ██████╗░███████╗██████╗░██████╗░██╗░░░██╗
██╔══██╗██╔══██╗██╔══██╗████╗░██║██╔════╝  ██╔══██╗██╔════╝██╔══██╗██╔══██╗╚██╗░██╔╝
██║░░██║██████╔╝██║░░██║██╔██╗██║█████╗░░  ██║░░██║█████╗░░██████╔╝██████╦╝░╚████╔╝░
██║░░██║██╔══██╗██║░░██║██║╚████║██╔══╝░░  ██║░░██║██╔══╝░░██╔══██╗██╔══██╗░░╚██╔╝░░
██████╔╝██║░░██║╚█████╔╝██║░╚███║███████╗  ██████╔╝███████╗██║░░██║██████╦╝░░░██║░░░
╚═════╝░╚═╝░░╚═╝░╚════╝░╚═╝░░╚══╝╚══════╝  ╚═════╝░╚══════╝╚═╝░░╚═╝╚═════╝░░░░╚═╝░░░
```

# Dev Environment

## Build Image

Build the server image with : 

```
docker build -f build/Dockerfile -t drone-derby-server .
```

or use docker-compose

```
docker-compose -f deployments/docker-compose.yml build
```

## Launch application locally

The application is composed of 2 services : 
* Golang Server
* Firestore database

You can start the stack locally thanks to docker-compose with :
```
PROJECT_ID=<your projectID> docker-compose -f deployments/docker-compose.yml up
```

## Configuration

There is some configuration you can provide to the server as environment variables

| Env Variable Name       | Description                                        |
|-------------------------|----------------------------------------------------|
| PROJECT_ID              | Your GCP project ID                                |
| ENV                     | If == DEV, the method name is logged by the logger |
| SWAGGER_UI_ENABLED      | if == true, enable the swagger UI                  |
| FIRESTORE_EMULATOR_HOST | Set the Firestore emulator host during development |

# Access the Service

When the service is started, you can access the Swagger UI [here](http://localhost:8080/swagger-ui/)

Then you have access to all the API endpoint. 

## Generation go-server from OpenAPI Spec

The drone-derby-server application implements an [OAS3](https://swagger.io/specification/) specification. You can read the specification in the api folder.
```

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/api/openapi.yaml \
    -g go-server  \
    --additional-properties=outputAsLibrary=true,sourceFolder=pkg/openapi  \
    --skip-validate-spec \
    -o /local
```


## Deploy on Cloud Run

//TODO