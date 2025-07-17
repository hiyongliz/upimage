# UpImage

A simple and efficient command-line tool for automated Docker image uploading to multiple cloud container registries. Supports Huawei Cloud SWR, Alibaba Cloud ACR, Tencent Cloud TCR, and other cloud service providers.

**📖 Documentation Languages:**
- [English Documentation](README_EN.md) (Current Page)
- [中文文档](README.md)

> **⚠️ Disclaimer**: This README document is AI-generated and may contain inaccurate or outdated information. If you encounter issues during use, please provide feedback through [Issues](https://github.com/hiyongliz/upimage/issues), and we will promptly correct and improve the documentation.

## Features

- 🚀 **One-Click Upload**: Complete image pull, tag, and push workflow with a single command
- 🔄 **Automated Workflow**: Optional automatic namespace creation and public repository settings
- 🎯 **Flexible Configuration**: Support for custom namespaces, regions, public settings, and more
- 🌐 **Multi-Cloud Support**: Support for major cloud service providers including Huawei Cloud SWR, Alibaba Cloud ACR, Tencent Cloud TCR
- 🌍 **Multi-Region Support**: Support for all available regions across cloud service providers
- 🛡️ **Error Handling**: Smart handling of existing namespaces to avoid duplicate creation errors
- ⚡ **Efficient and Convenient**: Reduce manual operations and improve DevOps workflow efficiency
- 📊 **Detailed Output**: Real-time display of operation progress and status information

## Installation

### Build from Source

```bash
git clone https://github.com/hiyongliz/upimage.git
cd upimage
make bin
```

After building, the `upimage` executable will be generated in the current directory.

### Direct Download

You can also download pre-compiled binary files from the [Releases](https://github.com/hiyongliz/upimage/releases) page.

### Install to System Path

```bash
# Move the executable to system PATH
sudo mv upimage /usr/local/bin/

# Verify installation
upimage --help
```

### Requirements

- Go 1.24.4 or higher
- Docker (for image operations)
- Cloud account and API credentials for the target cloud provider

## Configuration

### Huawei Cloud SWR

Set the required environment variables:

```bash
export HUAWEICLOUD_SDK_AK="your_access_key"
export HUAWEICLOUD_SDK_SK="your_secret_key"
```

### Alibaba Cloud ACR

Log in to Alibaba Cloud Container Registry:

```bash
docker login registry.cn-hangzhou.aliyuncs.com
```

### Tencent Cloud TCR

Log in to Tencent Cloud Container Registry:

```bash
docker login ccr.ccs.tencentyun.com
```

## Quick Start

### Huawei Cloud SWR

1. **Set environment variables**:
   ```bash
   export HUAWEICLOUD_SDK_AK="your_access_key"
   export HUAWEICLOUD_SDK_SK="your_secret_key"
   ```

2. **Upload image to Huawei Cloud SWR**:
   ```bash
   upimage nginx:latest --registry swr --namespace myproject
   ```

### Alibaba Cloud ACR

1. **Log in to Alibaba Cloud Container Registry**:
   ```bash
   docker login registry.cn-hangzhou.aliyuncs.com
   ```

2. **Upload image to Alibaba Cloud ACR**:
   ```bash
   upimage nginx:latest --registry acr --region cn-hangzhou --namespace myproject
   ```

### Tencent Cloud TCR

1. **Log in to Tencent Cloud Container Registry**:
   ```bash
   docker login ccr.ccs.tencentyun.com
   ```

2. **Upload image to Tencent Cloud TCR**:
   ```bash
   upimage nginx:latest --registry tcr --namespace myproject
   ```

## Usage

### Basic Usage

```bash
upimage <image> [flags]
```

### Command Line Parameters

- `--registry, -g`: Specify cloud service provider (supported: `swr`, `acr`, `tcr`, default: `swr`)
- `--region, -r`: Specify cloud service region (default: `cn-south-1`)
- `--namespace, -n`: Specify target namespace (default: `default`)
- `--create-namespace`: Whether to automatically create namespace (default: `true`, SWR only)
- `--public`: Whether to set repository as public (default: `false`, SWR only)
- `--help, -h`: Show help information

### Get Help

```bash
# View help information
upimage --help

# View version information
upimage --version
```

### Examples

#### Huawei Cloud SWR

```bash
# Basic usage - upload to Huawei Cloud SWR default namespace
upimage nginx:1.28 --registry swr

# Specify custom namespace
upimage nginx:1.28 --registry swr --namespace myproject

# Specify different region
upimage nginx:1.28 --registry swr --region cn-north-4

# Upload to custom namespace and set as public
upimage nginx:1.28 --registry swr --namespace myproject --public

# Don't auto-create namespace (namespace must exist)
upimage nginx:1.28 --registry swr --namespace existing-ns --create-namespace=false
```

#### Alibaba Cloud ACR

```bash
# Upload to Alibaba Cloud ACR Hangzhou region
upimage nginx:1.28 --registry acr --region cn-hangzhou --namespace myproject

# Upload to Alibaba Cloud ACR Beijing region
upimage nginx:1.28 --registry acr --region cn-beijing --namespace myproject
```

#### Tencent Cloud TCR

```bash
# Upload to Tencent Cloud TCR (region parameter is not effective for TCR)
upimage nginx:1.28 --registry tcr --namespace myproject
```

#### Complete Configuration Example

```bash
# Huawei Cloud SWR complete configuration
upimage nginx:1.28 \
  --registry swr \
  --region cn-north-4 \
  --namespace myproject \
  --public \
  --create-namespace
```

### Workflow

The tool supports multi-cloud image registry synchronization and executes corresponding workflows based on the selected cloud service provider:

#### Huawei Cloud SWR Workflow

When you execute `upimage nginx:1.28 --registry swr --namespace myproject --region cn-south-1 --public`:

1. **Parse image information**: Extract repository name and tag
2. **Create namespace**: Create namespace in the specified Huawei Cloud SWR region (if enabled and doesn't exist)
3. **Pull image**: `docker pull nginx:1.28`
4. **Re-tag**: `docker tag nginx:1.28 swr.cn-south-1.myhuaweicloud.com/myproject/nginx:1.28`
5. **Push image**: `docker push swr.cn-south-1.myhuaweicloud.com/myproject/nginx:1.28`
6. **Set permissions**: Set repository as public access (if enabled)

#### Alibaba Cloud ACR Workflow

When you execute `upimage nginx:1.28 --registry acr --namespace myproject --region cn-hangzhou`:

1. **Parse image information**: Extract repository name and tag
2. **Pull image**: `docker pull nginx:1.28`
3. **Re-tag**: `docker tag nginx:1.28 registry.cn-hangzhou.aliyuncs.com/myproject/nginx:1.28`
4. **Push image**: `docker push registry.cn-hangzhou.aliyuncs.com/myproject/nginx:1.28`

#### Tencent Cloud TCR Workflow

When you execute `upimage nginx:1.28 --registry tcr --namespace myproject`:

1. **Parse image information**: Extract repository name and tag
2. **Pull image**: `docker pull nginx:1.28`
3. **Re-tag**: `docker tag nginx:1.28 ccr.ccs.tencentyun.com/myproject/nginx:1.28`
4. **Push image**: `docker push ccr.ccs.tencentyun.com/myproject/nginx:1.28`

### Output Example

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
```

## Image Name Processing Rules

The tool now adopts a more flexible namespace management approach, supporting different image formats for multiple cloud service providers:

| Input Image | Default Namespace | Repository | Tag | Custom Namespace |
|-------------|------------------|------------|-----|------------------|
| `nginx:1.28` | `default` | `nginx` | `1.28` | Use `--namespace` to specify |
| `redis:latest` | `default` | `redis` | `latest` | Use `--namespace` to specify |
| `myapp` | `default` | `myapp` | `latest` | Use `--namespace` to specify |
| `registry.io/myapp:v1.0` | `default` | `myapp` | `v1.0` | Use `--namespace` to specify |

### Target Image Formats

Different cloud service providers have different image formats:

- **Huawei Cloud SWR**: `swr.{region}.myhuaweicloud.com/{namespace}/{repo}:{tag}`
- **Alibaba Cloud ACR**: `registry.{region}.aliyuncs.com/{namespace}/{repo}:{tag}`
- **Tencent Cloud TCR**: `ccr.ccs.tencentyun.com/{namespace}/{repo}:{tag}`

> **Important Change**: The tool no longer automatically parses namespaces from image names, but uses the `--namespace` parameter or default value `default`. This provides better control and consistency.

## Project Structure

```
upimage/
├── main.go                 # Program entry point
├── cmd/                    # CLI command definitions
│   └── root.go            # Root command and parameter handling
├── app/                    # Application business logic
│   └── app.go             # Core upload logic
├── internal/               # Internal packages
│   └── services/          # Service implementations
│       └── swr/           # Huawei Cloud SWR service
│           └── swr.go     # SWR operations
├── pkg/                    # Public packages
│   ├── swrapi/            # Huawei Cloud SWR API client
│   │   ├── swrapi.go      # API client implementation
│   │   ├── swrapi_test.go # API tests
│   │   └── models/        # Data models
│   │       └── response.go # Response structures
│   └── utils/             # Utility functions
│       ├── utils.go       # Image name parsing utilities
│       └── utils_test.go  # Utility tests
├── sync.sh                 # Batch sync script
├── Makefile               # Build configuration
└── README.md              # Project documentation
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - Powerful CLI framework
- [Fang](https://github.com/charmbracelet/fang) - Elegant CLI executor
- [Huawei Cloud Go SDK](https://github.com/huaweicloud/huaweicloud-sdk-go-v3) - Huawei Cloud service API client

## Performance and Limitations

### Performance Features
- 🚀 **Real-time Feedback**: Detailed operation progress display to keep users informed of current status
- 📦 **Incremental Upload**: Utilizes Docker layered storage, only uploads changed layers
- 🔄 **Error Recovery**: Detailed error information helps quickly locate problems
- ⚙️ **Flexible Configuration**: Supports various parameter combinations to meet different needs
- 🌐 **Multi-Cloud Adaptation**: Unified interface adapts to API differences across cloud service providers

### Usage Limitations

#### General Limitations
- Network Requirements: Requires stable network connection to corresponding cloud service providers
- Permission Requirements: Requires Docker runtime permissions and corresponding cloud service provider operation permissions

#### Cloud Service Provider Specific Limitations

**Huawei Cloud SWR**:
- Image Size: Single image not exceeding 10GB
- Namespace: Supports automatic creation and public settings
- Region: Supports all Huawei Cloud available regions

**Alibaba Cloud ACR**:
- Image Size: Personal edition single image not exceeding 1GB, Enterprise edition not exceeding 10GB
- Namespace: Needs to be pre-created in Alibaba Cloud console
- Region: Supports all Alibaba Cloud available regions

**Tencent Cloud TCR**:
- Image Size: Different limits based on instance type
- Namespace: Needs to be pre-created in Tencent Cloud console
- Region: Uses unified domain name, region parameter is not effective

### Best Practices
- Use nearby regions to reduce network latency
- Plan namespace structure properly
- Choose appropriate configurations based on cloud service provider characteristics
- Clean up unnecessary image versions regularly
- Recommend processing no more than 5 images simultaneously

### Supported Cloud Service Provider Regions

#### Huawei Cloud SWR Regions

- `cn-south-1` - South China 1 (Guangzhou) **[Default]**
- `cn-north-4` - North China 4 (Beijing-4)
- `cn-east-3` - East China 1 (Shanghai-1)
- `cn-east-2` - East China 2 (Shanghai-2)
- `cn-north-1` - North China 1 (Beijing-1)
- `cn-southwest-2` - Southwest 1 (Guiyang-1)

#### Alibaba Cloud ACR Regions

- `cn-hangzhou` - East China 1 (Hangzhou)
- `cn-beijing` - North China 2 (Beijing)
- `cn-shanghai` - East China 2 (Shanghai)
- `cn-shenzhen` - South China 1 (Shenzhen)
- `cn-qingdao` - North China 1 (Qingdao)

#### Tencent Cloud TCR

Tencent Cloud TCR uses a unified domain name and does not require region specification.

## Batch Synchronization

### Using sync.sh Script

The project provides a `sync.sh` script for batch image synchronization:

```bash
# Basic usage (Huawei Cloud SWR)
./sync.sh image.txt cn-south-1 myproject

# Specify cloud service provider
./sync.sh image.txt cn-south-1 myproject swr

# Alibaba Cloud ACR
./sync.sh image.txt cn-hangzhou myproject acr

# Tencent Cloud TCR
./sync.sh image.txt "" myproject tcr
```

### GitHub Actions Automation

The project includes GitHub Actions workflows that support:

- Automatic synchronization when pushing to main branch
- Manual trigger synchronization
- Multi-cloud service provider selection
- Custom regions and namespaces

For detailed configuration, see [.github/README.md](.github/README.md).

### Debug Mode

The tool has built-in detailed output information that requires no additional configuration:

```bash
upimage nginx:1.28 --namespace myproject --region cn-north-4
```

Output will include:
- Image parsing information
- Operation step progress
- Error details (if any)
- Success confirmation information

## Development

### Run Tests

```bash
go test ./...
```

### Build

```bash
# Build executable for current platform
make bin

# Or use Go command
go build -o upimage main.go
```

### Cross-Platform Build

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o upimage-linux main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o upimage.exe main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o upimage-darwin main.go
```

## Troubleshooting

### Common Issues

1. **Huawei Cloud SWR Authentication Failure**
   ```
   Error initializing Up: failed to create credentials
   ```
   Solution: Check that `HUAWEICLOUD_SDK_AK` and `HUAWEICLOUD_SDK_SK` environment variables are correctly set.

2. **Alibaba Cloud/Tencent Cloud Authentication Failure**
   ```
   Error executing upimage: failed to push image
   ```
   Solution: Ensure you have logged in to the corresponding cloud service provider using `docker login` command.

3. **Unsupported Cloud Service Provider**
   ```
   Invalid registry: xxx. Supported values are 'swr', 'acr' or 'tcr'.
   ```
   Solution: Use supported cloud service provider parameters: `swr` (Huawei Cloud), `acr` (Alibaba Cloud), `tcr` (Tencent Cloud).

4. **Docker Command Failure**
   ```
   Error executing Up: exit status 1
   ```
   Solution: Ensure Docker daemon is running and has permission to access the specified image.

5. **Network Connection Issues**
   - Ensure network connection is normal
   - Check if proxy configuration is needed
   - Verify that the corresponding cloud service is available in the current region

6. **Permission Issues**
   ```
   Error: permission denied
   ```
   Solution: Ensure the current user has Docker operation permissions, may need to add user to docker group.

7. **Namespace Does Not Exist**
   ```
   failed to push image: namespace not found
   ```
   Solution:
   - Huawei Cloud SWR: Enable `--create-namespace` parameter (enabled by default)
   - Alibaba Cloud ACR/Tencent Cloud TCR: Pre-create namespace in corresponding console

5. **Image Does Not Exist**
   ```
   Error: pull access denied for image
   ```
   Solution: Check if the image name is correct and ensure you have permission to access the image.

6. **Region Not Supported**
   ```
   Error: region not supported
   ```
   Solution: Check if the specified region is a valid cloud service provider region code.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork this repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please ensure your code follows the project's coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Huawei Cloud team for providing excellent SDK support
- Thanks to the Cobra team for the powerful CLI framework
- Thanks to all contributors who helped improve this project

## Support

If you encounter any issues or have questions, please:

1. Check the [FAQ](#troubleshooting) section first
2. Search existing [Issues](https://github.com/hiyongliz/upimage/issues)
3. Create a new issue with detailed information
4. Join our community discussions

---

**Made with ❤️ by the UpImage team**
