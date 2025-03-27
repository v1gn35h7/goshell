# Build stage
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /goshell ./cmd/goshell/


# Artifact stage
FROM alpine:3.14

WORKDIR /

COPY --from=build-stage /goshell /goshell


EXPOSE 8080
EXPOSE 8082

ENTRYPOINT ["/goshell", "start", "--configPath=/etc/goshell/configs/goshell"]
