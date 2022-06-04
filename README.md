# IBANator 

## How to run
To run the API, use the followowing command:
``` sh
make run
```
If make is not installed, while being in project root, run the following command:
``` sh
go run cmd/ibanator/main.go
```

## Make targets
* **build** - Builds the application. Binary is located at ./bin/ibanator
* **run** - Runs the application.
* **test** - Runs unit tests.
* **build-image** - Builds the Docker image. The image is tagged with ibanator:latest.

## Requirements:
* [Go](https://go.dev/) (written in 1.18.3)
* [Docker](https://www.docker.com/)
* [make](https://www.gnu.org/software/make/) (optional)
