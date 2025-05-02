# Go 語言 DDD 架構的資料夾結構

## 1. 建議的 DDD 資料夾架構

```
/project-root
│
├── /cmd                 # 主應用程式入口點（多個可執行檔入口）
│   └── /appname         # 每個應用程式都有自己的 main.go
│
├── /internal            # 內部使用的核心邏輯
│   ├── /domain          # 領域層：核心商業邏輯
│   │   ├── /model       # 領域模型和聚合根
│   │   │   ├── /entity          # 實體（Entity）
│   │   │   ├── /valueobject     # 值物件（Value Object）
│   │   │   └── /aggregate       # 聚合（Aggregate）
│   │   ├── /repository  # 定義存取接口（介面）
│   │   └── /service     # 領域服務
│   │
│   ├── /application     # 應用層：用戶案例和應用服務
│   │   ├── /dto         # 資料傳輸物件
│   │   └── /service     # 應用服務
│   │
│   ├── /infrastructure  # 基礎設施層
│   │   ├── /persistence # 實作存取層（資料庫）
│   │   ├── /messaging   # 訊息處理（如 Kafka、RabbitMQ）
│   │   └── /external    # 外部 API 調用
│   │
│   └── /interfaces      # 接口層：控制器和處理器
│       ├── /http        # HTTP 控制器
│       └── /grpc        # gRPC 控制器
│
├── /pkg                 # 可重用的第三方或工具程式碼
│   ├── /logger          # 日誌封裝
│   └── /config          # 配置管理
│
├── /migrations          # 資料庫 schema 管理
│
├── go.mod               # Go 模組定義
├── go.sum
└── README.md            # 專案說明
```

---

## 2. 各層職責說明

### **cmd**
存放應用程式的入口點，每個子資料夾包含一個 `main.go`，用於啟動不同的應用。

### **internal/domain**
核心商業邏輯，定義：
- **實體（Entity）**
- **值物件（Value Object）**
- **聚合（Aggregate）**
- **存取庫接口（Repository interfaces）**

#### **model/entity**
- 定義實體類型，通常包含唯一識別符（ID）和屬性。
- **範例**：  
  `user.go`
  ```go
  package entity

  type User struct {
      ID       string
      Name     string
      Email    string
      Password string
  }
  ```

#### **model/valueobject**
- 值物件是不可變的、無身份的，並且是原子值或多個值的組合。
- **範例**：  
  `address.go`
  ```go
  package valueobject

  type Address struct {
      Street string
      City   string
      Zip    string
  }
  ```

#### **model/aggregate**
- 聚合是包含多個實體和值物件的邏輯組合，具有一個根實體（Aggregate Root）。
- **範例**：  
  `order.go`
  ```go
  package aggregate

  import (
      "project/internal/domain/model/entity"
      "project/internal/domain/model/valueobject"
  )

  type Order struct {
      ID       string
      Customer entity.User
      Items    []valueobject.OrderItem
  }
  ```

---

### **internal/application**
應用服務（Application Service），包含具體的：
- **用戶案例（Use Cases）**
- **資料傳輸物件（DTOs）**

### **internal/infrastructure**
提供基礎設施的實現，如：
- 資料庫存取
- 第三方 API 調用
- 消息傳遞（如 Kafka、RabbitMQ）

### **internal/interfaces**
提供用於外部世界（如 HTTP、gRPC）的控制器和處理器。

### **pkg**
共用程式碼，如：
- 配置管理
- 日誌工具

---

## 3. 說明

1. Go 語言的 `internal` 目錄遵循包封裝規則，用於限制套件的可見性。
2. `pkg` 用於放置專案中可被其他專案或模組共用的工具函式和支援代碼。

---

這種設計能夠有效地將 **核心商業邏輯** 封裝在一個領域層，並且易於擴展和維護。
