FROM golang:alpine AS build-env

RUN apk add git

WORKDIR /go/src/math_app/

COPY go.mod .
COPY go.sum .
RUN go mod download

ENV CGO_ENABLED=0

COPY . .
RUN go test -cover ./...
RUN go install ./cmd/math_server


FROM alpine:latest

COPY --from=build-env /go/bin/math_server /app/math_server
ENTRYPOINT /app/math_server