# Service Discovery

Docs for tools used
- [Consul](https://www.consul.io/docs/index.html)
- [Registrator](http://gliderlabs.github.io/registrator/latest/)
- [Nginx](https://nginx.org/en/docs/)
- [Consul-Template](https://github.com/hashicorp/consul-template)
- [Docker + Docker Compose](https://docs.docker.com/compose/compose-file/)


# Quickstart

## Manage
#### Build

`docker-compose build --no-cache --force-rm`

#### Start, or recreated changed containers

`docker-compose up -d`

#### Tail logs

`docker-compose logs -f`

#### Stop

`docker-compose down`

#### Scale

`docker-compose up -d --scale foo=3 --scale bar=3`


## Access

- Consul: http://localhost:8500
- foo: http://localhost:8080/foo
- bar: http://localhost:8080/bar (should be dead)
