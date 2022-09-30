# declare fields
.PHONY: dev

# include config，initialize PROJECT_NAME params
include ./docker/dev/.env
PWD := $(shell pwd)
DEV_NAME = $(PROJECT_NAME)-$(USER_NAME)

# start dev environment
dev:
	go mod vendor && cd docker/dev && docker-compose -p "$(DEV_NAME)" down && docker-compose -p "$(DEV_NAME)" up --force-recreate