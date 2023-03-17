help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

up: ## dev server立ち上げ
	docker-compose up -d

down: ## dev server停止
	docker-compose down

test: ## test
	docker exec -i my-todo-api-1 /bin/bash -c "cd /go/src/app && go test -v ./..."
	docker exec -i my-todo-web-1 /bin/bash -c "npm test -- --coverage --watchAll=false"

go-gen: ## oapi-codegenでgoのコードを生成する
	oapi-codegen -package entities -generate types openapi.yaml > go_api/domain/entities/todo_gen.go

axios-gen: ## OpenAPI Generator TypeScript Axiosで型付きリクエストの自動生成 
	openapi-generator-cli generate -g typescript-axios -i ./openapi.yaml -o ./web/types/typescript-axios

ent-gen: ## entのschemaからコードを生成する
	docker exec -i my-todo-api-1 /bin/bash -c "cd /go/src/app && go generate ./ent"

sql-gen: ## atlasでmigrationのコードを生成する
	cd go_api && \
	atlas migrate diff migration_name \
	--dir "file://ent/migrate/migrations" \
	--to "ent://ent/schema" \
	--dev-url "docker://postgres/14/test?search_path=public"

install-pgo: ## pgoをインストール
	git clone https://github.com/CrunchyData/postgres-operator-examples || /bin/true
	cd postgres-operator-examples && \
	kubectl apply -k kustomize/install/namespace && \
	kubectl apply --server-side -k kustomize/install/default

docker-build-and-push: ## docker-build-and-push
	docker build -t ${DOCKERHUB_USERNAME}/my-todo-go-api:${GITHUB_SHA} ./go_api
	docker push ${DOCKERHUB_USERNAME}/my-todo-go-api:${GITHUB_SHA}
	docker build -t ${DOCKERHUB_USERNAME}/my-todo-web:${GITHUB_SHA} ./web
	docker push ${DOCKERHUB_USERNAME}/my-todo-web:${GITHUB_SHA}

deploy: ## kubernetes clusterにデプロイする
	@cat manifests/web.yaml | envsubst '$${DOCKERHUB_USERNAME} $${GITHUB_SHA}' | kubectl apply -f -
	@cat manifests/api.yaml | envsubst '$${DOCKERHUB_USERNAME} $${GITHUB_SHA}' | kubectl apply -f -
	kubectl apply -f manifests/ingress.yaml 
