#!/bin/sh
# Stop on error
set -e
# Get scripts folder
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Build docker image
docker build -t taskat/rubik-server -f dev.Dockerfile .
# Unset stopping on error
set +e