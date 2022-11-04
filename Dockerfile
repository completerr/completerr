FROM node:16.14.2-alpine3.15 AS JS_BUILD
COPY webapp /webapp
WORKDIR webapp
RUN npm install && npm run build

FROM golang:1.18.1-alpine3.15 AS GO_BUILD
RUN apk add build-base
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.13.5
WORKDIR /completerr
COPY --from=JS_BUILD /webapp/out* ./webapp/
COPY --from=GO_BUILD /go/bin/server ./
COPY config.example.yaml ./
COPY entrypoint.sh ./
ENTRYPOINT "/completerr/entrypoint.sh"
