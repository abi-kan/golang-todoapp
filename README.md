# golang-todoapp
This project is based on [this tutorial](https://youtu.be/uxzZJw-52qk)

## Prerequisites

- Docker with Docker Compose
- `make`
- Go 1.26+

## Setup

The Makefile includes `.env`, so create it before running any target
(it is git-ignored):

```bash
cp .env.example .env
```

## Common workflows

### Run everything in Docker

```bash
make env-up       # start Postgres and Redis
make migrate-up   # apply migrations
make deploy       # build and start the app
```

Stop with:

```bash
make undeploy
make env-down
```

### Run the app locally (dependencies in Docker)

```bash
make env-up             # start Postgres and Redis
make migrate-up         # apply migrations
make env-port-forward   # expose Postgres (5432) and Redis (6379) on localhost
make run                # go fmt, go vet, then go run
```

When finished:

```bash
make env-down
```

### Change the database schema

```bash
make migrate-create seq=add_tasks_priority   # creates an up/down pair in migrations/
# edit the new .up.sql and .down.sql files
make migrate-up                              # apply
make migrate-down                            # roll back
```

### Update the API docs

After changing handler annotations or DTOs:

```bash
make swagger-gen   # regenerates docs/
```

### Reset local data

```bash
make env-cleanup    # deletes out/pgdata and out/redisdata (asks for confirmation)
make logs-cleanup   # deletes out/logs (asks for confirmation)
```

## Makefile reference

### Environment (Postgres and Redis)

| Target             | Description                                                         |
|--------------------|---------------------------------------------------------------------|
| `env-up`           | Start Postgres and Redis containers                                 |
| `env-down`         | Stop Postgres, Redis and the port forwarders                        |
| `env-cleanup`      | Delete Postgres and Redis data in `out/` (asks first)               |
| `env-port-forward` | Expose Postgres and Redis on `127.0.0.1` (needed for `make run`)    |
| `env-port-close`   | Close the forwarded ports                                           |

### Migrations

| Target                    | Description                                                  |
|---------------------------|--------------------------------------------------------------|
| `migrate-up`              | Apply all pending migrations                                 |
| `migrate-down`            | Roll back migrations                                         |
| `migrate-create seq=name` | Create a new migration pair (`seq` is required)              |

### Application

| Target      | Description                                                      |
|-------------|------------------------------------------------------------------|
| `run`       | Run locally after `fmt` and `vet`; sets hosts to `localhost` and logs to `out/logs` |
| `deploy`    | Build the image and start the app container                      |
| `undeploy`  | Stop the app container                                           |

### Code quality and docs

| Target         | Description                                  |
|----------------|----------------------------------------------|
| `fmt`          | `go fmt ./...`                               |
| `vet`          | `go vet ./...`                               |
| `staticcheck`  | `staticcheck ./...` (must be installed)      |
| `swagger-gen`  | Regenerate Swagger docs in `docs/`           |

### Logs

| Target         | Description                                  |
|----------------|----------------------------------------------|
| `logs-cleanup` | Delete `out/logs` (asks first)               |

## Notes

- Everything generated at runtime (database files, Redis dumps, logs) lives in
  `out/`, which is git-ignored. Do not commit or share it.
- `make run` needs the port forwarders (`make env-port-forward`), because the
  app runs on your host and talks to the containers via `localhost`.