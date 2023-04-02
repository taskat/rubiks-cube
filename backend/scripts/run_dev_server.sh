#! /bin/bash
# Set stopping on error
set -e
# Get scripts folder
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Compile the source codes
go build -o dev_mode.exe ./scripts/devmode/devmode.go
# Run it
./dev_mode.exe $@
# Remove executable
rm dev_mode.exe
# Unset stopping on error
set +e