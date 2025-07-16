# UpImage

一个简单高效的命令行工具，用于自动化上传Docker镜像到华为云软件仓库(SWR)。

> **⚠️ 声明**: 本README文档由AI生成，可能存在不准确或过时的信息。如果在使用过程中遇到问题，请通过 [Issues](https://github.com/hiyongliz/upimage/issues) 反馈，我们会及时更正和完善文档。

## 功能特性

- 🚀 **一键上传**: 单一命令完成镜像拉取、标记、推送的完整流程
- 🔄 **自动化流程**: 自动创建命名空间、设置仓库为公开访问
- 🎯 **智能解析**: 自动解析镜像名称，提取命名空间、仓库名和标签
- 🌍 **多区域支持**: 支持华为云所有SWR区域，可通过 `--region` 参数指定
- 🛡️ **容错处理**: 智能处理已存在的命名空间，避免重复创建错误
- ⚡ **高效便捷**: 减少手动操作，提升DevOps工作效率

## 安装

### 从源码构建

```bash
git clone https://github.com/hiyongliz/upimage.git
cd upimage
make bin
```

构建完成后，会在当前目录生成 `upimage` 可执行文件。

### 直接下载

你也可以从 [Releases](https://github.com/hiyongliz/upimage/releases) 页面下载预编译的二进制文件。

### 安装到系统路径

```bash
# 将可执行文件移动到系统PATH中
sudo mv upimage /usr/local/bin/

# 验证安装
upimage --help
```

### 环境要求

- Go 1.24.4 或更高版本
- Docker (用于镜像操作)
- 华为云账号及API凭证

## 配置

在使用前，需要设置华为云认证环境变量：

```bash
export HUAWEICLOUD_SDK_AK="your_access_key"
export HUAWEICLOUD_SDK_SK="your_secret_key"
```

建议将这些环境变量添加到你的 shell 配置文件中（如 `~/.bashrc` 或 `~/.zshrc`）。

## 快速开始

1. **设置环境变量**:
   ```bash
   export HUAWEICLOUD_SDK_AK="your_access_key"
   export HUAWEICLOUD_SDK_SK="your_secret_key"
   ```

2. **上传你的第一个镜像**:
   ```bash
   upimage nginx:latest
   ```

3. **查看上传的镜像**:
   登录华为云控制台 → 容器镜像服务 → 我的镜像，即可看到上传的镜像。

## 使用方法

### 基本用法

```bash
upimage <image> [flags]
```

### 命令行参数

- `--region`: 指定华为云SWR区域（默认: `cn-south-1`）
- `--help, -h`: 显示帮助信息

### 获取帮助

```bash
# 查看帮助信息
upimage --help

# 查看版本信息
upimage --version
```

### 示例

```bash
# 上传一个标准镜像（使用默认区域）
upimage nginx:1.28

# 上传带命名空间的镜像
upimage myregistry/myapp:v1.0.0

# 上传镜像（将使用默认latest标签）
upimage redis

# 指定不同的区域上传镜像
upimage nginx:1.28 --region cn-north-4

# 上传到华东1区域
upimage myapp:latest --region cn-east-3
```

### 工作流程

当你执行 `upimage nginx:1.28 --region cn-south-1` 时，工具会自动执行以下步骤：

1. **拉取镜像**: `docker pull nginx:1.28`
2. **创建命名空间**: 在指定区域的华为云SWR中创建命名空间（如果不存在）
3. **重新标记**: `docker tag nginx:1.28 swr.cn-south-1.myhuaweicloud.com/nginx/nginx:1.28`
4. **推送镜像**: `docker push swr.cn-south-1.myhuaweicloud.com/nginx/nginx:1.28`
5. **设置公开**: 将仓库设置为公开访问

## 镜像名称解析规则

| 输入格式 | 命名空间 | 仓库名 | 标签 |
|---------|---------|--------|------|
| `nginx:1.28` | `nginx` | `nginx` | `1.28` |
| `myapp/redis:latest` | `myapp` | `redis` | `latest` |
| `redis` | `redis` | `redis` | `latest` |
| `myregistry/myapp` | `myregistry` | `myapp` | `latest` |

> **注意**: 如果镜像名称没有明确的命名空间，将使用默认命名空间 `lazylibrary`

## 项目结构

```
upimage/
├── main.go                 # 程序入口点
├── cmd/                    # CLI命令定义
│   └── root.go            # 根命令和参数处理
├── app/                    # 应用业务逻辑
│   └── app.go             # 主要业务流程实现
├── internal/services/      # 内部服务层
│   └── swr/               # SWR服务封装
│       └── swr.go         # SWR操作接口
├── pkg/                    # 可重用包
│   ├── swrapi/            # SWR API客户端
│   │   ├── swrapi.go      # API调用实现
│   │   └── models/        # 数据模型定义
│   │       └── response.go
│   └── utils/             # 工具函数
│       ├── utils.go       # 镜像解析工具
│       └── utils_test.go  # 单元测试
├── go.mod                  # Go模块定义
├── go.sum                  # 依赖版本锁定
├── LICENSE                 # MIT许可证文件
├── Makefile               # 构建脚本
└── README.md              # 项目文档
```

## 依赖库

- [Cobra](https://github.com/spf13/cobra) - 强大的CLI框架
- [华为云Go SDK](https://github.com/huaweicloud/huaweicloud-sdk-go-v3) - 华为云服务API客户端

## 性能和限制

### 性能特点
- 🚀 **并发处理**: 支持同时处理多个镜像操作
- 📦 **增量上传**: 利用Docker分层存储，只上传变更的层
- 🔄 **断点续传**: 网络中断时自动重试

### 使用限制
- 镜像大小：单个镜像不超过10GB（华为云SWR限制）
- 并发数量：建议同时处理的镜像数量不超过5个
- 网络要求：需要稳定的网络连接到华为云服务

### 最佳实践
- 使用就近的区域减少网络延迟
- 合理使用镜像标签避免覆盖
- 定期清理不需要的镜像版本

## 开发

### 运行测试

```bash
go test ./...
```

### 构建

```bash
# 构建当前平台可执行文件
make bin

# 或使用Go命令
go build -o upimage main.go
```

### 跨平台构建

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o upimage-linux main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o upimage.exe main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o upimage-darwin main.go
```

## 故障排除

### 常见问题

1. **认证失败**
   ```
   Error initializing Up: failed to create credentials
   ```
   解决方案：检查 `HUAWEICLOUD_SDK_AK` 和 `HUAWEICLOUD_SDK_SK` 环境变量是否正确设置。

2. **Docker命令失败**
   ```
   Error executing Up: exit status 1
   ```
   解决方案：确保Docker守护进程正在运行，且有权限访问指定的镜像。

3. **网络连接问题**
   - 确保网络连接正常
   - 检查是否需要配置代理
   - 验证华为云服务在当前区域是否可用

4. **权限问题**
   ```
   Error: permission denied
   ```
   解决方案：确保当前用户有Docker操作权限，可能需要将用户添加到docker用户组。

5. **镜像不存在**
   ```
   Error: pull access denied for image
   ```
   解决方案：检查镜像名称是否正确，确保有权限访问该镜像。

6. **区域不支持**
   ```
   Error: region not supported
   ```
   解决方案：检查指定的区域是否为有效的华为云区域代码。

### 支持的华为云区域

工具支持华为云所有SWR服务区域，常用区域包括：

- `cn-south-1` - 华南1（广州）**[默认]**
- `cn-north-4` - 华北4（北京四）
- `cn-east-3` - 华东1（上海一）
- `cn-east-2` - 华东2（上海二）
- `cn-north-1` - 华北1（北京一）
- `cn-southwest-2` - 西南1（贵阳一）

### 调试模式

可以通过设置环境变量启用详细日志：

```bash
export DEBUG=true
upimage nginx:1.28 --region cn-north-4
```

## 贡献

欢迎贡献代码！请遵循以下步骤：

1. Fork 这个仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 支持

如果你遇到问题或有功能建议，请：

- 创建 [Issue](https://github.com/hiyongliz/upimage/issues)
- 发送邮件到项目维护者
- 查看 [华为云SWR文档](https://support.huaweicloud.com/swr/)

---

Made with ❤️ for the DevOps community
