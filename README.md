# refresh-hash

The application contains two API servers: gRPC and HTTP. Read full [description](https://github.com/dolefir/refresh-hash/blob/main/task/DESCRIPTION.md).

## Run

__Prerequisites:__
* [Docker & docker-compose](https://www.docker.com/products/docker-desktop) (for 'Docker-compose' approach);
* [Golang](https://golang.org/dl/) v1.20+ (for 'Without Docker' approach).

#### Docker-compose

1. Config file already fill.
2. `$ make compose-up`

#### Without Docker

1. Config file already fill.
2. `$ make run`

#### HTTP server

1. Listen address `localhost:8080`

#### gRPC server

1. Listen address `localhost:8081`

### Run the test

1. `$ make test`
