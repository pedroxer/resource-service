generateGRPC:
	protoc --go_out=./proto_gen --go_opt=paths=source_relative --go_grpc_out=./proto_gen --go_grpc_opt=paths=source_relative ./internal/resource.proto
