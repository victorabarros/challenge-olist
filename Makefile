.DEFAULT_GOAL := help

APP_NAME?=$(shell pwd | xargs basename)
APP_DIR = /go/src/github.com/victorabarros/${APP_NAME}
PWD=$(shell pwd)

clean-up:
	@docker rm -f ${APP_NAME}

debug:
	@echo "\e[1m\033[32m\nDebug mode\e[0m"
	docker run -it -v ${PWD}:${APP_DIR} -w ${APP_DIR} \
		-p 8092:8092 --name ${APP_NAME} golang bash

fmt:
	docker run -v $(pwd):/go/src/github.com/victorabarros/fmt/ \
		-w /go/src/github.com/victorabarros/fmt/ golang:1.14 gofmt -e -l -s -w .

linter:
	# https://github.com/golangci/golangci-lint
	docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.24.0 golangci-lint run -v

log:
	docker logs ${APP_NAME} > logs.log

run:
# Como repassar a flag do comando?
	docker run -v ${PWD}:${APP_DIR} -w ${APP_DIR} \
		-p 8092:8092 --name ${APP_NAME} golang go run main.go

test:
	docker run -v ${PWD}:${APP_DIR} -w ${APP_DIR} \
		--env-file .env golang:1.14.2 \
		go test ./... -v -cover -race -coverprofile=c.out

test-log:
	@rm -rf dev/tests*.log
	@make test > dev/tests.log
	@cat dev/tests.log  | grep "coverage: " > dev/tests-summ.log

test-html-coverage:
	@rm -rf c.out
	@make test
	@go tool cover -html=c.out
