## [개    요]
  - Cloud-Barista 활용 실습 기간: 3주 ~ 4주 예상
  - Cloud-Barista 실습 계획: CB-Spider => CB-Tumblebug
  - 3주: CB-Spider 설치/실행/활용

***

## [환경 준비]
```
       (1) GCP-VM1: 아래 그림 참조 후 동일 조건으로 신규 생성
           - 목적: CB-Spider 설치/가동, CLI 시험
           - VM1 로그인 위한 접속 Key 및 Terminal
       (2) AWS 계정의 Credential 정보(지난 미션시 생성한거 활용)
            - 목적: Spider에 입력 후 Spider로 AWS에 명령 요청시 사용
       (3) Web Browser : 노트북/데스크탑 등에 설치된 아무 브라우저 가능
            - 목적: CB-Spider AdminWeb 접속
```
## [Docker 기반 실행]  
### (1) Docker 설치
  - sudo apt install -y docker.io
    
### (2) Docker 간단 실습

```
# local docker image 목록
sudo docker image list
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE

# 공개 이미지 pull & 실행
sudo docker run -p 8080:80 --name web -d nginx

# local docker image 목록
sudo docker image list
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE
nginx                       latest              4bb46517cac3        6 days ago          133MB

# 컨테이너 종료
sudo docker stop 6f03a6549195
6f03a6549195

# 동일 컨테이너 재실행
sudo docker run -p 8080:80 --name web -d nginx
docker: Error response from daemon: Conflict. The container name "/web" is already in use by container "6f03a6549195ce12282bec6e697f801e1850ff90e948645949f7ba954daef5a6". You have to remove (or rename) that container to be able to reuse that name.
See 'docker run --help'.

# 실행 컨테이너 목록
sudo docker ps -a

# 모든 컨테이너 목록(실행 이력 포함)
sudo docker ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                      PORTS               NAMES
6f03a6549195        nginx               "/docker-entrypoint.▒€▒"   8 minutes ago       Exited (0) 20 seconds ago                       web

# 실행 이력 삭제
sudo docker rm 6f03a6549195
6f03a6549195

# 재실행 가능
sudo docker run -p 8080:80 --name web -d nginx
0086054e6d19635acf9354ec8a67b664a7648ff633d5257c7fcdf1441800fc9b

# 종료시 실행 이력 없이 종료되도록 --rm 추가
sudo docker run --rm -p 8080:80 --name web -d nginx
```

### (3) Docker 기반 CB-Spider 실행

```
sudo docker run --rm -p 1024:1024 -p 2048:2048  -v /tmp/meta_db:/root/go/src/github.com/cloud-barista/cb-spider/meta_db --name cb-spider cloudbaristahub/cb-spider:v0.2.0-20200820
(cloudbaristahub:DockerHub임시계정)
 성공시 다음 메시지 출력함
```
```
[CB-Spider:Cloud Info Management Framework]

   Initiating REST API Server....__^..^__....


▒눊 http server started on [::]:1024

[CB-Spider:Cloud Info Management Framework]
   Initiating GRPC API Server....__^..^__....

 => grpc server started on :2048
```

#### AdminWeb: Browser로 접근
```
http://spider-server-publicip:1024/spider/adminweb
```


#### CLI: 다른 Terminal에서 실행

```
cd $CBSPIDER_ROOT/interface/cli/spider
./spider os list

cb-spider/interface/cli/spider# ./spider vpc --cname "aws-aws-(ohio)us-east-2-connection-config-01" list
vpc:
- IId:
    NameId: vpc-01
    SystemId: vpc-0c67242eab0b4f4ab
  IPv4_CIDR: 192.168.0.0/16
  SubnetInfoList:
  - IId:
      NameId: subnet-01
      SystemId: subnet-0061dc81ef4aa0388
    IPv4_CIDR: 192.168.1.0/24
    KeyValueList:
    - Key: VpcId
      Value: vpc-0c67242eab0b4f4ab
    - Key: MapPublicIpOnLaunch
      Value: "false"
    - Key: AvailableIpAddressCount
      Value: "251"
    - Key: AvailabilityZone
      Value: us-east-2a
    - Key: Status
      Value: available
  KeyValueList: null
```

## [Source 기반 설치/실행]

### go 설치
참고: https://gist.github.com/powerkimhub/9a722304a14c5af8b3dff56ab064fb43

### gcc 설치
sudo apt install -y  gcc

### cb-spider 소스 다운로드
```
go get -u -v github.com/cloud-barista/cb-spider
cd $GOPATH/src/github.com/cloud-barista/cb-spider
go get -u -v -t  ./... 
```

- 다음 메시지는 무시
package _/home/byoungseob/gosrc/src/github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/gcp/main/old/conf: cannot find package "_/home/byoungseob/gosrc/src/github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/gcp/main/old/conf" in any of:
        /home/byoungseob/go/src/_/home/byoungseob/gosrc/src/github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/gcp/main/old/conf (from $GOROOT)
        /home/byoungseob/gosrc/src/_/home/byoungseob/gosrc/src/github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/gcp/main/old/conf (from $GOPATH)


### 2020.08.19.현재 수정 필요한 package dependency
```
==================
	성공환경 활용 pkg path에서 실행> git show
	설치환경 pkg path에서 실행> git checkout 7dc0a2d6ddce55257ea8851e23b4fb9ef44fd4a0
==================
```

```
cd $GOPATH/src/github.com/Azure/go-autorest;
git checkout tags/autorest/azure/auth/v0.4.2; 
cd $GOPATH/src/github.com/Azure/azure-sdk-for-go;
git checkout tags/v37.2.0;
cd $GOPATH/src/github.com/docker/docker/vendor/github.com/containerd/containerd/errdefs;
git checkout dd16f2f21984338565ec8a8a896e7491d8948d93;
cd $GOPATH/src/github.com/docker/docker/client;
git checkout f6163d3f7a10c5d01a92bc8b86e204d784b2f6d6

rm -rf $GOPATH/src/github.com/docker/docker/vendor/github.com/docker/go-connections;
rm -rf $GOPATH/src/github.com/docker/docker/vendor/github.com/pkg;
rm -rf $GOPATH/src/go.etcd.io/etcd/vendor/golang.org/x/net/trace;
```

# vi ~/.bashrc
alias spider='cd /home/byoungseob/gosrc/src/github.com/cloud-barista/cb-spider'
source /home/byoungseob/gosrc/src/github.com/cloud-barista/cb-spider/setup.env



## 수고하셨습니다~
