postgres:
	sudo docker run --name postgres -p 1000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres
createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root SimpleBank
dropdb:
	sudo docker exec -it postgres dropdb --username=root SimpleBank
migrate:
	migrate -path db/migration -database "postgresql://root:root@localhost:1000/SimpleBank?sslmode=disable" -verbose up
downgrade:
	migrate -path db/migration -database "postgresql://root:root@localhost:1000/SimpleBank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.PHONY:postgres createdb dropdb migrate downgrade sqlc test