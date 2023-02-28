postgres:
	docker run --name web-chat -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=qwerty -d postgres

createdb:
	docker exec -it web-chat  createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it web-chat dropdb go-chat

postgresin:
	docker exec -it web-chat psql

migrateup:
	migrate -path server/db/migration -database postgresql://root:qwerty@localhost:5432/go-chat?sslmode=disable -verbose up

migratedown:
	migrate -path server/db/migration -database postgresql://root:qwerty@localhost:5432/go-chat?sslmode=disable -verbose down

.PHONY: postgres, createdb, postgresin, dropdb, migrateup, migratedown


