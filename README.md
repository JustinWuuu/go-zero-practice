# go-zero-practice

一個使用 [go-zero](https://github.com/zeromicro/go-zero) 練習微服務架構的小專案，示範 API Gateway 呼叫多個 RPC 服務的典型用法。

## 架構

```
user-api (HTTP, :8888)
  ├─ 呼叫 user-rpc  (gRPC, :8080) 取得 / 建立使用者
  └─ 呼叫 order-rpc (gRPC, :8081) 取得訂單資訊

order-rpc (gRPC, :8081)
  └─ 呼叫 user-rpc  (gRPC, :8080)（示範 RPC 之間互相呼叫）
```

三個服務都透過 etcd（`127.0.0.1:2379`）互相註冊與發現。

## 各服務說明

### user-api
對外的 HTTP Gateway，路由定義在 [`user-api/user.api`](user-api/user.api)：

| Method | Path       | 說明               |
|--------|------------|--------------------|
| GET    | /user/:id  | 查詢使用者資訊     |
| POST   | /user      | 建立使用者         |
| GET    | /order/:id | 查詢訂單資訊       |

### user-rpc
gRPC 服務，提供 `getUserInfo`、`createUser`（定義於 [`user-rpc/user.proto`](user-rpc/user.proto)）。

### order-rpc
gRPC 服務，提供 `getOrderInfo`（定義於 [`order-rpc/order.proto`](order-rpc/order.proto)）。

## 目錄結構

每個服務都是 go-zero 用 `goctl` 產生的標準結構：

```
<service>/
├── etc/                # 設定檔 (yaml)
├── internal/
│   ├── config/         # 設定結構
│   ├── handler/        # (僅 api) HTTP handler
│   ├── logic/          # 商業邏輯
│   ├── server/         # (僅 rpc) gRPC server
│   └── svc/            # 依賴注入的 ServiceContext
├── *.proto / *.api     # 服務定義
└── *.go                # 主程式入口
```

## 執行方式

需要先啟動本機 etcd（預設連線 `127.0.0.1:2379`），再依序啟動三個服務：

```bash
# etcd（需自行安裝，或用 docker 啟動）
etcd

# 分別在各自目錄下執行
cd user-rpc  && go run user.go -f etc/user.yaml
cd order-rpc && go run order.go -f etc/order.yaml
cd user-api  && go run user.go -f etc/user-api.yaml
```

服務啟動後可透過 `http://127.0.0.1:8888` 呼叫 user-api。

## 練習重點

- go-zero 的 `goctl` 專案生成流程（`.api` / `.proto` → handler / logic / server 骨架）
- API Gateway 與多個 RPC 服務之間的呼叫關係
- 服務間透過 etcd 做服務註冊與發現
