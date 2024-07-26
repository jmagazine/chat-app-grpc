.PHONY: generate_grpc_code

generate_grpc_code:
	pwd
	protoc --go_out=. --go_opt=paths=source_relative \
	       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	       chat\chat.proto
	protoc --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. chat/chat.proto