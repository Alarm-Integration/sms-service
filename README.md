## Run App

```bash
## run app
go run main.go
```

## Run Test

```bash
# run test
go test -cover ./...

# run test including coverage measurement
go test -coverprofile=coverage.out ./...


# measure test coverage by using each function
go tool cover -func=coverage.out

# measure test coverage by using html
go tool cover -html=coverage.out
```

## Docker

```bash
# docker image build
docker build . -t greatlaboratory/sms-service

# docker run container
docker run --rm -d -p 8081:8081 --name sms-service greatlaboratory/sms-service
```