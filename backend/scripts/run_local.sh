#! /bin/bash
# Get scripts folder
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Compile the source codes
go build -o rubik_server.exe src/main.go
# Run it
./rubik_server.exe $1
# Remove executable
rm rubik_server.exe