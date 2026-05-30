#!/usr/bin/env bash
set -eo pipefail

BASE_DIR="$(cd "$(dirname "$0")" && pwd)"
NODE_DIR="$BASE_DIR/nodes"
mkdir -p "$NODE_DIR"

REDIS_SERVER="/tmp/redis-4.0.9/src/redis-server"

rm -f "$NODE_DIR"/*.aof "$NODE_DIR"/*.rdb "$NODE_DIR"/*.nodes.conf "$NODE_DIR"/*.conf 2>/dev/null || true
rm -rf "$NODE_DIR"/appendonlydir 2>/dev/null || true

for i in $(seq 1 9); do
    PORT=$((6000 + i))
    CONF="$NODE_DIR/node_$PORT.conf"
    cat > "$CONF" <<EOF
port $PORT
cluster-enabled yes
cluster-config-file node_$PORT.nodes.conf
cluster-node-timeout 5000
appendonly yes
appendfilename node_$PORT.aof
dbfilename dump_$PORT.rdb
dir "$NODE_DIR"
daemonize no
EOF
done

echo "=== Starting Redis 4.0.9 nodes ==="
for i in $(seq 1 9); do
    PORT=$((6000 + i))
    CONF="$NODE_DIR/node_$PORT.conf"
    $REDIS_SERVER "$CONF" &
done

sleep 2

echo "=== Creating cluster via Valkey CLI ==="
NODES=""
for i in $(seq 1 9); do
    PORT=$((6000 + i))
    NODES="$NODES 127.0.0.1:$PORT"
done

redis-cli --cluster create $NODES --cluster-replicas 2 --cluster-yes

echo ""
echo "=== Cluster ready ==="
echo "CLI:      redis-cli -c -p 6001"
echo "Kill:     redis-cli -p PORT SHUTDOWN"
