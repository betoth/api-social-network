# Definindo variáveis
PROJECT_NAME = api-social-network
CMD_DIR = cmd/api
SRC_DIR = internal
BIN_DIR = bin

# Comandos
.PHONY: all clean build run test lint format docs install-lint

# Default target
all: clean test lint format run

# Limpar artefatos de compilação
clean:
	rm -rf $(BIN_DIR)/$(PROJECT_NAME)

# Compilar o projeto
build:
	go build -o $(BIN_DIR)/$(PROJECT_NAME) $(CMD_DIR)/main.go

# Executar o projeto
run:
	go run $(CMD_DIR)/main.go

# Rodar testes
test:
	go test ./...

# Instalar golangci-lint se não estiver presente
install-lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo >&2 "Installing golangci-lint..."; go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; }

# Verificar linting do código
lint: install-lint
	golangci-lint run

# Formatar o código
format:
	go fmt ./...

# Gerar documentação (Swagger)
docs:
	swag init -g $(CMD_DIR)/main.go

# Tarefas auxiliares
help:
	@echo "Comandos disponíveis:"
	@echo "  make all      - Limpar, testar, verificar lint, formatar e executar o projeto"
	@echo "  make clean    - Limpar artefatos de compilação"
	@echo "  make build    - Compilar o projeto"
	@echo "  make run      - Executar o projeto"
	@echo "  make test     - Rodar testes"
	@echo "  make lint     - Verificar linting do código"
	@echo "  make format   - Formatar o código"
	@echo "  make docs     - Gerar documentação (Swagger)"
	@echo "  make help     - Mostrar esta mensagem de ajuda"
