# golang.org - Database open-handle

[Link](https://go.dev/doc/database/open-handle)

# Prerequisites

-   Docker

# Setup databases

## MySQL

### Run

```bash
docker run --rm --detach --env-file .env --publish 3306:3306 --name open-handle-mysql kairatngo.jfrog.io/default-docker-virtual/mysql:8.0.27
```

# Watch logs

```bash
docker logs <name> -f
```

# Kill container

```bash
docker kill <name>
```
