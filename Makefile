PROTO_DIR=proto
PROTO_FILES=$(PROTO_DIR)/task.proto

.PHONY: proto

proto:
	protoc --go_out=. --go-grpc_out=. $(PROTO_FILES)
