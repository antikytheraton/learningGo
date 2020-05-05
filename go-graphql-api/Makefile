DOCKER_DB_NAME = go_graphql_api
DOCKER_DB_USER = postgres
DOCKER_DB_PASS = mysecretpassword
DOCKER_DB_PORT = 5432
DOCKER_DB_HOST = 127.0.0.1
DATABASE_URL = postgres://$(DOCKER_DB_USER):$(DOCKER_DB_PASS)@$(DOCKER_DB_HOST):$(DOCKER_DB_PORT)/$(DOCKER_DB_NAME)



define image-run
	docker run \
		--name $(DOCKER_DB_NAME) \
		-e POSTGRES_USER=$(DOCKER_DB_USER) \
		-e POSTGRES_PASSWORD=$(DOCKER_DB_PASS) \
		-e POSTGRES_DB=$(DOCKER_DB_NAME) \
		-p $(DOCKER_DB_PORT):5432 \
		-d postgres
endef

define psql-exec
	docker exec $(DOCKER_DB_NAME) psql $(DATABASE_URL) -c '$(1)'
endef

define psql-file
	docker exec $(DOCKER_DB_NAME) psql $(DATABASE_URL) < '$(1)';
endef

db-up:
	$(call image-run)

db-down:
	docker container stop $(DOCKER_DB_NAME)
	docker container rm $(DOCKER_DB_NAME)

db-reset: db-down db-up

db-populate:
	$(call psql-exec,DROP SCHEMA IF EXISTS public CASCADE)
	$(call psql-exec,CREATE SCHEMA public)
	$(call psql-exec,CREATE DATABASE go_graphql_db)
	$(call psql-file,./migrations/init.sql)