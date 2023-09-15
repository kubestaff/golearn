export DOCKER_IMAGE_DIGEST="$1"
docker compose stop golearn
docker compose pull
docker compose up -d golearn --remove-orphans