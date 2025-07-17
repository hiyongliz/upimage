# UpImage

一个简单高效的命令行工具，用于自动化上传Docker镜像到多个云容器镜像仓库。支持华为云SWR、阿里云ACR、腾讯云TCR等云服务商。

**📖 文档语言 / Documentation Languages:**
- [中文文档](README.md) (当前页面)
- [English Documentation](README_EN.md)

> **⚠️ 声明**: 本README文档由AI生成，可能存在不准确或过时的信息。如果在使用过程中遇到问题，请通过 [Issues](https://github.com/hiyongliz/upimage/issues) 反馈，我们会及时更正和完善文档。

## 功能特性

- 🚀 **一键上传**: 单一命令完成镜像拉取、标记、推送的完整流程
- 🔄 **自动化流程**: 可选择自动创建命名空间、设置仓库为公开访问
- 🎯 **灵活配置**: 支持自定义命名空间、区域、公开设置等选项
- � **多云支持**: 支持华为云SWR、阿里云ACR、腾讯云TCR等主流云服务商
- 🌍 **多区域支持**: 支持各云服务商的所有可用区域
- 🛡️ **容错处理**: 智能处理已存在的命名空间，避免重复创建错误
- ⚡ **高效便捷**: 减少手动操作，提升DevOps工作效率
- 📊 **详细输出**: 实时显示操作进度和状态信息

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

### 华为云SWR

1. **设置环境变量**:
   ```bash
   export HUAWEICLOUD_SDK_AK="your_access_key"
   export HUAWEICLOUD_SDK_SK="your_secret_key"
   ```

2. **上传镜像到华为云SWR**:
   ```bash
   upimage nginx:latest --registry swr --namespace myproject
   ```

### 阿里云ACR

1. **登录阿里云容器镜像服务**:
   ```bash
   docker login registry.cn-hangzhou.aliyuncs.com
   ```

2. **上传镜像到阿里云ACR**:
   ```bash
   upimage nginx:latest --registry acr --region cn-hangzhou --namespace myproject
   ```

### 腾讯云TCR

1. **登录腾讯云容器镜像服务**:
   ```bash
   docker login ccr.ccs.tencentyun.com
   ```

2. **上传镜像到腾讯云TCR**:
   ```bash
   upimage nginx:latest --registry tcr --namespace myproject
   ```

## 使用方法

### 基本用法

```bash
upimage <image> [flags]
```

### 命令行参数

- `--registry, -g`: 指定云服务商（支持: `swr`, `acr`, `tcr`，默认: `swr`）
- `--region, -r`: 指定云服务区域（默认: `cn-south-1`）
- `--namespace, -n`: 指定目标命名空间（默认: `default`）
- `--create-namespace`: 是否自动创建命名空间（默认: `true`，仅SWR支持）
- `--public`: 是否将仓库设置为公开（默认: `false`，仅SWR支持）
- `--help, -h`: 显示帮助信息

### 获取帮助

```bash
# 查看帮助信息
upimage --help

# 查看版本信息
upimage --version
```

### 示例

#### 华为云SWR

```bash
# 基本使用 - 上传到华为云SWR默认命名空间
upimage nginx:1.28 --registry swr

# 指定自定义命名空间
upimage nginx:1.28 --registry swr --namespace myproject

# 指定不同的区域
upimage nginx:1.28 --registry swr --region cn-north-4

# 上传到自定义命名空间并设置为公开
upimage nginx:1.28 --registry swr --namespace myproject --public

# 不自动创建命名空间（命名空间必须已存在）
upimage nginx:1.28 --registry swr --namespace existing-ns --create-namespace=false
```

#### 阿里云ACR

```bash
# 上传到阿里云ACR杭州区域
upimage nginx:1.28 --registry acr --region cn-hangzhou --namespace myproject

# 上传到阿里云ACR北京区域
upimage nginx:1.28 --registry acr --region cn-beijing --namespace myproject
```

#### 腾讯云TCR

```bash
# 上传到腾讯云TCR（区域参数对TCR无效）
upimage nginx:1.28 --registry tcr --namespace myproject
```

#### 完整配置示例

```bash
# 华为云SWR完整配置
upimage nginx:1.28 \
  --registry swr \
  --region cn-north-1 \
  --namespace myproject \
  --public \
  --create-namespace
```

### 工作流程

工具支持多云镜像仓库同步，根据选择的云服务商执行相应的流程：

#### 华为云SWR流程

当你执行 `upimage nginx:1.28 --registry swr --namespace myproject --region cn-south-1 --public` 时：

1. **解析镜像信息**: 提取仓库名和标签
2. **创建命名空间**: 在指定区域的华为云SWR中创建命名空间（如果启用且不存在）
3. **拉取镜像**: `docker pull nginx:1.28`
4. **重新标记**: `docker tag nginx:1.28 swr.cn-south-1.myhuaweicloud.com/myproject/nginx:1.28`
5. **推送镜像**: `docker push swr.cn-south-1.myhuaweicloud.com/myproject/nginx:1.28`
6. **设置权限**: 将仓库设置为公开访问（如果启用）

#### 阿里云ACR流程

当你执行 `upimage nginx:1.28 --registry acr --namespace myproject --region cn-hangzhou` 时：

1. **解析镜像信息**: 提取仓库名和标签
2. **拉取镜像**: `docker pull nginx:1.28`
3. **重新标记**: `docker tag nginx:1.28 registry.cn-hangzhou.aliyuncs.com/myproject/nginx:1.28`
4. **推送镜像**: `docker push registry.cn-hangzhou.aliyuncs.com/myproject/nginx:1.28`

#### 腾讯云TCR流程

当你执行 `upimage nginx:1.28 --registry tcr --namespace myproject` 时：

1. **解析镜像信息**: 提取仓库名和标签
2. **拉取镜像**: `docker pull nginx:1.28`
3. **重新标记**: `docker tag nginx:1.28 ccr.ccs.tencentyun.com/myproject/nginx:1.28`
4. **推送镜像**: `docker push ccr.ccs.tencentyun.com/myproject/nginx:1.28`

### 输出示例

```
Processing image: nginx:1.28
  Namespace: myproject
  Repository: nginx
  Tag: 1.28
  Target region: cn-south-1
Creating namespace "myproject" if it doesn't exist...
Pulling image "nginx:1.28"...
Tagging image as "swr.cn-south-1.myhuaweicloud.com/myproject/nginx:1.28"...
Pushing image "swr.cn-south-1.myhuaweicloud.com/myproject/nginx:1.28" to swr...
Setting repository "myproject"/"nginx" as public...
✅ Successfully synced image "nginx:1.28" to "swr.cn-south-1.myhuaweicloud.com/myproject/nginx:1.28"
Setting repository "myproject/nginx" as public...
✅ Successfully synced image "nginx:1.28" to "swr.cn-south-1.myhuaweicloud.com/myproject/nginx:1.28"
```

## 镜像名称处理规则

工具现在采用更灵活的命名空间管理方式，支持多云服务商的不同镜像格式：

| 输入镜像 | 默认命名空间 | 仓库名 | 标签 | 自定义命名空间 |
|---------|-------------|--------|------|-------------|
| `nginx:1.28` | `default` | `nginx` | `1.28` | 使用 `--namespace` 指定 |
| `redis:latest` | `default` | `redis` | `latest` | 使用 `--namespace` 指定 |
| `myapp` | `default` | `myapp` | `latest` | 使用 `--namespace` 指定 |
| `registry.io/myapp:v1.0` | `default` | `myapp` | `v1.0` | 使用 `--namespace` 指定 |

### 目标镜像格式

不同云服务商的镜像格式：

- **华为云SWR**: `swr.{region}.myhuaweicloud.com/{namespace}/{repo}:{tag}`
- **阿里云ACR**: `registry.{region}.aliyuncs.com/{namespace}/{repo}:{tag}`
- **腾讯云TCR**: `ccr.ccs.tencentyun.com/{namespace}/{repo}:{tag}`

> **重要变更**: 工具不再从镜像名称中自动解析命名空间，而是使用 `--namespace` 参数或默认值 `default`。这提供了更好的控制和一致性。

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
- [Fang](https://github.com/charmbracelet/fang) - 优雅的CLI执行器
- [华为云Go SDK](https://github.com/huaweicloud/huaweicloud-sdk-go-v3) - 华为云服务API客户端

## 性能和限制

### 性能特点
- 🚀 **实时反馈**: 详细的操作进度显示，让用户了解当前状态
- 📦 **增量上传**: 利用Docker分层存储，只上传变更的层
- 🔄 **错误恢复**: 详细的错误信息帮助快速定位问题
- ⚙️ **灵活配置**: 支持多种参数组合满足不同需求
- 🌐 **多云适配**: 统一接口适配不同云服务商的API差异

### 使用限制

#### 通用限制
- 网络要求：需要稳定的网络连接到对应云服务商
- 权限要求：需要Docker运行权限和相应云服务商的操作权限

#### 各云服务商特定限制

**华为云SWR**:
- 镜像大小：单个镜像不超过10GB
- 命名空间：支持自动创建和公开设置
- 区域：支持华为云所有可用区域

**阿里云ACR**:
- 镜像大小：个人版单个镜像不超过1GB，企业版不超过10GB
- 命名空间：需要预先在阿里云控制台创建
- 区域：支持阿里云所有可用区域

**腾讯云TCR**:
- 镜像大小：根据实例类型有不同限制
- 命名空间：需要预先在腾讯云控制台创建
- 区域：使用统一域名，区域参数无效

### 最佳实践
- 使用就近的区域减少网络延迟
- 合理规划命名空间结构
- 根据云服务商特性选择合适的配置
- 定期清理不需要的镜像版本
- 建议同时处理的镜像数量不超过5个

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

1. **华为云SWR认证失败**
   ```
   Error initializing Up: failed to create credentials
   ```
   解决方案：检查 `HUAWEICLOUD_SDK_AK` 和 `HUAWEICLOUD_SDK_SK` 环境变量是否正确设置。

2. **阿里云/腾讯云认证失败**
   ```
   Error executing upimage: failed to push image
   ```
   解决方案：确保已通过 `docker login` 命令登录到对应的云服务商。

3. **不支持的云服务商**
   ```
   Invalid registry: xxx. Supported values are 'swr', 'acr' or 'tcr'.
   ```
   解决方案：使用支持的云服务商参数：`swr`（华为云）、`acr`（阿里云）、`tcr`（腾讯云）。

4. **Docker命令失败**
   ```
   Error executing Up: exit status 1
   ```
   解决方案：确保Docker守护进程正在运行，且有权限访问指定的镜像。

5. **网络连接问题**
   - 确保网络连接正常
   - 检查是否需要配置代理
   - 验证对应云服务在当前区域是否可用

6. **权限问题**
   ```
   Error: permission denied
   ```
   解决方案：确保当前用户有Docker操作权限，可能需要将用户添加到docker用户组。

7. **命名空间不存在**
   ```
   failed to push image: namespace not found
   ```
   解决方案：
   - 华为云SWR：启用 `--create-namespace` 参数（默认启用）
   - 阿里云ACR/腾讯云TCR：在对应控制台预先创建命名空间

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

### 支持的云服务商区域

#### 华为云SWR区域

- `cn-south-1` - 华南1（广州）**[默认]**
- `cn-north-4` - 华北4（北京四）
- `cn-east-3` - 华东1（上海一）
- `cn-east-2` - 华东2（上海二）
- `cn-north-1` - 华北1（北京一）
- `cn-southwest-2` - 西南1（贵阳一）

#### 阿里云ACR区域

- `cn-hangzhou` - 华东1（杭州）
- `cn-beijing` - 华北2（北京）
- `cn-shanghai` - 华东2（上海）
- `cn-shenzhen` - 华南1（深圳）
- `cn-qingdao` - 华北1（青岛）

#### 腾讯云TCR

腾讯云TCR使用统一域名，不需要指定区域。

## 批量同步

### 使用sync.sh脚本

项目提供了`sync.sh`脚本用于批量同步镜像：

```bash
# 基本用法（华为云SWR）
./sync.sh image.txt cn-south-1 myproject

# 指定云服务商
./sync.sh image.txt cn-south-1 myproject swr

# 阿里云ACR
./sync.sh image.txt cn-hangzhou myproject acr

# 腾讯云TCR
./sync.sh image.txt "" myproject tcr
```

### GitHub Actions自动化

项目包含GitHub Actions工作流，支持：

- 推送到main分支时自动同步
- 手动触发同步
- 多云服务商选择
- 自定义区域和命名空间

详细配置参见 [.github/workflows/README.md](.github/workflows/README.md)。

### 调试模式

工具内置了详细的输出信息，无需额外配置：

```bash
upimage nginx:1.28 --namespace myproject --region cn-north-4
```

输出将包含：
- 镜像解析信息
- 操作步骤进度
- 错误详情（如果发生）
- 成功确认信息

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
