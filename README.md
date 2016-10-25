# wigo
Wiki Engine written in Go

## Quick Start

1. `go get` 명령어를 이용해 저장소를 `GOHOME`으로 가져옵니다.
```
go get github.com/GDG-SSU/wigo
```
2. 데이터베이스 연결 설정을 담은 쉘 환경 변수(Environment Variables)를 선언합니다.
```
export DB_USER="데이터베이스 접근 계정의 사용자 이름"
export DB_PASSWORD="데이터베이스 접근 계정의 비밀번호"
export DB_NAME="데이터베이스 이름"
```
3. `revel` 명령어를 이용해 애플리케이션을 실행합니다.
```
revel run github.com/GDG-SSU/wigo
```

## 패키지 설치

1. `curl https://glide.sh/get | sh` 명령어를 실행하여 glide를 설치한다.
2. 프로젝트 디렉토리(glide.yaml가 있는 디렉토리)에서 `glide install`을 실행한다.
```
glide install
```
