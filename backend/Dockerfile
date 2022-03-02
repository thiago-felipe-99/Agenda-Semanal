FROM golang:1.17.7-alpine as build

COPY . /build/

WORKDIR /build

RUN go build -o main

FROM alpine

COPY --from=build /build/main /app/main

ENTRYPOINT ["/app/main"]
