get_tools:
	protoc -I ../chat -I ../googleapis \
  --grpc-gateway_out ../chat_server \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  ../chat/chat.proto