FROM golang:1.19 AS build

WORKDIR /test
# Copy the go source
COPY . .
RUN GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .


FROM alpine:latest
COPY --from=build /test/pushtest /
ENTRYPOINT ["/pushtest"]