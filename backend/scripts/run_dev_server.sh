#! /bin/bash
# Get scripts folder
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Compile the source codes
go build -o dev_mode.exe ./scripts/devmode/devmode.go ./scripts/devmode/watcher.go ./scripts/devmode/restarter.go
# Run it
./dev_mode.exe $@
# Remove executable
rm dev_mode.exe