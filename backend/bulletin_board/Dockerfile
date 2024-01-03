# Use an official Golang runtime as a parent image
FROM golang:latest  AS builder

# 빌드 이미지

# 작업 디렉토리 설정
WORKDIR /app

# Go 모듈 파일 복사
COPY go.mod go.sum ./

# 의존성 다운로드, Beego 및 Bee 명령어 툴 포함
RUN go mod download
RUN go get github.com/beego/beego/v2@latest
RUN go get github.com/beego/bee/v2@latest

# 소스 코드 복사
COPY . .

# Beego 애플리케이션 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 최종 실행 이미지
FROM alpine:latest  
WORKDIR /root/

# 빌드된 실행 파일 복사
COPY --from=builder /app/main .

# 포트 8080 열기 (Beego 기본 포트)
EXPOSE 8080

# 실행 명령
CMD ["./main"]