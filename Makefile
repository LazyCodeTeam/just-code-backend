run:
	go run cmd/justcode/main.go

test:
	go test ./...

dev_infra:
	cd deployments/infrastructure/dev && terraform init && terraform apply

dev_outputs:
	cd deployments/infrastructure/dev && terraform output all_outputs

spec:
	swagger generate spec -o ./api/swagger.yaml --scan-models

migrate_up:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path db/migration up
	
migrate_down:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path db/migration down -all

sqlc:
	sqlc generate
