FROM golang:1.17.7-alpine

RUN apk update && apk upgrade
RUN apk add git

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
