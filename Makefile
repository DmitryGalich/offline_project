init_db:
	docker-compose	up	offline_project_db	-d

open_db:
	docker exec -it offline_project_db bash

migrate_up:
	docker-compose	up	offline_project_db_migrate_up

migrate_down:
	docker-compose	up	offline_project_db_migrate_down

MIGRATE_TOOL_STRING:=docker-compose	run	--rm	offline_project_db_migration_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"

migration_version:
	${MIGRATE_TOOL_STRING}	version

# Call it like this: make force_migration_version VERSION=1
VERSION:=""
force_migration_version:
	${MIGRATE_TOOL_STRING}	force	${VERSION}