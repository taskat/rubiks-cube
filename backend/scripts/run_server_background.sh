#! /bin/bash
# Stop on error
set -e
# Get scripts folder"
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Compile the source codes
go build -o rubik_server.exe src/main.go
# Run it
./rubik_server.exe &
# Get pid
PID=$!
echo "Server started with PID: $PID"
# Save pid to file
echo $PID > rubik_server.pid
# Unset stopping on error
set +e