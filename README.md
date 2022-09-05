# Nidus API

Home monitoring API part of the
[Nidus](https://github.com/alexandrelamberty/nidus) project.

## Features

- [ ] Manage devices, zones, capabilities
- [ ] Pair device
- [ ] Create alerts and notifications

## Technolgies and frameworks

- Docker
- MongoDB
- Fiber

## Usage

```bash
go run cmd/api/main
```

## Build and run with Docker

Build the image, see: [Dockerfile](./Dockerfile)

```bash
docker build -t alexandrelamberty/nidus-api
```

Run the image, we specify the ports mapping, environment variables file and network to join

```bash
docker run -p 8080:8080 --network=nidus-server_default --env-file .env -name nidus-api -d alexandrelamberty/nidus-api_latest
```

## References

- <https://github.com/gofiber/recipes/tree/master/clean-architecture>
- <https://docs.docker.com/language/golang/>