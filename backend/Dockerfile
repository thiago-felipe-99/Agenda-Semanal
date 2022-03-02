FROM golang:1.17.7-alpine as build

COPY . /build

WORKDIR /build

RUN go build -o main

FROM heroku/heroku:20

COPY --from=build /build/main /app/main

ENV HOME /app
WORKDIR /app
RUN useradd -m heroku
USER heroku

ENTRYPOINT ["./main"]
