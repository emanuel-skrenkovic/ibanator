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

[Docker Compose][3] is also available. To run the application using DockerCompose, run the following command in the project root directory:

``` sh
docker-compose up -d
```

## Requirements:
* [Go][1] (written in 1.18.3)
* [Docker][2]
* [Docker-Compose][3] (optional)
* [make][4] (optional)

## API definition
The API contains a [Swagger][5] definition which is accessible at http://localhost:8080/swagger/index.html while the application is running.

## Make targets
* **build** - Builds the application. Binary is located at ./bin/ibanator
* **run** - Runs the application.
* **test** - Runs unit tests.
* **build-image** - Builds the Docker image. The image is tagged with ibanator:latest.

[1]: https://go.dev/
[2]: https://www.docker.com/
[3]: https://docs.docker.com/compose/
[4]: https://www.gnu.org/software/make/
[5]: https://swagger.io
