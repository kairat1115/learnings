# How to do DDD in GO

[Link](https://habr.com/ru/company/domclick/blog/592087/)

# Repositories

## Memory

No need to run anything

## MongoDB

### Prerequisites

-   Docker

### Run

```bash
docker run --rm --detach --publish 27017:27017 --name tavern-mongodb kairatngo.jfrog.io/default-docker-virtual/mongo:5.0.5-focal
```
