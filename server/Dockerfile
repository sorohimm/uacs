FROM golang:1.17.1-alpine3.14 as build-env

WORKDIR /server/cmd
COPY . /server

RUN go build -o uacs .

# Final stage
WORKDIR /
FROM alpine:latest
COPY --from=build-env server/cmd/uacs /server/cmd/uacs

ENTRYPOINT ["server/cmd/uacs"]