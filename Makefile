rbacgen-create:
	go run cmd/rbacgen/main.go --create --specs cmd/rbacgen/timetableApi.yaml --tmpl cmd/rbacgen/roleTmpl.yaml

rbacgen-fill:
	go run cmd/rbacgen/main.go --fill --tmpl cmd/rbacgen/roleTmpl.yaml

dev-schema-up:
	migrate -source file://./pkg/rbac/repository/postgres/migrations -database postgres://postgres:superuser@localhost:5432/roles?sslmode=disable up

dev-schema-down:
	migrate -source file://./pkg/rbac/repository/postgres/migrations -database postgres://postgres:superuser@localhost:5432/roles?sslmode=disable down