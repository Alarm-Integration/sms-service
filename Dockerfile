FROM golang:1.16 as builder
ENV APP_HOME /go/src/myapp
WORKDIR $APP_HOME
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o main

FROM debian:buster
ENV APP_HOME /go/src/myapp
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY --chown=0:0 --from=builder $APP_HOME/main $APP_HOME
CMD ["./main"]