help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

up: ## dev server立ち上げ
	docker-compose up -d

down: ## dev server停止
	docker-compose down

test: ## test
	docker exec  -i my-todo-api-1 /bin/bash -c "cd /go/src/app && go test ./..."

go-gen: ## oapi-codegenでgoのコードを生成する
	oapi-codegen -package entities -generate types openapi.yaml > go_api/domain/entities/todo_gen.go

ent-gen: ## entのschemaからコードを生成する
	docker exec -i my-todo-api-1 /bin/bash -c "cd /go/src/app && go generate ./ent"

sql-gen: ## atlasでmigrationのコードを生成する
	cd go_api && \
	atlas migrate diff migration_name \
	--dir "file://ent/migrate/migrations" \
	--to "ent://ent/schema" \
	--dev-url "docker://postgres/15/test?search_path=public"

install-pgo: ## pgoをインストール
	git clone https://github.com/CrunchyData/postgres-operator-examples || /bin/true
	cd postgres-operator-examples && \
	kubectl apply -k kustomize/install/namespace && \
	kubectl apply --server-side -k kustomize/install/default

docker-build-and-push: ## docker-build-and-push
	docker build -t hdys/my-todo-go-api:latest ./go_api
	docker push hdys/my-todo-go-api
	docker build -t hdys/my-todo-web:latest ./web
	docker push hdys/my-todo-web
