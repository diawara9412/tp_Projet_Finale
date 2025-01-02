
FROM golang:1.23 AS build
WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go build -o vote-api


FROM ubuntu:22.04
WORKDIR /app


RUN apt-get update && apt-get install -y \
    libc6 \
    ca-certificates


COPY --from=build /app/vote-api /app/vote-api


RUN chmod +x /app/vote-api


EXPOSE 8080


CMD ["/app/vote-api"]