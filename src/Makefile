.PHONY: generate_grpc_code

generate_grpc_code:
	echo %cd%
	protoc --go_out=. --go_opt=paths=source_relative \
	       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	       chat/chat.proto
