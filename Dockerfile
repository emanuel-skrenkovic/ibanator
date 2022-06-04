FROM golang:1.18.3 AS build
ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go mod download && \
go build -o ibanator cmd/ibanator/main.go

FROM alpine
WORKDIR /
COPY --from=build /app/ibanator /ibanator
EXPOSE 8080
ENTRYPOINT ["/ibanator"]
