#!/bin/bash

# Quick test script for a single MySQL version
# Usage: ./quick-test.sh [5.7|8.0|8.4]

set -e

VERSION=${1:-8.0}
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
DBMIX_BIN="../src/dbmix"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

main() {
    # Map versions to ports
    case $VERSION in
        5.7)
            PORT=3357
            ;;
        8.0)
            PORT=3380
            ;;
        8.4)
            PORT=3384
            ;;
        *)
            print_error "Unsupported MySQL version: $VERSION"
            echo "Supported versions: 5.7, 8.0, 8.4"
            exit 1
            ;;
    esac

    print_status "Quick test for MySQL $VERSION"

    # Check if dbmix binary exists
    if [[ ! -f "$DBMIX_BIN" ]]; then
        print_status "Building dbmix..."
        cd ../src && go build -o dbmix . && cd ../test_containers
    fi

    # Start the MySQL container
    print_status "Starting MySQL $VERSION container..."
    cd "$SCRIPT_DIR/mysql-$VERSION"

    if ! ./run.sh start; then
        print_error "Failed to start MySQL $VERSION container"
        exit 1
    fi

    print_success "MySQL $VERSION container started"

    # Wait for container to be fully ready
    print_status "Waiting for MySQL to be fully ready..."
    sleep 5

    # Run dbmix test
    print_status "Running dbmix test..."
    cd "$SCRIPT_DIR"

    if $DBMIX_BIN -host=localhost -port=$PORT -user=root -password=rootpass -database=testdb -outfile="quick-test-mysql$VERSION"; then
        print_success "dbmix test completed successfully"

        # Show output file info
        if [[ -f "quick-test-mysql$VERSION.md" ]]; then
            local size=$(stat -f%z "quick-test-mysql$VERSION.md" 2>/dev/null || stat -c%s "quick-test-mysql$VERSION.md" 2>/dev/null || echo "0")
            print_success "Output file created: quick-test-mysql$VERSION.md (${size} bytes)"
            print_status "First few lines of output:"
            head -20 "quick-test-mysql$VERSION.md"
        fi
    else
        print_error "dbmix test failed"
        cd "$SCRIPT_DIR/mysql-$VERSION"
        ./run.sh stop
        exit 1
    fi

    # Stop the container
    print_status "Stopping MySQL $VERSION container..."
    cd "$SCRIPT_DIR/mysql-$VERSION"
    ./run.sh stop

    print_success "Quick test completed for MySQL $VERSION"
}

main