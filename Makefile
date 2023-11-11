POSTGRES_COMPOSE_SERVICE:=offline_project_db
POSTGRES_DB:=offline_project_db
POSTGRES_PORT:=5432
POSTGRES_USER:=postgres
POSTGRES_PASSWORD:=123
POSTGRES_CONNECTING_STRING:="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_COMPOSE_SERVICE}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

INIT_DB_ARGS:=ARG1=${POSTGRES_DB}	ARG2=${POSTGRES_USER}	ARG3=${POSTGRES_PASSWORD}
init_db:
	${INIT_DB_ARGS}	docker-compose	up	${POSTGRES_COMPOSE_SERVICE}	-d

MIGRATE_UP_ARGS:=ARG1=${POSTGRES_CONNECTING_STRING}
migrate_up:
	docker-compose	up	offline_project_db_migrate_up	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	up

