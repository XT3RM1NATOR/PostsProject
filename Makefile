post_proto:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=require_unimplemented_servers=true:. --go-grpc_opt=paths=source_relative     protos/post_service/post_service.proto

user_proto:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=require_unimplemented_servers=true:. --go-grpc_opt=paths=source_relative     protos/user_service/user_service.proto

