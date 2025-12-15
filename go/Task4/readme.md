# Gin Blog Backend

一个基于 **Go + Gin + GORM + JWT** 的博客系统后端，实现了用户注册、登录认证、文章增删改查、评论与添加评论等核心功能，适合作为学习和中小型项目的后端模板。

---

## 一、项目功能

### 1. 用户模块
- 用户注册
- 用户登录
- JWT 鉴权认证

### 2. 文章模块
- 创建文章（需登录）
- 修改文章（仅作者）
- 删除文章（仅作者）
- 文章列表
- 文章详情

### 3. 评论模块
- 添加评论（需登录）
- 评论列表（按文章）

---

## 二、技术栈

- **Go**
- **Gin** v1.11.0 Web 框架
- **GORM**  v1.31.1 ORM（支持 MySQL / PostgreSQL / SQLite）
- **JWT** V5 用户认证
- **Zap** v1.27.1 日志

---

## 三、运行环境要求

| 组件 | 版本要求 |
|---|---|
| Go | >= 1.25 |
| 数据库 | MySQL 5.7+（或其他 GORM 支持的数据库） |
| 操作系统 | Windows / Linux / macOS |

---

## 四、项目结构示例

```text
blog-backend
├── db
├── log
├── model
├── service
├── main.go
├── go.mod
└── README.md
```

---

## 五、依赖安装

### 1️⃣ 安装 Go 依赖

```bash
go mod tidy
```

主要依赖：

```text
github.com/gin-gonic/gin
gorm.io/gorm
gorm.io/driver/mysql
github.com/golang-jwt/jwt/v5
go.uber.org/zap
```

---

### 2️⃣ 数据库准备（以 MySQL 为例）

```sql
CREATE DATABASE blog DEFAULT CHARSET utf8mb4;
```

配置数据库连接（示例）：

```yaml
database:
  dsn: root:password@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local
```

启动时会自动进行数据表迁移（AutoMigrate）。

---

## 六、启动方式

### 方式一：直接运行

```bash
go run cmd/main.go
```

### 方式二：编译后运行

```bash
go build -o blog
./blog
```

默认启动地址：

```text
http://localhost:8080
```

---

## 七、接口测试说明

可以使用 **Postman / Apifox / curl** 对接口进行测试。

### 通用说明
- 登录成功后会返回 `token`
- 需要认证的接口需在 Header 中携带：

```text
Authorization: Bearer <token>
```

---

## 八、接口说明与测试用例（基于当前路由）

> 接口前缀统一为：`/api/v1`

---

### 1️⃣ 用户注册

**POST** `/api/v1/register`

请求体：

```json
{
  "email": "test@example.com",
  "password": "123456",
  "nickname": "testuser"
}
```

响应示例：

```json
{
  "code": 0,
  "message": "register success"
}
```

测试结果：✅ 注册成功

---

### 2️⃣ 用户登录

**POST** `/api/v1/login`

请求体：

```json
{
  "email": "test@example.com",
  "password": "123456"
}
```

响应示例：

```json
{
  "token": "<jwt_token>"
}
```

测试结果：✅ 返回 JWT Token

---

### 🔐 鉴权说明（重要）

以下接口需要登录后访问，请在 Header 中携带 JWT：

```text
Authorization: Bearer <jwt_token>
```

---

### 3️⃣ 创建文章

**POST** `/api/v1/post`

请求体：

```json
{
  "title": "第一篇博客",
  "content": "这是文章内容"
}
```

测试结果：✅ 创建成功

---

### 4️⃣ 修改文章（仅作者）

**PUT** `/api/v1/post/:ID`

示例：

```text
PUT /api/v1/post/1
```

请求体：

```json
{
  "title": "更新后的标题",
  "content": "更新后的内容"
}
```

测试结果：✅ 修改成功

---

### 5️⃣ 删除文章（仅作者）

**DELETE** `/api/v1/post/:ID`

示例：

```text
DELETE /api/v1/post/1
```

测试结果：✅ 删除成功

---

### 6️⃣ 获取文章列表

**GET** `/api/v1/post`

响应示例：

```json
[
  {
    "id": 1,
    "title": "第一篇博客",
    "author": "testuser"
  }
]
```

测试结果：✅ 返回文章列表

---

### 7️⃣ 获取文章详情

**GET** `/api/v1/post/:ID`

示例：

```text
GET /api/v1/post/1
```

测试结果：✅ 返回文章详情

---

### 8️⃣ 添加评论

**POST** `/api/v1/comment`

请求体：

```json
{
  "post_id": 1,
  "content": "写得不错！"
}
```

测试结果：✅ 评论成功

---

### 9️⃣ 获取文章评论列表

**GET** `/api/v1/post/:ID/comment`

示例：

```text
GET /api/v1/post/1/comment
```

响应示例：

```json
[
  {
    "id": 1,
    "content": "写得不错！",
    "user": "testuser"
  }
]
```

测试结果：✅ 返回评论列表

---

## 九、常见问题

### Q1：返回 401 Unauthorized？
- 未携带 Token
- Token 已过期
- Authorization 格式错误

### Q2：数据库连接失败？
- 检查 DSN
- 确认数据库已启动

---

## 十、后续可扩展方向

- 文章标签 / 分类
- 点赞 / 收藏
- 评论回复
- 管理员后台
- 接入 Redis / 缓存

---

## License

MIT License

