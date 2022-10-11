# StarkNet Stack

## Setup

### Install Docker && Docker Compose

<https://docs.docker.com/engine/install/ubuntu>

```bash
export UID=$(id -u)
export GID=$(id -g)
```

### Reset

```bash
docker rm -f $(docker ps -a -q)
docker volume rm $(docker volume ls -q)
```
