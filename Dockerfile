FROM golang:1.22.2-alpine3.19

WORKDIR /src/app

RUN apk add --no-cache git

# Install air
RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy