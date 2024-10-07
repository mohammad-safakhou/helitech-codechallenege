# Code Challenge

This is a Go application that implements a Todo service with a file upload feature, database migrations, SQS queue
integration, Docker setup, and unit testing & benchmarking.

## Prerequisites

- Docker
- Docker Compose
- Make
- Go (1.23 or newer)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/mohammad-safakhou/helitech-codechallenege
   ```

2. Build and run the Docker containers<pre>make run </pre>
## Usage

### API Endpoints

- `/v1/upload`: Allows users to upload a file. The uploaded file is stored in an S3 bucket, and a `fileId` is returned in the response.
- `/v1/todo`: Allows users to create a `TodoItem`. The request should include the description, dueDate, and an optional `fileId` if the user has uploaded a file.

### Running Tests

Run the following command to execute the unit tests:

```sh
make test
```

### Running Benchmarks

Run the following command to execute the benchmarks:
```sh
make benchmark
```

## Built With

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [PostgreSQL](https://www.postgresql.org/)
- [LocalStack](https://localstack.cloud/)
