#!/bin/sh
# Stop on error
set -e
# Compile the source codes
go build -o rubik_server ./src/main.go
# Run it
./rubik_server &
# Get pid
SERVER_PID=$!
# Save pid to file
echo $SERVER_PID > /tmp/server.pid
# Unset stopping on error
set +e