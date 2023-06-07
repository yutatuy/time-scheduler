
DOCKER_COMPOSE := docker compose
DOCKER_COMPOSE_EXEC := ${DOCKER_COMPOSE} exec
TARGET_MIGRATION_DIR := /app/go/config/goose/migrations
GOOSE_DBSTRING_TEST := 'docker:docker@tcp(test-db:3306)/time_scheduler_test'

help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## ファイルを指定してtestの実行
	${DOCKER_COMPOSE_EXEC} go go test -v $${FILE}

bash: ## goコンテナにアタッチ
	bash: ## goコンテナにアタッチ
	${DOCKER_COMPOSE_EXEC} -it go sh

goose-up: ## DBのup
	${DOCKER_COMPOSE_EXEC} -it go sh -c "cd ${TARGET_MIGRATION_DIR} && goose up"

goose-down: ## DBのdown
	${DOCKER_COMPOSE_EXEC} -it go sh -c "cd ${TARGET_MIGRATION_DIR} && goose down"

goose-reset: ## DBのreset
	${DOCKER_COMPOSE_EXEC} -it go sh -c "cd ${TARGET_MIGRATION_DIR} && goose reset"

goose-up-test: ## test用DBのup
	${DOCKER_COMPOSE_EXEC} -it go sh -c "cd ${TARGET_MIGRATION_DIR} && GOOSE_DBSTRING=${GOOSE_DBSTRING_TEST} goose up"

goose-down-test: ## test用DBのdown
	${DOCKER_COMPOSE_EXEC} -it go sh -c "cd ${TARGET_MIGRATION_DIR} && GOOSE_DBSTRING=${GOOSE_DBSTRING_TEST} goose down"

goose-reset-test: ## test用DBのreset
	${DOCKER_COMPOSE_EXEC} -it go sh -c "cd ${TARGET_MIGRATION_DIR} && GOOSE_DBSTRING=${GOOSE_DBSTRING_TEST} goose reset"


