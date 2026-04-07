BINARY     := databasemix
SRC_DIR    := src
BUILD_DIR  := .
OUTDIR     := test_output

# Docker command (use sudo if not in docker group)
DOCKER_COMPOSE := sudo docker compose

# MySQL test containers
MYSQL_VERSIONS := 5.7 8.0 8.4
MYSQL_PORTS_5.7 := 3357
MYSQL_PORTS_8.0 := 3380
MYSQL_PORTS_8.4 := 3384

# PostgreSQL test containers
PG_VERSIONS := 16 17 18
PG_PORTS_16 := 5416
PG_PORTS_17 := 5417
PG_PORTS_18 := 5418

.PHONY: build clean test test-mysql test-postgres test-mysql-% test-postgres-% \
        containers-up containers-down containers-up-mysql containers-up-postgres \
        containers-down-mysql containers-down-postgres help

## ---- Build ----

build: ## Build the binary
	cd $(SRC_DIR) && go build -o ../$(BINARY) .

clean: ## Remove the binary and test output
	rm -f $(BINARY)
	rm -rf $(OUTDIR)

## ---- Test (full) ----

test: build test-mysql test-postgres ## Run all tests (MySQL + PostgreSQL)

## ---- MySQL tests ----

test-mysql: build $(addprefix test-mysql-,$(MYSQL_VERSIONS)) ## Run all MySQL tests

test-mysql-%: build
	@mkdir -p $(OUTDIR)
	@VERSION=$*; PORT=$(MYSQL_PORTS_$*); \
	CONTAINER=dbmix-mysql$$(echo $$VERSION | tr -d '.')-test; \
	COMPOSE=test_containers/mysql-$$VERSION/docker-compose.yml; \
	echo "=== MySQL $$VERSION (port $$PORT) ==="; \
	$(DOCKER_COMPOSE) -f $$COMPOSE up -d --build; \
	echo "Waiting for MySQL $$VERSION to be ready..."; \
	for i in $$(seq 1 90); do \
		sudo docker exec $$CONTAINER mysqladmin ping -h localhost -u root -prootpass 2>/dev/null && break; \
		sleep 3; \
	done; \
	echo "Waiting for init scripts to complete..."; \
	sleep 10; \
	echo "Running databasemix against MySQL $$VERSION ..."; \
	./$(BINARY) -type=mysql -host=localhost -port=$$PORT -user=root -password=rootpass -database=testdb \
		-outfile=$(OUTDIR)/mysql-$$VERSION 2>&1; \
	STATUS=$$?; \
	$(DOCKER_COMPOSE) -f $$COMPOSE down; \
	if [ $$STATUS -ne 0 ]; then echo "FAIL: MySQL $$VERSION"; exit 1; fi; \
	echo "OK: MySQL $$VERSION -> $(OUTDIR)/mysql-$$VERSION.md"

## ---- PostgreSQL tests ----

test-postgres: build $(addprefix test-postgres-,$(PG_VERSIONS)) ## Run all PostgreSQL tests

test-postgres-%: build
	@mkdir -p $(OUTDIR)
	@VERSION=$*; PORT=$(PG_PORTS_$*); \
	CONTAINER=dbmix-postgres$$VERSION-test; \
	COMPOSE=test_containers/postgres-$$VERSION/docker-compose.yml; \
	echo "=== PostgreSQL $$VERSION (port $$PORT) ==="; \
	$(DOCKER_COMPOSE) -f $$COMPOSE up -d --build; \
	echo "Waiting for PostgreSQL $$VERSION to be ready..."; \
	for i in $$(seq 1 60); do \
		sudo docker exec $$CONTAINER pg_isready -U postgres 2>/dev/null && break; \
		sleep 2; \
	done; \
	echo "Running databasemix against PostgreSQL $$VERSION ..."; \
	./$(BINARY) -type=postgres -host=localhost -port=$$PORT -user=postgres -password=rootpass -database=testdb \
		-outfile=$(OUTDIR)/postgres-$$VERSION 2>&1; \
	STATUS=$$?; \
	$(DOCKER_COMPOSE) -f $$COMPOSE down; \
	if [ $$STATUS -ne 0 ]; then echo "FAIL: PostgreSQL $$VERSION"; exit 1; fi; \
	echo "OK: PostgreSQL $$VERSION -> $(OUTDIR)/postgres-$$VERSION.md"

## ---- Container management ----

containers-up: containers-up-mysql containers-up-postgres ## Start all test containers

containers-down: containers-down-mysql containers-down-postgres ## Stop all test containers

containers-up-mysql: ## Start all MySQL test containers
	@for v in $(MYSQL_VERSIONS); do \
		echo "Starting MySQL $$v ..."; \
		$(DOCKER_COMPOSE) -f test_containers/mysql-$$v/docker-compose.yml up -d --build; \
	done

containers-up-postgres: ## Start all PostgreSQL test containers
	@for v in $(PG_VERSIONS); do \
		echo "Starting PostgreSQL $$v ..."; \
		$(DOCKER_COMPOSE) -f test_containers/postgres-$$v/docker-compose.yml up -d --build; \
	done

containers-down-mysql: ## Stop all MySQL test containers
	@for v in $(MYSQL_VERSIONS); do \
		$(DOCKER_COMPOSE) -f test_containers/mysql-$$v/docker-compose.yml down 2>/dev/null || true; \
	done

containers-down-postgres: ## Stop all PostgreSQL test containers
	@for v in $(PG_VERSIONS); do \
		$(DOCKER_COMPOSE) -f test_containers/postgres-$$v/docker-compose.yml down 2>/dev/null || true; \
	done

## ---- Help ----

help: ## Show this help
	@grep -E '^[a-zA-Z_%-]+:.*?## ' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-28s\033[0m %s\n", $$1, $$2}'
