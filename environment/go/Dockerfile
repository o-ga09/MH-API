#API用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.24-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -trimpath -ldflags "-w -s" -o main ./app/cmd/api/main.go

#バッチ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.24-bullseye as deploy-batch-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -trimpath -ldflags "-w -s" -o main ./app/cmd/batch/main.go

#-----------------------------------------------
#API デプロイ用コンアテナ
FROM centos:centos7 as deploy-api

RUN yum -y update

EXPOSE "8080"

COPY --from=deploy-builder /app/main .

CMD ["./main"]

#-----------------------------------------------
#バッチ デプロイ用コンアテナ
FROM centos:centos7 as deploy-batch

RUN yum -y update

EXPOSE "8080"

COPY --from=deploy-batch-builder /app/main .

CMD ["./main"]

#-----------------------------------------------
#ローカル開発環境で利用するホットリロード環境
FROM golang:1.24 as dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go install github.com/air-verse/air@latest
CMD ["air"]