# StarkNet Stack

## Setup

```bash
sudo apt install -y jq curl net-tools
```

### Install Docker

<https://docs.docker.com/engine/install/ubuntu>

### Install Docker Compose

<https://docs.docker.com/compose/install/linux>

## Run

```bash
docker compose -f dc-l1.yaml up -d
```

```bash
docker compose -f dc-l2.yaml up -d
```

## Stop

```bash
docker compose -f <layer> down
docker compose rm
```

## Tail Logs

```bash
docker container logs -f $(docker ps | grep besu | awk '{print $1}')
docker container logs -f $(docker ps | grep lighthouse | awk '{print $1}')
docker container logs -f $(docker ps | grep pathfinder | awk '{print $1}')
```

## Start Indexer

```bash
./indexer/indexer.sh
```
