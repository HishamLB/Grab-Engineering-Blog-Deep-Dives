#!/usr/bin/env bash
set -euo pipefail

echo "=== Killing all redis-server processes ==="
pkill redis-server || echo "  (none running)"
sleep 1

NODE_DIR="$(cd "$(dirname "$0")" && pwd)/nodes"
if [ -d "$NODE_DIR" ]; then
    echo "=== Cleaning up node data ==="
    rm -rf "$NODE_DIR"/*.aof "$NODE_DIR"/*.rdb "$NODE_DIR"/*.nodes.conf
    echo "  done"
fi

echo "=== Cluster torn down ==="
