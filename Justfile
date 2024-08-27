migrate_up:
	docker run -i --net host --rm \
		-e DATABASE_URL="tcp:localhost:3306*testdb/hybird/" \
		-v ${PWD}/db/migrations:/app/db/migrations hybird/goose \
		goose -dir=db/migrations mysql "hybird:test@/testdb" up

migrate_down:
	docker run -i --net host --rm \
		-e DATABASE_URL="tcp:localhost:3306*testdb/hybird/" \
		-v ${PWD}/db/migrations:/app/db/migrations hybird/goose \
		goose -dir=db/migrations mysql "hybird:test@/testdb" down

client build="":
  docker compose -f docker-compose-client.yml up {{ build }}

air build="":
  docker compose up {{ build }}