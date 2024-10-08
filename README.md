# spot-hinta-influxdb

## What is this

This project exports data from https://api.spot-hinta.fi into influxdb. It performs
data fetching periodically, controlled by environment variables.

It is written in go, so having either Docker or the golang toolchain installed is
a requirement.

## How do I use it

There are multiple different ways to run this solution. The recommended way
to run it is using `Docker`.

### Using Docker

#### With external influxdb

Modify environment variables in ./docker-compose.yml.example to match the influxdb
instance and database (and user, and password, etc.) you already have. Once the
parameters are correct, it's as simple as executing

```sh
cp docker-compose.yml{.example,} && docker-compose up -d
```

and voilá, today's data should already have found its way to influxdb.

#### Without external influxdb

The ./docker-compose.influx.yml comes preconfigured with an integrated influxdb
instance. You might want to change environment variables in it, but don't have to.

```sh
docker-compose -f ./docker-compose.influx.yml up -d
```

To explore your data in this integrated influxdb instance, execute


```sh
docker exec -it spot-hinta-influxdb_influxdb_1 influx -database spot_price -execute 'SELECT * FROM spot_price;'
```

### Locally

#### Without explicitly compiling it

Please make a copy of `./scripts/run.sh.example` to `./scripts/run.sh`:

```sh
cp ./scripts/run.sh{.example,}
```

Now edit the environment variables found in `./scripts/run.sh`, after which you may run it:

```sh
PWD=spot-hinta-influxdb
./scripts/run.sh
```

#### Explicitly compiling it

##### Inside of docker

```sh
./scripts/build-docker.sh && ./bin/spot-hinta-influxdb
```

##### Outside of docker

```sh
PWD=spot-hinta-influxdb
./scripts/build.sh && ./bin/spot-hinta-influxdb
```

## TODO

* Docs
* Makefile
* GH Actions build pipeline
* GH Actions CI pipeline
* k8s deployment
* GH Actions release pipeline
* GCP function terraforms
* Azure function terraforms
* Systemd unit
* Ansible deployment
* Sonar to CI pipeline
* Tests (tirsk)
* More measurements (ATM only exporting `/today` endpoint, as that's all I'm interested in)

## Author

[Axel Latvala](https://www.linkedin.com/in/axel-latvala)
