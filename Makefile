PROTO_DIR=proto
PROTO_FILES=$(PROTO_DIR)/stock.proto
OUT_DIR=$(PROTO_DIR)/generated

proto:
	protoc \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)