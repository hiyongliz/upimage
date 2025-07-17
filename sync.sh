#!/bin/bash

# sync.sh - Batch sync Docker images to Huawei Cloud SWR
# Usage: ./sync.sh [image_list_file] [region] [namespace]

set -euo pipefail

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Default values
IMAGE_FILE="${1:-image.txt}"
REGION="${2:-cn-south-1}"
NAMESPACE="${3:-default}"
UPIMAGE_BINARY="./upimage"
REGISTRY="${4:-swr}" # Default to SWR, can be acr or tcr

# Logging functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if upimage binary exists
check_binary() {
    if [[ ! -f "$UPIMAGE_BINARY" ]]; then
        log_error "upimage binary not found at $UPIMAGE_BINARY"
        log_info "Please run 'make build' to build the binary first"
        exit 1
    fi

    if [[ ! -x "$UPIMAGE_BINARY" ]]; then
        log_error "upimage binary is not executable"
        log_info "Please run 'chmod +x $UPIMAGE_BINARY'"
        exit 1
    fi
}

# Check if image list file exists
check_image_file() {
    if [[ ! -f "$IMAGE_FILE" ]]; then
        log_error "Image list file not found: $IMAGE_FILE"
        log_info "Please create the file with one image per line"
        exit 1
    fi
}

# Check required environment variables
check_env() {
    if [[ "${REGISTRY}" != "swr" ]]; then
        return
    fi

    if [[ -z "${HUAWEICLOUD_SDK_AK:-}" ]]; then
        log_error "HUAWEICLOUD_SDK_AK environment variable is not set"
        exit 1
    fi

    if [[ -z "${HUAWEICLOUD_SDK_SK:-}" ]]; then
        log_error "HUAWEICLOUD_SDK_SK environment variable is not set"
        exit 1
    fi
}

# Process images
sync_images() {
    local total_count=0
    local success_count=0
    local failed_images=()

    log_info "Starting image sync to region: $REGION"
    log_info "Target namespace: $NAMESPACE"
    log_info "Reading images from: $IMAGE_FILE"
    echo

    while IFS= read -r line || [[ -n "$line" ]]; do
        # Skip empty lines and comments
        if [[ -z "$line" ]] || [[ "$line" =~ ^[[:space:]]*# ]]; then
            continue
        fi

        # Trim whitespace
        image=$(echo "$line" | sed 's/^[[:space:]]*//;s/[[:space:]]*$//')
        
        if [[ -z "$image" ]]; then
            continue
        fi

        total_count=$((total_count + 1))
        log_info "Processing image [$total_count]: $image"

        # Run upimage with error handling
        if $UPIMAGE_BINARY "$image" --region "$REGION" --namespace "$NAMESPACE" --registry "$REGISTRY"; then
            success_count=$((success_count + 1))
            log_success "Successfully synced: $image"
        else
            log_error "Failed to sync: $image"
            failed_images+=("$image")
        fi
        echo
    done < "$IMAGE_FILE"

    # Summary
    echo "========================================="
    log_info "Sync Summary:"
    log_info "Total images: $total_count"
    log_success "Successful: $success_count"
    
    if [[ ${#failed_images[@]} -gt 0 ]]; then
        log_error "Failed: ${#failed_images[@]}"
        log_error "Failed images:"
        for img in "${failed_images[@]}"; do
            echo "  - $img"
        done
        exit 1
    else
        log_success "All images synced successfully!"
    fi
}

# Show usage
show_usage() {
    echo "Usage: $0 [image_list_file] [region] [namespace]"
    echo
    echo "Arguments:"
    echo "  image_list_file  Path to file containing list of images (default: image.txt)"
    echo "  region          Huawei Cloud region (default: cn-south-1)"
    echo "  namespace       Target namespace (default: default)"
    echo
    echo "Environment variables required:"
    echo "  HUAWEICLOUD_SDK_AK  Huawei Cloud Access Key"
    echo "  HUAWEICLOUD_SDK_SK  Huawei Cloud Secret Key"
    echo
    echo "Examples:"
    echo "  $0                                    # Use default image.txt, cn-south-1, default"
    echo "  $0 my-images.txt                      # Use custom image file"
    echo "  $0 my-images.txt cn-north-4           # Use custom file and region"
    echo "  $0 my-images.txt cn-north-4 myproject # Use custom file, region, and namespace"
}

# Main function
main() {
    # Handle help flag
    if [[ "${1:-}" == "-h" ]] || [[ "${1:-}" == "--help" ]]; then
        show_usage
        exit 0
    fi

    log_info "Starting upimage batch sync tool"
    
    # Run checks
    check_binary
    check_image_file
    check_env
    
    # Start sync process
    sync_images
}

# Run main function
main "$@"
