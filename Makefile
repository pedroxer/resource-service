generateGRPC:
	protoc --go_out=./internal/proto_gen --go_opt=paths=source_relative --go_grpc_out=./internal/proto_gen --go_grpc_opt=paths=source_relative ./docs/protos/resource.proto
