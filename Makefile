go-gen: #oapi-codegenでgoのコードを生成する
	oapi-codegen -package entities -generate types openapi.yaml > go_api/domain/entities/todo_gen.go

ent-gen: #entのschemaからコードを生成する
	docker exec -i my-todo-api-1 /bin/bash -c "cd /go/src/app && go generate ./ent"

sql-gen: #atlasでmigrationのコードを生成する
	cd go_api && \
	atlas migrate diff migration_name \
		--dir "file://ent/migrate/migrations" \
		--to "ent://ent/schema" \
		--dev-url "docker://postgres/15/test?search_path=public"
