#!/bin/sh
# Stop on error
set -e
# Get scripts folder
SCRIPTS="$(dirname $(realpath "$0"))"
cd ${SCRIPTS}/..
# Run docker container
docker run -v D:/egyetem/msc/dipterv/rubiks-cube/backend:/app/dev/ -p 8080:8080 taskat/rubik-server:latest
# Unset stopping on error
set +e