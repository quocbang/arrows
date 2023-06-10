# Server binary builder
FROM golang:${go_version}-alpine${alpine_version} as base

RUN apk add --no-cache git
RUN apk add --update --no-cache curl build-base

RUN apk update && apk add openssh

RUN go mod download
RUN go vet ./...
RUN go test ./...