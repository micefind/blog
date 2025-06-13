# 项目名称

## 简介

本项目是一个使用 Go 和 Gin 框架开发的 API 服务。它实现了用户管理、文章管理等功能。

## 环境要求

- Go 1.18 或更高版本
- MySQL 数据库
- Redis（可选）

## 安装和运行

### 1. 克隆仓库

```bash
git clone https://github.com/micefind/blog.git
```

### 2. 配置数据库

在 config/config.yaml 中配置 MySQL 数据库连接信息。

```yaml
mysql:
  host: "localhost"
  port: 3306
  user: "root"
  password: "123456"
  dbname: "blog_db"
```

### 3. 安装依赖

```
go mod tidy
```

### 4. 运行项目

```
go run main.go
```

项目默认会在 8080 端口启动，访问 http://localhost:8080 进行测试。

### 5. 测试接口

GET /api/user/list - 获取所有用户
