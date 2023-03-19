#! /bin/bash
# Get scripts folder
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Compile the source codes
go build src/main.go
# Run it
./main.exe $1
# Remove executable
rm main.exe