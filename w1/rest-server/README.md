## [개    요]
  - CB-Contributhon 학습용 예제: Echo 기반 REST 서비스 시험

***

## [목    차]

1. [시나 리오](#시나-리오)
2. [실행 환경](#실행-환경)
3. [서버 가동](#서버-가동)
4. [호출 시험](#호출-시험)

***

## [시나 리오]
  - VM1: Rest 서버 가동
  - VM2: Rest 호출 시험
    - VM2에서 VM1의 REST API 호출

## [실행 환경]

  - 공통: 리눅스 (검증시험:Ubuntu 18.04) 
  - VM1: 이전 단계 생성한 VM, Go 개발환경 설치 완료
  - VM2: 본 시험 과정에서 생성, 생성 방법 아래 내용 참고

## [서버 가동]
  - VM1 환경

### (1) Web Service 개발관련
  - echo framework 활용한 웹서버 개발
  - REST API 제공
  - 소스 다운로드 후 시험 우선 수행
  - 소스 수정 및 시험 반복을 통한 분석
  - 소스 설명: 차주 미팅 
  
### (2) 실행 환경 설정
  - VM1 보안그룹(SecurityGroup) 설정
    - 8080 port inbound 추가
    - 참고
      ```
      포트	프로토콜	소스	launch-wizard-6
      8080	tcp	0.0.0.0/0	
      ```
      
  - 기존 VM 설정 중 Go 환경변수 변경: ~$GOROOT 설정 제거~
      - mkdir $HOME/gosrc
      - .bashrc 추가했었던 GO 관련 설정을 다음 줄로 교체
        - export GOPATH=$HOME/gosrc;GOROOT=$HOME/go;export PATH=$PATH:$GOROOT/bin;
      - 적용 위해서 logout 후 다시 login
      - 참고: Go 환경 설정 방법(최신 gist): https://gist.github.com/powerkimhub/d1d6b260228746e14151685bbf2cdf03

### (3) REST 서버 소스 다운로드
  - go get 이용한 저장소 다운로드
	  - $ go get -u github.com/cb-contributhon/cb-contributhon-2020
  - 다운로드 받은 저장소 위치로 이동
	  - $ cd $GOPATH/src/github.com/cb-contributhon/cb-contributhon-2020
  - 환경변수 설정: Log 설정 파일 위치 지정
    - $ export CBLOG_ROOT=/home/ubuntu/go/src/github.com/cb-contributhon/cb-contributhon-2020/w1/rest-server
    - 반복 사용 위해서는 .bashrc 끝에 추가
      - 적용 위해서 $ source .bashrc 실행


### (4) REST 서버 가동
  - 서비스 가동 시도
	- $ go run myserver.go
  - 필요한 패키지 부재로 아래와 같은 오류 발생 

      ```
      ubuntu@ip-172-31-40-73:~/w1/rest-server$ go run myserver.go
      myserver.go:14:2: cannot find package "github.com/labstack/echo" in any of:
        /home/ubuntu/go/src/github.com/labstack/echo (from $GOROOT)
        /home/ubuntu/go/src/src/github.com/labstack/echo (from $GOPATH)
      myserver.go:16:2: cannot find package "github.com/labstack/echo/middleware" in any of:
        /home/ubuntu/go/src/github.com/labstack/echo/middleware (from $GOROOT)
        /home/ubuntu/go/src/src/github.com/labstack/echo/middleware (from $GOPATH)
      myserver.go:11:2: cannot find package "github.com/sirupsen/logrus" in any of:
        /home/ubuntu/go/src/github.com/sirupsen/logrus (from $GOROOT)
        /home/ubuntu/go/src/src/github.com/sirupsen/logrus (from $GOPATH)
      ```

  - 필요 package import
    - 다음처럼 실행: 
```    
    - $ go get `go run myserver.go 2>&1 |grep cannot |awk '{print $5}' | sed 's/"//g'`
      - 출력무시: go get: no install location for directory /home/ubuntu/w1/rest-server outside GOPATH
```      
  - 서비스 가동 재시도
	- $ go run myserver.go
  - 다음 메시지 출력이면 성공: 안되면 slack 문의

      ```
      [CB-Contributhon:Test REST Framework]

         Initiating REST Server....__^..^__....
      ```

## [호출 시험]
  - VM2 환경
  
### (1) GCP VM2 생성
  - 필요시 김윤곤멘토 자료 참고
  
### (2) VM2 로그인
  - 예시: $ ssh -i private_key user@34.72.180.11
    - key: VM2에 설정한 private key
    - user: VM2 접근 가능한 user
    - IP: 소유 VM2 Public IP
  
### (3) REST API 호출 시험
  - $ curl -sX GET http://3.19.54.10:8080/2020/myinfo |json_pp
    - IP: 소유 VM1 Public IP

## 수고하셨습니다~
