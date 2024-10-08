#!/bin/bash
REPO=vonlatvala/spot-hinta-influxdb
FQT="${REPO}:$(<VERSION)"
docker build . --tag "${FQT}" && docker push "${FQT}"
