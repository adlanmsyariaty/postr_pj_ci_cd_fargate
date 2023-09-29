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

1. Run the command below to start the Docker Compose:

    ```bash
    docker-compose up --build
    ```

2. Install `golang-migrate` and run the command below to migrate all the required tables:

    ```bash
    make migrateup
    ```

   The service will be available at `http://localhost:8080`.
