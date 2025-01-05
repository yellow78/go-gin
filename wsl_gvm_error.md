# GVM 和 Go Modules 介紹

## 1. GVM（Go 版本管理器）

### 介紹
GVM 是一個開源工具，用於管理多個 Go 版本。它允許開發者安裝、卸載和切換不同的 Go 版本，方便在不同專案中使用不同的 Go 版本。

### 安裝
可以從 GVM 的 GitHub 倉庫下載並安裝。安裝後，使用 `gvm listgvm gos` 命令查看已安裝的 Go 版本。例如，安裝 Go 1.4 版本時，可能需要使用 `-B` 參數來避免編譯錯誤。

### 使用
使用 `gvm use go1.4` 命令切換到 Go 1.4 版本。切換後，GVM 會自動設置相應的環境變數，方便開發者使用。

## 2. Go Modules

### 介紹
Go Modules 是 Go 1.11 引入的套件管理工具，允許開發者在不依賴 GOPATH 的情況下管理專案的依賴。它使得套件管理變得更加靈活和高效。

### 啟用
在 Go 1.11 或更高版本中，使用 `GO111MODULE=on` 環境變數來啟用 Go Modules。在專案目錄中，使用 `go mod init` 命令初始化 Go Modules。

### 使用
在專案中，使用 `go mod` 命令來管理依賴，例如 `go mod tidy` 來整理依賴，`go mod vendor` 來創建 vendor 目錄等。

## 3. GVM 和 Go Modules 的結合使用

### 安裝 Go 版本
使用 GVM 安裝所需的 Go 版本，例如 Go 1.12。安裝後，使用 `gvm use go1.12` 切換到該版本。

### 初始化 Go Modules
在專案目錄中，使用 `go mod init` 命令初始化 Go Modules。如果遇到錯誤，可能需要指定模組路徑，例如 `go mod init /go/mod-demo`。

### 管理依賴
使用 `go mod` 命令來管理專案的依賴，例如 `go mod tidy` 來整理依賴，`go mod vendor` 來創建 vendor 目錄等。

## 4. 常見問題

### 編譯錯誤
在安裝 Go 版本時，可能會遇到編譯錯誤。例如，安裝 Go 1.4 版本時，可能需要使用 `-B` 參數來避免編譯錯誤。

### 環境變數
切換 Go 版本後，GVM 會自動設置相應的環境變數，但有時可能需要手動設置。可以使用 `gvm listgvm gos` 命令查看當前使用的 Go 版本。

## 總結
GVM 和 Go Modules 是管理 Go 環境和套件的有力工具。GVM 允許開發者管理多個 Go 版本，而 Go Modules 則提供了靈活的套件管理方式。結合使用這兩者，可以有效地管理 Go 專案的環境和依賴。

原文連結：[GVM 和 Go Mod](https://medium.com/golang-%E7%AD%86%E8%A8%98/gvm-go-mod-492a54c15c41)