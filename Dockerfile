FROM golang:1.23-bookworm AS builder

RUN mkdir /sln

COPY ./src/spot-hinta-influxdb/ /sln

WORKDIR /sln

RUN go install && go build

FROM debian:bookworm

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends ca-certificates && apt-get clean

RUN mkdir /sln

COPY --from=builder /sln/spot-hinta-influxdb /sln/

WORKDIR /sln

ENTRYPOINT ["/sln/spot-hinta-influxdb"]
