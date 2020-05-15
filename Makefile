rbacgen-create:
	go run cmd/rbacgen/*.go --create

rbacgen-fill:
	go run cmd/rbacgen/*.go --fill

dev-schema-up:
	migrate -source file://./pkg/rbac/repository/postgres/migrations -database postgres://postgres:superuser@localhost:5432/roles?sslmode=disable up

dev-schema-down:
	migrate -source file://./pkg/rbac/repository/postgres/migrations -database postgres://postgres:superuser@localhost:5432/roles?sslmode=disable down