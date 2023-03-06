go-gen: #oapi-codegenでgoのコードを生成する
	oapi-codegen -package main -generate types openapi.yaml > go_api/types_gen.go
	oapi-codegen -package main -generate server openapi.yaml > go_api/server_gen.go
