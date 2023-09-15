export DOCKER_IMAGE_DIGEST="$1"
docker compose -f infra/docker-compose.yaml stop golearn
docker compose -f infra/docker-compose.yaml pull
docker compose -f infra/docker-compose.yaml up -d golearn --remove-orphans