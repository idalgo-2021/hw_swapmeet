PROTO_DIR := proto
OUTPUT_DIR := pkg/api/swapmeet_grpc

.PHONY: all
all: generate

.PHONY: generate
generate:
	mkdir -p $(OUTPUT_DIR)
	protoc -I$(PROTO_DIR) \
		--go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(OUTPUT_DIR) --grpc-gateway_opt=paths=source_relative \
		$(PROTO_DIR)/swapmeet.proto

.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR)
