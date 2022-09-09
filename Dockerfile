# syntax = docker/dockerfile:1.4

FROM golang:1.19.0 AS builder

ENV GOPRIVATE=dev.azure.com

RUN git config --global url."git@ssh.dev.azure.com:v3/ciaalicorp/".insteadof "https://dev.azure.com/ciaalicorp/"

RUN mkdir -p -m 0700 ~/.ssh && ssh-keyscan ssh.dev.azure.com >> ~/.ssh/known_hosts

WORKDIR /src

COPY go.* .
RUN --mount=type=ssh go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=ssh CGO_ENABLED=0 G00S=linux go build -o app main.go

FROM alpine AS dockerize
LABEL maintainer="Alicia <bot_alicia@alicorp.com.pe> (@aliciabot)"

COPY --from=builder /src/app .

CMD ["./app"]
