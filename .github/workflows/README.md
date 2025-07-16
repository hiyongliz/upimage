# GitHub Actions Configuration

This document explains how to set up GitHub Actions for automatic Docker image synchronization to Huawei Cloud SWR.

## Required Secrets

Configure the following secrets in your GitHub repository settings:

### Huawei Cloud API Credentials
- `HUAWEICLOUD_SDK_AK`: Your Huawei Cloud Access Key
- `HUAWEICLOUD_SDK_SK`: Your Huawei Cloud Secret Key

### SWR Registry Credentials (Optional)
If you need to pull from private registries:
- `SWR_REGISTRY_USER`: SWR registry username
- `SWR_REGISTRY_PASSWORD`: SWR registry password

## Workflow Triggers

The sync workflow can be triggered in several ways:

### 1. Automatic Triggers
- **Push to main branch**: When `image.txt` or the workflow file is modified
- **File changes**: Only when relevant files are changed to avoid unnecessary runs

### 2. Manual Trigger
- **Workflow Dispatch**: Manually trigger from GitHub Actions tab with custom parameters:
  - Target region selection
  - Custom image file path

## Workflow Jobs

### 1. Validate Job
- Checks if the image list file exists
- Counts valid images (non-empty, non-comment lines)
- Fails early if no valid images are found

### 2. Sync Job
- Sets up Go environment with caching
- Runs tests to ensure code quality
- Builds the upimage binary
- Authenticates with registries
- Executes the batch sync process
- Uploads logs as artifacts

### 3. Notify Job
- Provides clear success/failure notifications
- Includes image count and target region information

## Image List Format

The `image.txt` file supports:

```text
# Comments start with # and are ignored
# Empty lines are also ignored

# Basic image names
nginx:latest
redis:7.0

# Images with registry
docker.io/library/alpine:latest
quay.io/prometheus/prometheus:latest
```

## Monitoring and Debugging

### Logs
- All sync operations are logged with color-coded output
- Failed images are clearly listed in the summary
- Logs are uploaded as artifacts for later analysis

### Artifacts
- Sync logs are retained for 7 days
- Download from the Actions run page for troubleshooting

### Error Handling
- Early validation prevents wasted CI time
- Clear error messages for common issues
- Graceful handling of partial failures

## Best Practices

1. **Test Locally First**: Always test your image list with `./sync.sh` locally before pushing
2. **Small Batches**: Keep image lists manageable (< 50 images per run)
3. **Monitor Usage**: Check your Huawei Cloud SWR quotas and usage
4. **Security**: Never commit credentials to the repository
5. **Regions**: Use the closest region for better performance

## Troubleshooting

### Common Issues

1. **Missing Secrets**
   ```
   Error: HUAWEICLOUD_SDK_AK environment variable is not set
   ```
   Solution: Add the required secrets in repository settings

2. **Invalid Region**
   ```
   Error: region not supported
   ```
   Solution: Use a valid Huawei Cloud region code

3. **Authentication Failed**
   ```
   Error: failed to create credentials
   ```
   Solution: Verify your AK/SK credentials are correct

4. **Docker Pull Failed**
   ```
   Error: pull access denied
   ```
   Solution: Ensure the image exists and you have access permissions

### Getting Help

- Check the Actions logs for detailed error messages
- Review the sync summary for failed images
- Download artifacts for offline analysis
- Create an issue if you encounter persistent problems
