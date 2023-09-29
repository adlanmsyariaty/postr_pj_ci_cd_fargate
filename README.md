# PostrServer
Services related to posr application

## Development

### Requirements

To develop this service, you need to have the following installed in your machine:

1. Golang
2. Docker 20
3. Docker Compose 1.29

### Running the Service Locally

We use Docker Compose to run the service locally. To run the service, follow these steps:

1. We need to build the service first. Run the command below to build the service:

    ```bash
    ./gradlew bootJar
    ```

2. Run the command below to start the Docker Compose:

    ```bash
    docker-compose up --build
    ```

   The service will be available at `http://localhost:8080`.
