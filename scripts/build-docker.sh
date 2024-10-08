#!/bin/bash
docker run --rm \
    -v ./src:/tmp/sln \
    -v ./bin:/tmp/dockerdist \
    --entrypoint go \
    -u "$UID:$GID" \
    --workdir /tmp/sln \
	-e GOCACHE=/tmp/.cache \
    golang:1.23-bookworm \
    build -C /tmp/sln -o /tmp/dockerdist/ .
