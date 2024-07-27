.PHONY: generate_grpc_code

generate_js_code:
	pwd
	protoc -I=chat/ -I=../../googleapis \
       --js_out=import_style=commonjs,binary:chat_server \
       --grpc-web_out=import_style=commonjs,mode=grpcwebtext:chat_server \
       chat/chat.proto