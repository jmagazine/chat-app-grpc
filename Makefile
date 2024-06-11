.PHONY: generate_grpc_code

generate_grpc_code:
	pwd
	protoc --go_out=. --go_opt=paths=source_relative \
	       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	       src/chat/chat.proto
