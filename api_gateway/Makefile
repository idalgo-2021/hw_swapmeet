PROTO_DIR := proto
OUTPUT_DIR := pkg/api

PROTO_FILES := auth.proto swapmeet.proto

.PHONY: all generate clean

all: generate

generate:
	@mkdir -p $(OUTPUT_DIR)
	@for proto in $(PROTO_FILES); do \
		proto_name=$${proto%.*}; \
		mkdir -p $(OUTPUT_DIR)/$$proto_name; \
		protoc -I$(PROTO_DIR) \
			--go_out=$(OUTPUT_DIR)/$$proto_name --go_opt=paths=source_relative \
			--go-grpc_out=$(OUTPUT_DIR)/$$proto_name --go-grpc_opt=paths=source_relative \
			$(PROTO_DIR)/$$proto; \
	done

clean:
	rm -rf $(OUTPUT_DIR)