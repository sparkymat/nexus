all: nexus

nexus:
	CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o nexus main.go

nexus-worker:
 	CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o nexus-worker worker/main.go

start-app:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	# Install godotenv with 'go install github.com/joho/godotenv/cmd/godotenv@latest'
	reflex -s -r '\.go$$' -- godotenv -f .env go run main.go

start-worker:
	reflex -s -r '\.go$$' -- godotenv -f .env go run worker/main.go

start-view:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	# Install templ with 'go install github.com/a-h/templ/cmd/templ@latest'
	reflex -s -r '\.templ$$' -- templ generate

db-migrate:
	migrate -path migrations -database "postgres://127.0.0.1/nexus?sslmode=disable" up

db-schema-dump:
	pg_dump --schema-only -O nexus > internal/database/schema.sql

sqlc-gen:
	sqlc generate

.PHONY: nexus nexus-worker start-app start-worker start-scheduler start-view db-migrate db-schema-dump sqlc-gen
