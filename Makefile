include .env

migrateup:
	migrate  -path db/migrations -database "postgres://${POSTGRES_USER}:$(POSTGRES_PASSWORD)@localhost:5432/${POSTGRES_DB}?sslmode=disable" -verbose up



migratedown:
	migrate  -path db/migrations -database "postgres://${POSTGRES_USER}:$(POSTGRES_PASSWORD)@localhost:5432/${POSTGRES_DB}?sslmode=disable" -verbose down
	


