PROTO_DIR=proto
PB_DIR=pkg/pb

proto:
	protoc -I=$(PROTO_DIR) \
		--go_out=$(PB_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PB_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/node.proto

run-master:
	go run ./cmd/master

run-worker:
	go run ./cmd/worker
