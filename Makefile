#"postgres://your_username:your_password@your_db_docker_compose_service:5432/your_database?sslmode=disable"

# To check current version: docker-compose run offline_project_db_migration_tool -path /migrations -database "postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable" version
# To check current version: docker-compose run offline_project_db_migration_tool -path /migrations -database "postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable" force <desired version>
# Then run service: offline_project_db_migrate_up

init_db:
	docker-compose	up	offline_project_db	-d

check_migration_version:
	docker-compose	run	offline_project_db_migration_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	version

migrate_up:
	docker-compose	run	offline_project_db_migration_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	up

migrate_down:
	docker-compose	run	offline_project_db_migration_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	down	-all

force_migrate:
	docker-compose	run	offline_project_db_migration_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	force