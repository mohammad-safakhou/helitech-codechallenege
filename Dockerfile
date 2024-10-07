FROM golang:1.23.2 AS build

RUN mkdir /app

WORKDIR /app

COPY . .
RUN GOPROXY=goproxy.io,direct CGO_ENABLED=0 go build -o app

FROM alpine:3.9

COPY --from=build /app/app app
COPY --from=build /app/config/config.json .
RUN chmod +x /app
