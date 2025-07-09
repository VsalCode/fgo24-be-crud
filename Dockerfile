FROM golang:1.22-alpine AS build

WORKDIR /buildapp
COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o goapp main.go

FROM alpine:3.20

WORKDIR /app
COPY --from=build /buildapp/goapp /app/goapp
COPY .env /app/.env

RUN chmod +x /app/goapp

EXPOSE 8080

ENTRYPOINT ["/app/goapp"]