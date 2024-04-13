#API用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.22-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -trimpath -ldflags "-w -s" -o app ./app/cmd/api/.

#バッチ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.22-bullseye as deploy-batch-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -trimpath -ldflags "-w -s" -o app ./app/cmd/batch/.

#-----------------------------------------------
#API デプロイ用コンアテナ
FROM centos:centos7 as deploy-api

RUN yum -y update

EXPOSE "8080"

COPY --from=deploy-builder /app/app .

CMD ["./app"]

#-----------------------------------------------
#バッチ デプロイ用コンアテナ
FROM centos:centos7 as deploy-batch

RUN yum -y update

EXPOSE "8080"

COPY --from=deploy-builder /app/app .

CMD ["./app"]

#-----------------------------------------------
#ローカル開発環境で利用するホットリロード環境
FROM golang:1.22 as dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]