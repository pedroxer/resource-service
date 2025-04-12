generateGRPC:
	protoc --go_out=./proto_gen --go_opt=paths=source_relative --go_grpc_out=./proto_gen --go_grpc_opt=paths=source_relative ./internal/resource.proto

runDocker:
	docker build -t resource-service .
	docker run --env-file .env -p 8081:8081 resource-service
