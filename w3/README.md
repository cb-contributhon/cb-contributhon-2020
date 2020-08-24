# [개    요]
  - Cloud-Barista 활용 실습 기간: 3주 ~ 4주 예상
  - Cloud-Barista 실습 계획: CB-Spider => CB-Tumblebug
  - 3주: CB-Spider 설치/실행/활용

***

# [환경 준비]
## 1. 사전 준비
```
       (1) GCP-VM1: 아래 VM1 설정 참조 후 신규 생성
           - 목적: CB-Spider 설치/가동, CLI 시험
           - VM1 로그인 위한 접속 Key 및 Terminal
       (2) AWS 계정의 Credential 정보(지난 미션시 생성한거 활용)
            - 목적: Spider에 입력 후 Spider로 AWS에 명령 요청시 사용
       (3) Web Browser : 노트북/데스크탑 등에 설치된 아무 브라우저 가능
            - 목적: CB-Spider AdminWeb 접속
```

## 2. VM1 설정
```
  - OS:  Ubuntu 18.04
  - 사양: 2 vCPU, 4GB, 20GB Disk 이상
  - 방화벽(보안그룹):
    - CB-Spider 서버용: 1024(Rest Server), 2048(gRPC Server) 포트 개발
    - Nginx Test용: 8080 포트 개방

  - VM1 로그인 후
  - sudo apt update -y  // 옵션, 필요시 실행
```

***

# [CB-Spider 서버 가동]
## 1. Docker기반 가동 방법
## 2. 소스기반 가동 방법

## 1. Docker 기반 가동
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

### (3) Docker 기반 CB-Spider 서버 실행

```
sudo docker run --rm -p 1024:1024 -p 2048:2048  -v /tmp/meta_db:/root/go/src/github.com/cloud-barista/cb-spider/meta_db --name cb-spider cloudbaristaorg/cb-spider:v0.2.0-20200821 
```

성공시 다음 메시지 출력함

```
[CB-Spider:Cloud Info Management Framework]

   Initiating REST API Server....__^..^__....


 => http server started on [::]:1024

[CB-Spider:Cloud Info Management Framework]
   Initiating GRPC API Server....__^..^__....

 => grpc server started on :2048
```



## 2. 소스기반 가동

### (1) go 설치
참고: https://gist.github.com/powerkimhub/9a722304a14c5af8b3dff56ab064fb43

### (2) gcc 설치
sudo apt install -y  gcc

### (3) cb-spider 소스 다운로드
```
go get -u -v github.com/cloud-barista/cb-spider
cd $GOPATH/src/github.com/cloud-barista/cb-spider
go get -u -v -t  ./... 
```

- 다음 메시지는 무시
package _/home/byoungseob/gosrc/src/github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/gcp/main/old/conf: cannot find package "_/home/byoungseob/gosrc/src/github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/gcp/main/old/conf" in any of:
        /home/byoungseob/go/src/_/home/byoungseob/gosrc/src/github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/gcp/main/old/conf (from $GOROOT)
        /home/byoungseob/gosrc/src/_/home/byoungseob/gosrc/src/github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/gcp/main/old/conf (from $GOPATH)


### (4) Spider 개발시 활용한 패키지 버전 수정 필요(2020.08.19.현재)
```
================== 이 부분은 참고만
	성공환경 활용 pkg path에서 실행> git show   // 개발시 활용한 패키지의 commit 버전을 복사
	설치환경 pkg path에서 실행> git checkout 7dc0a2d6ddce55257ea8851e23b4fb9ef44fd4a0
==================
```
  - 다음 블럭을 복사 후 VM1에서 실행
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

### (5) vi ~/.bashrc  // 끝에 다음 2줄을 추가하고 source .bashrc 또는 나갔다 들옴.
```
alias spider='cd $HOME/gosrc/src/github.com/cloud-barista/cb-spider'
source $HOME/gosrc/src/github.com/cloud-barista/cb-spider/setup.env
```

### (6) 소스 기반 CB-Spider 서버 가동

#### 1. Stop cb-spider continaer
Docker 컨테이너로 서버 가동 중이면, Spider 서버 컨테이너 종료
```
sudo docker stop cb-spider
```
#### 2. Change directory to api-runtime
api-runtime 위치로 이동
```
cd $CBSPIDER_ROOT/api-runtime
```
#### 3. Run cb-spider server
cb-spider 서버 가동
```
go run *.go
```

***

# CB-Spider 간단 활용 방법
- Docker 기반 서버 가동 또는 소스 기반 서버 가동 후에 실행

## 1. AdminWeb 이용한 활용: Browser로 접근

  - 접속 링크: http://spider-server-publicip:1024/spider/adminweb
  - 가이드 참고: https://drive.google.com/file/d/13x8Amdsoq3RabZhNHD1VC933LeYYLX4z/view?usp=sharing
  - AWS 대상으로 시험 추천(검증이 많이 됨, 아직 불안할 수도 있음^^)
    - CB-Spider로 생성한 자원은 가급적 CB-Spider로 삭제
  - CB-Spider 메타 정보 초기화 방법
    - spider 서버 종료
    - rm -rf $CBSPIDER_ROOT/meta_db/dat
    - spider 서버 가동

## 2. CLI 이용한 활용: 다른 Terminal에서 실행
  - 본 가이드의 CB-Spider 소스 기반 설치 완료 후 CLI 활용    

```
cd $CBSPIDER_ROOT/interface/cli/spider
go build spider.go  // spider cli build
./spider os list    // 현재 Spider가 제공하는 CloudOS 목록 제공

./spider vpc --cname "aws-aws-(ohio)us-east-2-connection-config-01" list    // cname이 생성한거와 동일한지 확인 필요
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

## 수고하셨습니다~

