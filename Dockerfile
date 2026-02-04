FROM golang:1.25.6 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/web ./cmd/web

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=build /app/bin/web /app/web
COPY config.json /app/config.json

EXPOSE 3000

CMD ["/app/web"]
