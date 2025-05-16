多商户发卡系统开发方案
### 一、系统架构设计

**整体架构**：
- 前端：Vue.js + ElementUI + 拖拽构建器组件
- 后端：Go Gin框架
- 数据库：MySQL(主库) + Redis(缓存)
- 部署：Nginx + Let's Encrypt

**模块化设计**：
1. 用户模块 - 负责用户认证、权限管理
2. 商户模块 - 商户入驻、等级管理、审核流程
3. 页面装修模块 - 拖拽式页面构建器、模板管理
4. 卡密模块 - 卡密生成、存储、分发、核销
5. 支付模块 - 支付渠道集成、订单管理、对账
6. 域名模块 - 二级域名分配、HTTPS证书管理
7. 系统管理模块 - 全局配置、数据统计
### 二、环境准备

**需要安装的软件**：
1. Go语言环境 (下载地址：https://go.dev/dl/)
2. Node.js (下载地址：https://nodejs.org/)
3. MySQL数据库 (下载地址：https://dev.mysql.com/downloads/installer/)
4. Redis (下载地址：https://github.com/microsoftarchive/redis/releases)
5. VSCode开发工具 (已安装)
6. Git版本控制工具 (下载地址：https://git-scm.com/downloads)

**开发工具配置**：
1. 在VSCode中安装Go扩展
2. 安装Vue.js扩展
3. 配置Go环境变量
### 三、数据库设计

**主要数据表**：

1. 用户表 (users)

id: 用户ID
username: 用户名
password: 密码(加密存储)
email: 邮箱
phone: 手机号
role: 角色(普通用户/管理员)
created_at: 创建时间
updated_at: 更新时间

2. 商户表 (merchants)

id: 商户ID
user_id: 关联用户ID
name: 商户名称
logo: 商户logo
status: 状态(待审核/已通过/已拒绝)
level: 商户等级(普通/VIP)
domain: 二级域名
created_at: 创建时间
updated_at: 更新时间

3. 页面模板表 (page_templates)

id: 模板ID
merchant_id: 商户ID
name: 模板名称
layout_data: 页面布局数据(JSON格式)
is_default: 是否默认模板
created_at: 创建时间
updated_at: 更新时间

4. 卡密表 (cards)

id: 卡密ID
merchant_id: 商户ID
product_id: 产品ID
card_code: 卡密内容(加密存储)
status: 状态(未使用/已使用/已过期)
expire_at: 过期时间
created_at: 创建时间
used_at: 使用时间

5. 订单表 (orders)

id: 订单ID
merchant_id: 商户ID
user_id: 用户ID
card_id: 卡密ID
amount: 金额
payment_method: 支付方式
transaction_id: 交易ID
status: 状态(待支付/已支付/已取消)
created_at: 创建时间
paid_at: 支付时间

6. 支付渠道表 (payment_channels)

id: 渠道ID
name: 渠道名称
config: 配置信息(JSON格式)
status: 状态(启用/禁用)
created_at: 创建时间
updated_at: 更新时间
### 四、功能模块实现

**1. 用户模块**

**功能说明**：
- 用户注册/登录
- 密码找回/重置
- 个人信息管理

**开发要求**：
- 使用JWT进行身份验证
- 密码使用bcrypt加密存储
- 实现RBAC权限控制

**2. 商户模块**

**功能说明**：
- 商户入驻申请
- 商户信息管理
- 商户等级管理
- 商户审核流程

**开发要求**：
- 实现多等级商户功能
- 设计人工审核流程
- 商户信息修改需审核

**3. 页面装修模块**

**功能说明**：
- 拖拽式页面构建器
- 页面元素管理
- 模板保存/加载
- 响应式布局设计

**开发要求**：
- 使用Vue.js实现拖拽功能
- 支持PC和移动端自适应
- 设计所见即所得(WYSIWYG)界面
- 页面数据JSON格式存储

**4. 卡密模块**

**功能说明**：
- 卡密批量生成
- 卡密加密存储
- 卡密分发
- 卡密核销
- 卡密查询统计

**开发要求**：
- 支持自定义卡密格式
- 卡密内容加密存储
- 实现卡密使用日志
- 支持批量操作

**5. 支付模块**

**功能说明**：
- 支付渠道管理
- 订单创建
- 支付回调处理
- 对账功能
- 退款管理

**开发要求**：
- 支持多种支付渠道
- 实现支付安全机制
- 设计幂等性回调处理
- 每日对账功能

**6. 域名模块**

**功能说明**：
- 二级域名分配
- 自定义域名管理
- HTTPS证书申请/续期
- 域名绑定管理

**开发要求**：
- 实现自动化证书管理
- 支持域名解析验证
- 证书到期提醒

**7. 系统管理模块**

**功能说明**：
- 系统配置管理
- 数据统计分析
- 操作日志管理
- 商户数据管理

**开发要求**：
- 设计系统配置中心
- 实现数据可视化
- 操作日志审计功能
### 五、开发流程

**1. 后端开发**

**技术栈**：
- 框架：Gin
- 数据库：GORM
- 认证：JWT
- 缓存：Redis
- 配置：Viper

**目录结构**：

card-system/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/       # 配置管理
│   ├── controller/   # 控制器层
│   ├── model/        # 模型层
│   ├── repository/   # 数据访问层
│   ├── service/      # 业务逻辑层
│   ├── middleware/   # 中间件
│   └── utils/        # 工具函数
├── pkg/
│   ├── auth/         # 认证模块
│   ├── payment/      # 支付模块
│   ├── domain/       # 域名模块
│   └── card/         # 卡密模块
├── static/           # 静态文件
└── templates/        # 模板文件

**2. 前端开发**

**技术栈**：
- 框架：Vue.js
- UI组件库：ElementUI
- 状态管理：Vuex
- 路由：Vue Router
- 拖拽组件：Vue.Draggable

**目录结构**：

card-system-frontend/
├── src/
│   ├── api/          # API请求
│   ├── assets/       # 静态资源
│   ├── components/   # 组件
│   ├── router/       # 路由
│   ├── store/        # 状态管理
│   ├── utils/        # 工具函数
│   └── views/        # 视图
│       ├── admin/    # 管理后台
│       └── merchant/ # 商户后台
├── public/           # 公共文件
└── package.json      # 依赖配置
### 六、部署方案

**部署架构**：

客户端 <-> Nginx <-> Go应用 <-> MySQL
                         <-> Redis

**部署步骤**：

1. 服务器环境准备：
   - 安装Nginx
   - 安装MySQL
   - 安装Redis
   - 配置防火墙

2. 应用部署：
   - 上传后端代码到服务器
   - 配置环境变量
   - 构建并运行Go应用
   - 上传前端代码到Nginx目录

3. Nginx配置：
   - 配置主域名
   - 配置二级域名通配符
   - 配置HTTPS证书
   - 设置反向代理

4. 证书管理：
   - 安装Certbot
   - 申请证书
   - 配置自动续期

5. 性能优化：
   - 配置Gzip压缩
   - 设置静态文件缓存
   - 配置数据库连接池
   - 实现Redis缓存
### 七、开发进度安排
**第一阶段：基础架构搭建**
- 完成Go项目初始化
- 配置数据库连接
- 实现用户认证模块
- 搭建前端基础框架

**第二阶段：核心功能开发**
- 开发商户入驻与管理模块
- 实现卡密生成与管理功能
- 开发支付集成模块
- 实现页面装修模块基础功能

**第三阶段：高级功能开发**
- 完善页面装修拖拽功能
- 开发域名管理模块
- 实现系统管理功能
- 开发数据统计分析功能

**第四阶段：测试与优化**
- 编写单元测试
- 进行集成测试
- 性能优化
- 安全漏洞扫描

**第五阶段：部署与上线**
- 完善部署文档
- 实现自动化部署脚本
- 系统上线
- 用户培训

好的！我们从最基础的环境搭建开始，逐步完成各个模块。以下是第一步：**环境准备与项目初始化**。

### **第一步：安装必备工具**
#### **1. 安装Go语言环境**
- 下载地址：[Go官方下载页](https://go.dev/dl/)，选择Windows版本（如go1.21.0.windows-amd64.msi）。
- 安装后，验证环境变量：
  - 打开命令提示符，输入 `go version`，应显示版本号（如 `go version go1.21.0 windows/amd64`）。
  - 若提示找不到命令，手动将 `C:\Program Files\go\bin` 添加到系统环境变量的 `Path` 中。

#### **2. 安装Node.js**
- 下载地址：[Node.js LTS版](https://nodejs.org/en/download/)，选择Windows Installer。
- 安装后，验证：
  bash
  node -v    # 显示版本号（如v20.6.1）
  npm -v     # 显示npm版本
 
#### **3. 安装MySQL**
- 下载地址：[MySQL Installer](https://dev.mysql.com/downloads/installer/)，选择Community Server。
- 安装时记住设置的 **root密码**，并创建一个数据库（后续步骤用）。
  sql
  -- 创建数据库（可在MySQL命令行执行）
  CREATE DATABASE card_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
 
#### **4. 安装Redis**
- 下载地址：[Redis for Windows](https://github.com/microsoftarchive/redis/releases)，下载 `Redis-x64-xxx.msi`。
- 安装后，启动Redis服务：
  - 打开命令提示符，进入Redis安装目录（默认 `C:\Program Files\Redis`），输入：
    bash
    redis-server --service-start
### **第二步：初始化Go项目（后端）**
#### **1. 创建项目目录**
bash
# 在任意目录创建项目文件夹
mkdir card-system-backend
cd card-system-backend
# 初始化Go模块（注意替换你的模块名，如github.com/你的用户名/card-system）
go mod init card-system

#### **2. 安装依赖**
bash
# 安装Gin框架
go get -u github.com/gin-gonic/gin
# 安装GORM数据库驱动（MySQL）
go get -u gorm.io/gorm gorm.io/driver/mysql
# 安装JWT认证库
go get -u github.com/golang-jwt/jwt/v4
# 安装Viper配置库
go get -u github.com/spf13/viper
# 安装Redis缓存库
go get -u github.com/go-redis/redis/v8

#### **3. 创建配置文件 `config/config.go`**
go
package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("app")       // 配置文件名（app.yaml）
	viper.SetConfigType("yaml")      // 配置文件类型
	viper.AddConfigPath("config/")   // 配置文件路径
	viper.AutomaticEnv()             // 读取环境变量

	// 初始化默认配置
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.dsn", "root:你的数据库密码@tcp(127.0.0.1:3306)/card_system?charset=utf8mb4&parseTime=True&loc=Local")
	viper.SetDefault("jwt.secret", "your-secret-key-123")
	viper.SetDefault("redis.addr", "localhost:6379")
}

// 获取端口
func Port() int {
	return viper.GetInt("server.port")
}

// 获取数据库DSN
func DSN() string {
	return viper.GetString("database.dsn")
}

// 获取JWT密钥
func JWTSecret() string {
	return viper.GetString("jwt.secret")
}

// 获取Redis地址
func RedisAddr() string {
	return viper.GetString("redis.addr")
}
#### **4. 创建数据库模型 `models/user.go`**
go
package models

import (
	"gorm.io/gorm"
	"time"
)

// 用户模型
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"not null" json:"-"` // 密码不返回给前端
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Phone     string         `json:"phone"`
	Role      string         `gorm:"default:user" json:"role"` // 角色：user/admin
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// 加密密码
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// 验证密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
### **第三步：编写基础服务（用户认证）**
#### **1. 创建用户控制器 `controllers/user_controller.go`**
go
package controllers

import (
	"card-system/models"
	"card-system/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// 注册用户
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 加密密码
	if err := user.SetPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}

	// 保存到数据库（假设已初始化DB，后续步骤完善数据库连接）
	if result := utils.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户创建失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// 登录获取JWT
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询用户
	var dbUser models.User
	if result := utils.DB.Where("username = ?", user.Username).First(&dbUser); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}

	// 验证密码
	if !dbUser.CheckPassword(user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	// 生成JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    dbUser.ID,
		"role":  dbUser.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 过期时间24小时
		"iat":   time.Now().Unix(),
		"nbf":   time.Now().Unix(),
	})

	tokenStr, err := token.SignedString([]byte(config.JWTSecret()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}
#### **2. 创建工具包 `utils/utils.go`（暂存数据库连接）**
go
package utils

import (
	"card-system/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := config.DSN()
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
### **第四步：编写主入口 `cmd/main.go`**
go
package main

import (
	"card-system/config"
	"card-system/controllers"
	"card-system/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 初始化配置
	config.Init()

	// 初始化数据库
	if err := utils.InitDB(); err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer func() {
		sqlDB, _ := utils.DB.DB()
		sqlDB.Close()
	}()

	// 初始化Gin引擎
	r := gin.Default()

	// 跨域中间件
	r.Use(gin.Cors(gin.CorsConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 用户路由
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
	}

	// 启动服务
	log.Printf("服务启动在 :%d", config.Port())
	r.Run(":" + strconv.Itoa(config.Port()))
}
### **当前进度检查**
1. 确保所有依赖已安装：
   bash
   go mod tidy # 下载缺失依赖
   
2. 在VSCode中打开项目目录，结构应如下：
   
   card-system-backend/
   ├── cmd/
   │   └── main.go
   ├── config/
   │   └── config.go
   ├── controllers/
   │   └── user_controller.go
   ├── models/
   │   └── user.go
   ├── utils/
   │   └── utils.go
   ├── go.mod
   └── go.sum
   
3. 创建配置文件 `config/app.yaml`（与config.go同目录）：
   yaml
   server:
     port: 8080
   database:
     dsn: "root:你的数据库密码@tcp(127.0.0.1:3306)/card_system?charset=utf8mb4&parseTime=True&loc=Local"
   jwt:
     secret: "your-secret-key-123" # 请替换为复杂密钥
   redis:
     addr: "localhost:6379"
   
### **下一步计划**
1. 完成数据库迁移（创建表）。
2. 开发商户入驻模块。
3. 搭建前端基础框架（Vue.js）。



没问题，我会详细指导你如何在 MySQL 8 中创建数据库。这里提供两种方法：使用命令行工具和使用 MySQL Workbench。


### **方法一：使用命令行工具创建数据库**
#### 步骤 1：打开 MySQL 命令行客户端
1. **Windows 系统**：
   - 按下 `Win + R` 打开运行窗口。
   - 输入 `cmd` 并回车，打开命令提示符。
   - 在命令提示符中输入以下命令（替换 `your_username` 为你的 MySQL 用户名，通常默认是 `root`）：
     ```bash
     mysql -u your_username -p
     ```
   - 按回车后，系统会提示你输入密码。输入密码（输入时密码不会显示，输完直接按回车）。

#### 步骤 2：创建数据库
- 成功登录后，在 MySQL 命令行中输入以下命令创建数据库：
  ```sql
  CREATE DATABASE card_system CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
  ```
  - **说明**：
    - `CHARACTER SET utf8mb4`：支持存储包括 emoji 在内的所有 Unicode 字符。
    - `COLLATE utf8mb4_unicode_ci`：设置不区分大小写的排序规则。

#### 步骤 3：验证数据库创建
- 输入以下命令查看所有数据库：
  ```sql
  SHOW DATABASES;
  ```
- 如果你看到 `card_system` 在列表中，说明创建成功。

#### 步骤 4：退出 MySQL 命令行
- 输入以下命令退出：
  ```sql
  EXIT;
  ```

### **方法二：使用 MySQL Workbench 创建数据库**
#### 步骤 1：打开 MySQL Workbench
1. 在开始菜单或桌面上找到 MySQL Workbench 并打开。
2. 连接到你的 MySQL 服务器（如果是第一次连接，需要配置连接信息）：
   - 点击 `+` 按钮添加新连接。
   - 填写连接名称（如 `Local MySQL`）、主机名（通常是 `localhost`）、端口（默认 `3306`）、用户名（默认 `root`）和密码。
   - 点击 `Test Connection` 测试连接，成功后点击 `OK`。

#### 步骤 2：打开 SQL 查询窗口
- 双击已创建的连接，打开主界面。
- 在左侧导航栏中，点击 `SCHEMAS` 标签。
- 右键点击空白处，选择 `Create Schema`（创建模式）。

#### 步骤 3：配置数据库
- 在弹出的对话框中：
  - **Schema Name**：输入 `card_system`。
  - **Default Character Set**：选择 `utf8mb4`。
  - **Default Collation**：选择 `utf8mb4_unicode_ci`。
- 点击 `Apply`，然后在新窗口中点击 `Apply` 确认创建。

#### 步骤 4：验证数据库创建
- 在左侧 `SCHEMAS` 列表中，刷新后应该能看到 `card_system`。


### **后续配置**
创建数据库后，需要在你的项目配置文件中更新数据库连接信息。打开 `card-system/config/config.yaml`，修改以下部分（替换为你的实际信息）：

```yaml
db:
  host: localhost
  port: 3306
  user: root          # 你的 MySQL 用户名
  password: yourpassword  # 你的 MySQL 密码
  name: card_system   # 刚创建的数据库名
  ssl_mode: disable
```
### **常见问题解决**
1. **忘记 MySQL 密码**：
   - 如果你忘记了 `root` 密码，可以通过以下步骤重置：
     ```bash
     # 1. 停止 MySQL 服务
     net stop mysql80  # 服务名可能是 MySQL80 或 MySQL

     # 2. 以安全模式启动 MySQL（跳过权限验证）
     mysqld --skip-grant-tables

     # 3. 打开新的命令提示符，直接登录 MySQL
     mysql -u root

     # 4. 在 MySQL 中重置密码
     USE mysql;
     ALTER USER 'root'@'localhost' IDENTIFIED BY 'new_password';
     FLUSH PRIVILEGES;
     EXIT;

     # 5. 停止 MySQL 服务，然后正常启动
     ```

