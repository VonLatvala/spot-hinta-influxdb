#!/bin/bash
REPO=vonlatvala/spot-hinta-influxdb
FQT="${REPO}:v$(<VERSION)"
docker tag "$FQT" "$REPO:latest"
