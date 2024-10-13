# ベースイメージとしてGolangの公式イメージを使用
FROM golang:1.22-alpine

# 作業ディレクトリを設定
WORKDIR /app

# 必要なパッケージをインストール
RUN apk add --no-cache git

# ソースコードをコピー
COPY . .

# 依存関係をダウンロードしてビルド
RUN go mod download
RUN go build -o /app/migrate cmd/migrate/main.go

# 実行コマンドを設定
CMD ["/app/migrate", "-migration-dir=migrations", "-seed-dir=seeds"]
