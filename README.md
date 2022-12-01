[![Go](https://github.com/alexandrelamberty/nidus-api/actions/workflows/go.yml/badge.svg)](https://github.com/alexandrelamberty/nidus-api/actions/workflows/go.yml)
[![Docker](https://github.com/alexandrelamberty/nidus-api/actions/workflows/docker.yml/badge.svg)](https://github.com/alexandrelamberty/nidus-api/actions/workflows/docker.yml)

# Nidus API

Home monitoring API part of the
[Nidus](https://github.com/alexandrelamberty/nidus) project, see the [Nidus API Specification](https://github.com/alexandrelamberty/nidus-api-spec)

## Features

- [x] Manage devices, zones, capabilities
- [x] Pair device
- [ ] Create alerts and notifications
- [ ] Security
  - [ ] Key
- [ ] Tests

## Technolgies and frameworks

- [Docker](https://www.docker.com/)
- [Go](https://go.dev/)
- [Fiber](https://gofiber.io/)

## Usage

This application is part of a Docker stack and rely on a MongoDB database service. see:
[Nidus](https://github.com/alexandrelamberty/nidus) project to launch the
complete stack or only specific services.

### Run with Go

If the database service is up and running, create an .env file and fill it
accordingly with the `database` service configuration.

```properties
ENV=dev
PAIRING_KEY=9fca54477c8ad4e70dc5e1084f884aad
JWT_SECRET=d7a481461577ba4c3c4c6946cca7204b
JWT_EXPIRE=90
BCRYPT_HASH=7f91317e30a02bc7b87205e95b842df2
DATABASE_URI=mongodb://nidus:nidus@localhost:27017/nidus
```

Run the application:

```bash
go run cmd/api/main
```

Go to <http://localhost:3333>

### Tests with Go

> To implement

### Build and run with Docker

Build the image, see: [Dockerfile](./Dockerfile).

```bash
docker build . -t alexandrelamberty/nidus-api:latest
```

Run the image, we specify the ports mapping, environment variables file and
network to join.

```bash
docker run -p 3333:3333 --network=nidus_default --env-file .env --name nidus-api -d alexandrelamberty/nidus-api:latest
```

## Push to Docker Hub

Docker Hub [Nidus API](https://hub.docker.com/repository/docker/alexandrelamberty/nidus-api)

> Automated with GitHub Action, see: [docker.yml](./.github/workflows/docker.yml)

```bash
docker tag alexandrelamberty/nidus-api:latest alexandrelamberty/nidus-api:latest
docker push alexandrelamberty/nidus-api:latest
```

## References

- <https://github.com/gofiber/recipes/tree/master/clean-architecture>
- <https://docs.docker.com/language/golang/>
