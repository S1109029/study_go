FROM golang:1.22.0-alpine

WORKDIR /app
COPY ./src .

RUN apk upgrade --update && \
    apk --no-cache add git && \
    apk upgrade git

RUN go mod init my-project

RUN go mod tidy

RUN go mod download
# watchexec のインストール
RUN apk add --no-cache -X http://dl-cdn.alpinelinux.org/alpine/edge/community watchexec

# ホットリロード用のコマンドを設定
CMD ["watchexec", "-w", ".", "--force-poll", "100", "-r", "go", "run", "main.go"]
