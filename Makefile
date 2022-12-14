WORKDIR = $(shell pwd)
DOCKER_DIR = $(WORKDIR)/docker
ANTLR_IMAGE = antlr4
MOUNT_DIR = /datastx
BASE_ANTLR_CMD=docker run -it -v $(PWD):/datastx antlr4 -Dlanguage=Go
SNOWFLAKE_MOUNT_PARSER_DIR = $(MOUNT_DIR)/internal/databases/snowflake/snowflake-parser
POSTGRES_MOUNT_PARSER_DIR = $(MOUNT_DIR)/internal/databases/postgres/postgres-parser
LOOKER_MOUNT_PARSER_DIR = $(MOUNT_DIR)/internal/visual/looker/looker-parser
JINJA_MOUNT_PARSER_DIR =  $(MOUNT_DIR)/internal/jinja2/jinja-parser
PARSER_PATTERN = -visitor

build-base-datastx:
	docker build -t datastx . -f $(DOCKER_DIR)/Dockerfile.datastx

test: build-base-datastx
	docker run -it datastx go test ./internal/...

build-datastx: build-base-datastx
	docker build -t datastx-prod . -f $(DOCKER_DIR)/Dockerfile.datastxslim

run-datastx: build-datastx
	docker run -it datastx-prod

repl:
	go run cmd/main.go

purge-branches:
	git branch | grep -v "main" | xargs git branch -D 

antlr-build:
	echo "Building $(ANTLR_IMAGE) docker image"
	echo "WORKDIR: $(WORKDIR)"
	docker build -t $(ANTLR_IMAGE) $(WORKDIR) -f $(DOCKER_DIR)/Dockerfile.antlr

antlr-build-looker:
	echo "Building $(ANTLR_IMAGE) docker image"
	echo "WORKDIR: $(WORKDIR)"
	docker build -t $(ANTLR_IMAGE) $(WORKDIR) -f $(DOCKER_DIR)/Dockerfile.lookerantlr

antlr-snowflake: antlr-build
	$(BASE_ANTLR_CMD) $(SNOWFLAKE_MOUNT_PARSER_DIR)/SnowflakeLexer.g4 $(PARSER_PATTERN) -o $(SNOWFLAKE_MOUNT_PARSER_DIR)
	$(BASE_ANTLR_CMD) $(SNOWFLAKE_MOUNT_PARSER_DIR)/SnowflakeParser.g4 $(PARSER_PATTERN) -o $(SNOWFLAKE_MOUNT_PARSER_DIR)

antlr-postres: antlr-build
	$(BASE_ANTLR_CMD) $(POSTGRES_MOUNT_PARSER_DIR)/PostgreSQLLexer.g4 $(PARSER_PATTERN) -o $(POSTGRES_MOUNT_PARSER_DIR)
	$(BASE_ANTLR_CMD) $(POSTGRES_MOUNT_PARSER_DIR)/PostgreSQLParser.g4 $(PARSER_PATTERN) -o $(POSTGRES_MOUNT_PARSER_DIR)

antlr-looker: antlr-build-looker
	$(BASE_ANTLR_CMD) $(LOOKER_MOUNT_PARSER_DIR)/LookML.g4 $(PARSER_PATTERN) -o $(LOOKER_MOUNT_PARSER_DIR)
	
antlr-jinja2: antlr-build-looker
	$(BASE_ANTLR_CMD) $(JINJA_MOUNT_PARSER_DIR)/Jinja.g4 $(PARSER_PATTERN) -o $(JINJA_MOUNT_PARSER_DIR)

clean-docker:
	docker system prune -a

