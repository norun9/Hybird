migrate_up:
	docker run -i --net host --rm \
		-e DATABASE_URL="tcp:localhost:3306*testdb/root/" \
		-v ${PWD}/db/migrations:/app/db/migrations hybird/goose \
		goose -dir=db/migrations mysql "root:test@/testdb" up

migrate_down:
	docker run -i --net host --rm \
		-e DATABASE_URL="tcp:localhost:3306*testdb/root/" \
		-v ${PWD}/db/migrations:/app/db/migrations hybird/goose \
		goose -dir=db/migrations mysql "root:test@/testdb" down