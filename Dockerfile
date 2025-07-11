FROM golang:alpine AS build
WORKDIR /buildapp
COPY . .
RUN go mod download
RUN go build -o becrud main.go

FROM alpine:3.22
WORKDIR /app
COPY --from=build /buildapp/becrud /app/becrud
COPY db.sql /app/db.sql  

EXPOSE 8080
ENTRYPOINT ["/app/becrud"]
