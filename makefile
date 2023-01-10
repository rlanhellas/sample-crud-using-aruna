.PHONY: help run startinfra test coverage build stopinfra checkinfra swagger
.DEFAULT: help

PROJECT=sample_crud

help:
	@echo "[Makefile] Comandos disponíveis:"
	@echo ""
	@echo "|> COMANDOS"
	@echo ""
	@echo "help - Lista todos os comandos disponíveis"
	@echo "run - Sobe a infra e executa a aplicação"
	@echo "startinfra - Sobe via Docker a infra necessária"
	@echo "stopinfra - Desliga a infra que está rodando no docker"
	@echo "checkinfra - Verifica o status da infra"
	@echo "test - Execute os testes"
	@echo "coverage - Executa os testes e mostra o coverage"

build:
	@echo "iniciando build do projeto ..."
	go mod tidy
	mkdir -p runlocal
	go build -tags=go_json,nomsgpack -o runlocal/main

test: build
	@echo "inicando testes ..."
	go clean -testcache && go test -v -coverpkg=./... -coverprofile=coverage.out ./... | tree coveragereport.out

coverage: test
	@echo "iniciando coverage ..."
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out

startinfra:
	@echo "Subindo infra necessária para o projeto ..."
	docker-compose -f ./deployments/docker-compose.yml -p $(PROJECT) up -d
	@sleep 30

stopinfra:
	@echo "Parando infra ..."
	docker-compose -f ./deployments/docker-compose.yml -p $(PROJECT) down --remove-orphans

checkinfra:
	@echo "Mostrando o status da infra atual ..."
	docker-compose -f ./deployments/docker-compose.yml -p $(PROJECT) ps

swagger:
	go install github.com/swaggo/swag/cmd/swag@v1.8.1
	swag fmt
	swag init --parseDependency --parseDepth=1 --output ./shared

run: swagger build startinfra
	./runlocal/main
