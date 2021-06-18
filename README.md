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

## Gitlab Runner

```bash
# docker run gitlab-runner container
docker run 
--detach \
--name gitlab-runner \
--restart always \
--volume /srv/gitlab-runner/config:/etc/gitlab-runner: \
--volume /var/run/docker.sock:/var/run/docker.sock \
gitlab/gitlab-runner:latest

# docker execute gitlab-runner container by using bash
docker exec -it gitlab-runner bash

# register gitlab repository for CI/CD
gitlab-runner register -n \
--url [registration url] \
--registration-token [registration token] \
--description gitlab-runner \
--executor docker \
--docker-image docker:latest \
--docker-volumes /var/run/docker.sock:/var/run/docker.sock
```