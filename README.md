# hr-service

## Overview
The `hr-service` is a Go-based application designed to manage employee attendance and leave records. It provides a set of RESTful APIs to handle various operations such as adding, updating, deleting, and querying attendance and leave data.

## Features
- Add, update, delete, and list attendance records
- Add, update, delete, and list leave records
- Query attendance and leave records
- Generate attendance statistics

## Requirements
- Go 1.18 or higher

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/z9905080/hr-service.git
    ```
2. Navigate to the project directory:
    ```sh
    cd hr-service
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage
### Docker Compose
1. Build the Docker image:
    ```sh
    docker-compose build
    ```
   
2. Start the Docker container:
    ```sh
    docker-compose up
    ```

### Local
1. Execute the following command to migrate the database schema:
    ```sh
    go run . migrate --config config/local.json
    ```
   
2. Start the application:
    ```sh
    go run . api --config config/local.json
    ```
   
3. The application will be available at `http://localhost:8080`.
4. To stop the application, press `Ctrl+C`.
5. To Rollback the database schema, execute the following command:
    ```sh
    go run . rollback --config config/local.json
    ```

## API Endpoints
link: [API Endpoints](./API_Endpoint.md)


## License
This project is licensed under the MIT License.