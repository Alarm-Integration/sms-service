FROM golang:1.16 as builder
ENV APP_USER app
ENV APP_HOME /go/src/myapp
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
WORKDIR $APP_HOME
USER $APP_USER
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o main

FROM debian:buster
ENV APP_USER app
ENV APP_HOME /go/src/myapp
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY --chown=0:0 --from=builder $APP_HOME/main $APP_HOME
USER $APP_USER
CMD ["./main"]