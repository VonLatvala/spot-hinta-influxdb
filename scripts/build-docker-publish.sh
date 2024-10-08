#!/bin/bash
REPO=vonlatvala/spot-hinta-influxdb
FQT="${REPO}:v$(<VERSION)"
docker build . --tag "${FQT}" && docker push "${FQT}"
