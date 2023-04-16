#!/bin/sh
# Stop on error
set -e
kill $1
# Unset stopping on error
set +e