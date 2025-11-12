# GRPC-GO-DEMO

一个完整的 gRPC Go 示例项目，包含服务器、客户端和 Web 前端调用。

## 🌟 功能特性

- ✅ gRPC 服务器实现（Calculator 加法服务）
- ✅ Go 客户端示例
- ✅ **Web 前端调用 gRPC（通过 HTTP Gateway）**
- ✅ 美观的用户界面
- ✅ 完善的错误处理
- ✅ CORS 支持

## 🚀 快速开始

### 1. 启动 gRPC 服务器

```bash
# Windows
启动服务器.bat

# Linux/Mac
go run server/main.go
```

### 2. 启动 HTTP Gateway

```bash
# Windows
web\启动网关.bat

# Linux/Mac
go run web/gateway.go
```

### 3. 访问 Web 界面

打开浏览器访问：`http://localhost:8080/test.html`

## 📁 项目结构

```
grpc-golang-demo/
├── proto/                      # Protocol Buffers 定义
│   └── calculator.proto        # Calculator 服务定义
├── gen/proto/                  # 生成的 protobuf 代码
├── server/                     # gRPC 服务器
│   └── main.go                 # 服务器实现
├── client/                     # gRPC Go 客户端
│   └── main.go                 # 客户端示例
├── web/                        # Web 前端
│   ├── gateway.go              # HTTP Gateway（桥接浏览器和gRPC）
│   ├── test.html               # 前端页面
│   └── 启动网关.bat            # Windows 启动脚本
├── 启动服务器.bat              # Windows 启动脚本
├── 快速启动指南.md             # 快速入门指南
└── 运行说明.md                 # 详细技术文档
```

## 🏗️ 架构说明

```
浏览器 (test.html)
    ↓ HTTP POST (JSON)
HTTP Gateway (gateway.go:8080)
    ↓ gRPC
gRPC Server (server/main.go:50051)
```

## 📖 文档

- [快速启动指南.md](./快速启动指南.md) - 3步快速上手
- [运行说明.md](./运行说明.md) - 详细的技术文档和架构说明

## 🛠️ 技术栈

- **Go 1.24+** - 编程语言
- **gRPC** - RPC 框架
- **Protocol Buffers** - 数据序列化
- **HTML/CSS/JavaScript** - 前端技术

## 💪 gRPC 的优势

### 📊 性能优势
- **更高效的序列化**：Protocol Buffers 比 JSON 更紧凑，序列化/反序列化速度快 **5-10 倍**
- **二进制传输**：相比 JSON 文本传输，消息大小减少 **50-80%**
- **多路复用（HTTP/2）**：单一连接支持多个并发请求，减少连接开销

### 🚀 开发体验
- **强类型定义**：Proto 文件定义契约，避免字段类型错误
- **自动代码生成**：支持多种编程语言，减少重复代码
- **向后兼容**：方便 API 版本升级而不破坏现有客户端
- **服务通信标准化**：统一的 RPC 调用方式

### 🔌 功能完善
- **流式通信**：支持服务端流、客户端流、双向流
- **内置错误处理**：标准化的错误码和信息
- **超时控制**：原生支持请求超时设置
- **认证和加密**：内置 TLS/SSL 支持

## 🎯 解决的场景问题

### 问题1️⃣：**微服务间通信效率低**
- ❌ 传统 REST API：每个请求都是独立连接，序列化效率低
- ✅ gRPC 方案：HTTP/2 多路复用 + 高效二进制协议，减少延迟和带宽占用

### 问题2️⃣：**API 版本管理困难**
- ❌ REST API：新增字段需要版本迭代（v1/v2/v3...），客户端兼容性复杂
- ✅ gRPC 方案：Proto 字段自动编号，新增字段不影响旧版本，自动向后兼容

### 问题3️⃣：**服务间通信数据量大**
- ❌ JSON：冗余字段名和类型信息，消息体膨胀
- ✅ Protocol Buffers：紧凑的二进制格式，带宽节省 50-80%

### 问题4️⃣：**实时数据推送困难**
- ❌ REST API：只能 Polling，效率低，服务器压力大
- ✅ gRPC Streaming：原生支持服务端推送，实时性强

### 问题5️⃣：**跨语言调用复杂**
- ❌ 自定义协议：每种语言需要独立实现序列化/反序列化
- ✅ gRPC：官方支持 10+ 语言，一次定义多语言使用

## 📈 性能对比

| 指标 | REST (JSON) | gRPC | 提升 |
|------|-----------|------|------|
| 消息大小 | 100KB | 20KB | 80% 减少 |
| 序列化速度 | 10ms | 2ms | 5x 快速 |
| HTTP 连接 | 新建连接 | HTTP/2 复用 | 减少开销 |
| 带宽占用 | 基准值 | 20-30% | 节省 70% |

## 🎯 使用场景

这个项目展示了如何在 Web 应用中调用 gRPC 服务，适合：

- 学习 gRPC 基础
- 了解如何在浏览器中调用 gRPC
- 构建微服务架构的参考
- 前后端分离项目的参考

## 📝 API 示例

### HTTP Gateway API

**请求：**
```bash
POST http://localhost:8080/api/add
Content-Type: application/json

{
  "a": 10,
  "b": 20
}
```

**响应：**
```json
{
  "result": 30
}
```

## 🔧 开发工具

### [Buf Tools](https://github.com/bufbuild/buf)

By using the Buf plugin to manage grpc gateway,
it can efficiently configure various plugins, saving us the trouble of manually editing sh

### 重新生成 Proto 代码

```bash
# 如果修改了 proto 文件，需要重新生成代码
./build-proto.sh
```

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

---

⭐ 如果这个项目对你有帮助，请给个 Star！
