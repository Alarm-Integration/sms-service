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

# gocovey - localhost:8080
gocovey
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
- https://github.com/callistaenterprise/goblog/tree/P8/accountservice
- https://github.com/streadway/amqp/blob/master/_examples/simple-consumer/consumer.go

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

## SMS API
- [cool-sms](https://coolsms.co.kr/)
- sample request
```json
{
    "messages": [
        {
            "to": "01000000001",
            "from": "029302266",
            "text": "내용",
            "type": "SMS"
        }
    ]
}
```

- sample response
```json
{
    "count": {
        "total": 1,
        "sentTotal": 0,
        "sentFailed": 0,
        "sentSuccess": 0,
        "sentPending": 0,
        "sentReplacement": 0,
        "refund": 0,
        "registeredFailed": 0,
        "registeredSuccess": 1
    },
    "countForCharge": {
        "sms": {
            "82": 1
        },
        "lms": {},
        "mms": {},
        "ata": {},
        "cta": {},
        "cti": {}
    },
    "balance": {
        "requested": 0,
        "replacement": 0,
        "refund": 0,
        "sum": 0
    },
    "point": {
        "requested": 50,
        "replacement": 0,
        "refund": 0,
        "sum": 0
    },
    "app": {
        "profit": {
            "sms": 0,
            "lms": 0,
            "mms": 0,
            "ata": 0,
            "cta": 0,
            "cti": 0
        },
        "appId": null,
        "version": null
    },
    "serviceMethod": "MT",
    "sdkVersion": null,
    "osPlatform": null,
    "log": [
        {
            "createAt": "2021-01-23T10:47:52.771Z",
            "message": "[::ffff:127.0.0.1] 메시지 그룹이 생성되었습니다."
        },
        {
            "createAt": "2021-01-23T10:47:52.788Z",
            "message": "국가코드(82)의 단문문자(SMS) 1 건이 추가되었습니다."
        },
        {
            "createAt": "2021-01-23T10:47:52.803Z",
            "message": "메시지를 발송했습니다.",
            "oldBalance": 100,
            "newBalance": 100,
            "oldPoint": 100,
            "newPoint": 50,
            "totalPrice": 0
        }
    ],
    "status": "SENDING",
    "dateSent": "2021-01-23T10:47:52.803Z",
    "dateCompleted": null,
    "isRefunded": false,
    "flagUpdated": false,
    "prepaid": true,
    "strict": false,
    "masterAccountId": null,
    "_id": "G4V20210123194752TX3BUSQGG4ESXWK",
    "accountId": "12925149",
    "apiVersion": "4",
    "customFields": {},
    "hint": null,
    "groupId": "G4V20210123194752TX3BUSQGG4ESXWK",
    "price": {
        "82": []
    },
    "dateCreated": "2021-01-23T10:47:52.773Z",
    "dateUpdated": "2021-01-23T10:47:52.803Z"
}
```

- sample code
```go
package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
)

func main() {
  uri := "http://api.coolsms.co.kr/messages/v4/send-many"
  data := strings.NewReader(`{"messages":[{"to":"01000000001","from":"029302266","text":"내용","type":"SMS"}]}`)

  req, err := http.NewRequest("POST", uri, data)
  if err != nil { 
  	panic(err)
  }

  req.Header.Set("Authorization", "HMAC-SHA256 apiKey=NCSAYU7YDBXYORXC, date=2019-07-01T00:41:48Z, salt=jqsba2jxjnrjor, signature=1779eac71a24cbeeadfa7263cb84b7ea0af1714f5c0270aa30ffd34600e363b4")
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil { 
  	panic(err)
  }
  defer resp.Body.Close()

  bytes, _ := ioutil.ReadAll(resp.Body)
  str := string(bytes)
  fmt.Println(str)
}
```