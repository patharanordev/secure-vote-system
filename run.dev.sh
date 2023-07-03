docker-compose -f docker-compose.dev.yml down -v && \

mkdir -p pgdata && \
docker-compose -f docker-compose.dev.yml up --build