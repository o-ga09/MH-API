#API用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.25.4-alpine3.22 as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -trimpath -ldflags "-w -s" -o main ./cmd/api/main.go

#バッチ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.25.4-alpine3.22 as deploy-batch-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -trimpath -ldflags "-w -s" -o main ./cmd/batch/main.go

#MCP用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.25.4-alpine3.22 as deploy-mcp-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -trimpath -ldflags "-w -s" -o main ./cmd/mcp/main.go

#Agent用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.25.4-alpine3.22 as deploy-agent-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -trimpath -ldflags "-w -s" -o main ./cmd/agent/main.go

#-----------------------------------------------
#API デプロイ用コンアテナ
FROM ubuntu:22.04 as deploy-api

RUN apt update
RUN apt-get install -y ca-certificates openssl

EXPOSE "8080"

COPY --from=deploy-builder /app/main .

CMD ["./main"]

#-----------------------------------------------
#バッチ デプロイ用コンアテナ
FROM ubuntu:22.04 as deploy-batch

RUN apt update
RUN apt-get install -y ca-certificates openssl

EXPOSE "8080"

COPY --from=deploy-batch-builder /app/main .

CMD ["./main"]

#-----------------------------------------------
#MCP デプロイ用コンテナ
FROM ubuntu:22.04 as deploy-mcp

RUN apt update
RUN apt-get install -y ca-certificates openssl

EXPOSE "8080"

COPY --from=deploy-mcp-builder /app/main .

CMD ["./main"]

#-----------------------------------------------
#Agent デプロイ用コンテナ
FROM ubuntu:22.04 as deploy-agent

RUN apt update
RUN apt-get install -y ca-certificates openssl

EXPOSE "8080"

COPY --from=deploy-agent-builder /app/main .

CMD ["./main"]

#-----------------------------------------------
#ローカル開発環境で利用するホットリロード環境
FROM golang:1.25 as dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go install github.com/air-verse/air@latest
CMD ["air"]
#-----------------------------------------------
#ローカル開発環境で利用するホットリロード環境
FROM golang:1.25 as dev-mcp

WORKDIR /app

COPY go.mod go.sum ./

RUN go install github.com/air-verse/air@latest
CMD ["air"]
#-----------------------------------------------
#Agent ローカル開発環境で利用するホットリロード環境
FROM golang:1.25.4-alpine3.22 as dev-agent

WORKDIR /app

COPY go.mod go.sum ./

RUN go install github.com/air-verse/air@latest
CMD ["air"]
