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
docker run --rm -d -p 8081:8081 --name sms-service --network ai-network greatlaboratory/sms-service
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

## load config file from spring config server
- https://github.com/spf13/viper
- https://dzone.com/articles/go-microservices-part-8-centralized-configuration

## kafka produce & consume test
```bash
//토픽 리스트 확인
./kafka-topics.sh --bootstrap-server 139.150.75.240:9093 --list
__confluent.support.metrics


//토픽 생성
./kafka-topics.sh --bootstrap-server 139.150.75.240:9093 --create --topic email --partitions 3 --replication-factor 3
./kafka-topics.sh --bootstrap-server 139.150.75.240:9093 --create --topic slack --partitions 3 --replication-factor 2
./kafka-topics.sh --bootstrap-server 139.150.75.240:9094 --create --topic sms --partitions 3 --replication-factor 2

//토픽 리스트 재확인
./kafka-topics.sh --bootstrap-server 139.150.75.240:9093 --list
__confluent.support.metrics
email
slack
sms

//sms 토픽 상세 정보 확인
./kafka-topics.sh --bootstrap-server 139.150.75.240:9093 --describe --topic sms
Topic: sms    PartitionCount: 3       ReplicationFactor: 3    Configs:
        Topic: sms    Partition: 0    Leader: 2       Replicas: 2,3,1 Isr: 2,3,1
        Topic: sms    Partition: 1    Leader: 3       Replicas: 3,1,2 Isr: 3,1,2
        Topic: sms    Partition: 2    Leader: 1       Replicas: 1,2,3 Isr: 1,2,3
 
//sms 토픽 상세 정보 확인
./kafka-topics.sh --bootstrap-server 139.150.75.240:9093 --describe --topic sms
Topic: sms      PartitionCount: 3       ReplicationFactor: 2    Configs:
        Topic: sms      Partition: 0    Leader: 2       Replicas: 2,1   Isr: 2,1
        Topic: sms      Partition: 1    Leader: 3       Replicas: 3,2   Isr: 3,2
        Topic: sms      Partition: 2    Leader: 1       Replicas: 1,3   Isr: 1,3
        
// 컨슈머 그룹 목록 확인
./kafka-consumer-groups.sh --bootstrap-server 139.150.75.240:9093 --list

//컨슈머 상태와 오프셋 확인
./kafka-consumer-groups.sh --bootstrap-server 139.150.75.240:9093 --group sms --describe

// sms 토픽으로 레코드 프로듀싱하기
./kafka-console-producer.sh --bootstrap-server 139.150.75.240:9093 --topic sms
```