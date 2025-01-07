PROTO_DIR := proto
OUTPUT_DIR := pkg/api/auth_grpc

.PHONY: check-plugins
check-plugins:
	@if ! [ -x "$$(which protoc-gen-go)" ]; then \
		echo "protoc-gen-go не найден. Установите его командой:"; \
		echo "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; \
		exit 1; \
	fi
	@if ! [ -x "$$(which protoc-gen-go-grpc)" ]; then \
		echo "protoc-gen-go-grpc не найден. Установите его командой:"; \
		echo "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; \
		exit 1; \
	fi
	@if ! [ -x "$$(which protoc-gen-grpc-gateway)" ]; then \
		echo "protoc-gen-grpc-gateway не найден. Установите его командой:"; \
		echo "go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest"; \
		exit 1; \
	fi

.PHONY: all
all: check-plugins generate

.PHONY: generate
generate:
	@echo "Генерация файлов из $(PROTO_DIR)/auth.proto в $(OUTPUT_DIR)..."
	mkdir -p $(OUTPUT_DIR)
	protoc -I$(PROTO_DIR) \
		--go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(OUTPUT_DIR) --grpc-gateway_opt=paths=source_relative \
		$(PROTO_DIR)/auth.proto

.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR)
