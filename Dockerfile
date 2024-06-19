FROM golang:1.22 AS build

WORKDIR /app

RUN apt update && apt install -y ca-certificates git-core ssh

ARG CGO_ENABLED=0
ENV CGO_ENABLED=${CGO_ENABLED}

ARG GOOS=linux
ARG GOOS=${GOOS}

ARG GO111MODULE=on
ENV GO111MODULE=${GO111MODULE}

COPY go.mod go.sum ./
RUN go mod download && go mod tidy && go mod verify

COPY . .
RUN go build -o bin/api ./cmd/api

FROM gcr.io/distroless/static-debian11

WORKDIR /app
COPY --from=build --chown=nonroot:nonroot /app/bin /app/bin/
USER nonroot

ENTRYPOINT ["/app/bin/api"]
