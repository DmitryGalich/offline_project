
init_db:
	docker-compose	up	offline_project_db	-d

check_migration_version:
	docker-compose	run	offline_project_db_migrate_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	version

migrate_up:
	docker-compose	run	offline_project_db_migrate_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	up

migrate_down:
	docker-compose	run	offline_project_db_migrate_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	down	-all