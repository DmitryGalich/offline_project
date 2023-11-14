#"postgres://your_username:your_password@your_db_docker_compose_service:5432/your_database?sslmode=disable"

# To check current version: docker-compose run offline_project_db_migration_tool -path /migrations -database "postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable" version
# To check current version: docker-compose run offline_project_db_migration_tool -path /migrations -database "postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable" force <desired version>
# Then run service: offline_project_db_migrate_up

POSTGRES_COMPOSE_SERVICE:=offline_project_db
POSTGRES_DB:=offline_project_db
POSTGRES_PORT:=5432
POSTGRES_USER:=postgres
POSTGRES_PASSWORD:=123
POSTGRES_CONNECTING_STRING:="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_COMPOSE_SERVICE}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

INIT_DB_ARGS:=ARG1=${POSTGRES_DB}	ARG2=${POSTGRES_USER}	ARG3=${POSTGRES_PASSWORD}
init_db:
	${INIT_DB_ARGS}	docker-compose	up	offline_project_db	-d

check_migration_version:
	docker-compose	run	--rm	offline_project_db_migration_tool	-path	/migrations	-database	${POSTGRES_CONNECTING_STRING}	version

migrate_up:
	docker-compose	run	--rm	offline_project_db_migration_tool	-path	/migrations	-database	${POSTGRES_CONNECTING_STRING}	up

migrate_down:
	docker-compose	run	--rm	offline_project_db_migration_tool	-path	/migrations	-database	${POSTGRES_CONNECTING_STRING}	down	-all

# Call it like this: make force_migrate VERSION=1
VERSION:=""
force_migrate:
	docker-compose	run	--rm	offline_project_db_migration_tool	-path	/migrations	-database	${POSTGRES_CONNECTING_STRING}	force	${VERSION}

