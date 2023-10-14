#デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.18.2-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -trimpath -ldflags "-w -s" -o app ./cmd/.

#-----------------------------------------------
#デプロイ用コンアテナ
FROM centos:centos7 as deploy

RUN yum -y update

EXPOSE "8080"

COPY --from=deploy-builder /app/app .

CMD ["./app"]

#-----------------------------------------------
#ローカル開発環境で利用するホットリロード環境
FROM golang:1.21 as dev

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]