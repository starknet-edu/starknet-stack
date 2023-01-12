# StarkNet Stack

## Setup

### Install Docker && Docker Compose

<https://docs.docker.com/engine/install/ubuntu>
<https://docs.docker.com/compose/install/linux>

### Reset

```bash
docker rm -f $(docker ps -a -q)
docker volume rm $(docker volume ls -q)
```

## Run

```bash
docker compose up
```

## Stop

```bash
docker compose down
docker compose rm
```
