.PHONY: generate_ts_code generate_go_code

# Paths
BACKEND_SRC_DIR = ./backend/src
FRONTEND_SRC_DIR = ./frontend/src
INCLUDE_DIR = googleapis
PROTOC_GEN_TS_PATH = $(shell pwd)/../frontend/node_modules/.bin/protoc-gen-ts
OUT_DIR = ./backend/src/chat_server/gen
GO_OUT_DIR = ./backend/src/chat_server/gen
TS_OUT_DIR = ./frontend/src/gen

generate_go_code:
	pwd
	@protoc --proto_path=$(BACKEND_SRC_DIR)/protos --go_out=$(GO_OUT_DIR) --go-grpc_out=$(GO_OUT_DIR) $(BACKEND_SRC_DIR)/protos/chat.proto


generate_ts_code:
	@protoc --proto_path=$(BACKEND_SRC_DIR)/protos --proto_path=$(INCLUDE_DIR) --grpc-web_out=import_style=typescript,mode=grpcwebtext:$(TS_OUT_DIR) $(BACKEND_SRC_DIR)/protos/chat.proto

generate_all: generate_go_code generate_ts_code
