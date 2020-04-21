
# Sample Golang CSV Report Server

Welcome to a sample csv report generator developed with golang

# Pre-Requisites for Development
* Go 1.14.x
* Docker-CE version 19.03.5 and above
* Docker-compose version 1.22.*
* Git
* Visual Studio Code **(with Golang language server)**

# Test Deployment

## Server Software Requirements
* Docker-CE version 1.18.3 and above
* Docker-compose version 1.25.* and above
* Git
* Go 1.14.x

## Deploy Postgresql
1. ``cd`` to the project directory
2. Run:

    * ``docker pull postgres:11.2`` to pull the postgresql docker image
    * ``docker-compose -f docker-compose-pq.yaml up -d`` to bring up the postgresql db
    * ``docker exec -it $(docker container ls -f name=postgresql* -q) bash -c "psql -U hppoc -v dbname=reportserverdb -v user=hppoc -v passwd=hppoc -f /scripts/createpg.sql"`` to setup the database and tables

## Run the HTTP Server
1. ``cd`` to the project directory
2. Update the STATIC_FILEPATH environment variable appropriately
3. Run:

    * ``go build ./...`` to get the module dependencies and build
    * ``go run ./cmd/reportserver/`` to run the http server

## Import the Postman Collection
* Import the postman collection file ``reportserver.postman_collection.json``

# Troubleshooting
