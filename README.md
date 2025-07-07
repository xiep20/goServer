# Gin Demo 项目说明

## 项目简介

本项目是一个基于 [Gin](https://github.com/gin-gonic/gin) 框架的 Go Web 应用，集成了表单处理、文件上传、静态资源服务、模板渲染、重定向等常见功能，可以部署前端页面的静态资源。

尝试使用go做一个简单服务端的测试代码。

## 主要功能

- **欢迎接口**：简单的 Hello World 及参数演示。
- **表单处理**：支持普通表单提交、单文件上传、多文件上传。
- **模板渲染**：支持 HTML 和 Go 模板（`.tmpl`）渲染。
- **静态资源服务**：可访问 `static/` 目录下的前端资源（如 Cesium）。
- **重定向**：支持接口重定向到外部网站。
- **目录浏览**：可通过接口浏览服务器目录或指定文件。

## 目录结构

```
.
├── main.go                // 项目主程序，Gin 路由与业务逻辑
├── go.mod, go.sum         // Go 依赖管理文件
├── static/                // 前端静态资源（含 Cesium 三维地球库）
├── templates/             // HTML 页面与 Go 模板
├── upload/                // 上传文件存储目录
├── favicon.ico            // 网站图标
└── gin-demo/              // 预留子模块目录（当前为空）
```

## 依赖环境

- Go 1.19 及以上
- Gin v1.9.1

依赖包详见 `go.mod`。

## 启动方式

1. **安装依赖**

   ```bash
   go mod tidy
   ```

2. **运行项目**

   ```bash
   go run main.go
   ```

3. **访问服务**

   - 默认监听端口：`http://127.0.0.1:8000`
   - 示例接口：
     - 欢迎页：`GET /come/`
     - 表单提交：`POST /form/formSubmit`
     - 单文件上传：`POST /form/upload`
     - 多文件上传：`POST /form/multiUpload`
     - 模板页：`GET /views/html/login`、`GET /views/tmpl/index`
     - 静态资源：`/static/Cesium/Cesium.js` 等
     - 目录浏览：`/showDir`
     - 重定向：`/redirect/bd`

## 典型页面与接口示例

- `templates/login.html`：登录表单页面，演示表单提交。
- `templates/file.html`：单文件上传页面。
- `templates/multiFile.html`：多文件上传页面。
- `templates/map.html`、`templates/map2.html`：Cesium 三维地球演示页面。

## 备注

- 上传的文件会存储在 `upload/` 目录下。
- Cesium 相关静态资源位于 `static/Cesium/`，可直接用于三维地球可视化开发。
- `gin-demo/` 目录当前为空，可用于后续扩展。

---

如需更详细的接口说明或自定义功能，请补充需求！ 