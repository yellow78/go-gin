# Dockerfile
FROM golang:1.23.4-alpine

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 並下載依賴
COPY go.mod go.sum ./
RUN go mod download

# 複製專案檔案
COPY . .

# 編譯 Go 應用程式
RUN go build -o main ./cmd/api

# 暴露應用程式埠
EXPOSE 8080

# 執行應用程式
CMD ["./main"]