docker compose -f infra/docker-compose.yaml stop golearn
docker compose -f infra/docker-compose.yaml up --build golearn -d
docker compose -f infra/docker-compose.yaml up -d --remove-orphans