# Build stage
FROM golang:1.25 AS build-stage

WORKDIR /app

# Artifact stage
FROM alpine:3.14

WORKDIR /

COPY goshell /goshell


EXPOSE 8080
EXPOSE 8082

ENTRYPOINT ["/goshell", "start", "--configPath=/etc/goshell/configs/goshell"]
