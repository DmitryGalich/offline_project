init_db:
	docker-compose	up	offline_project_db	-d

migrate_up:
	docker-compose up	offline_project_db_migrate_up

migrate_down:
	docker-compose up	offline_project_db_migrate_down

check_migration_version:
	docker-compose up	offline_project_db_check_version

check_migration_version_via_tool:
	docker-compose	run	offline_project_db_migrate_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	version

migrate_up_via_tool:
	docker-compose	run	offline_project_db_migrate_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	up

migrate_down_via_tool:
	docker-compose	run	offline_project_db_migrate_tool	-path	/migrations	-database	"postgres://postgres:123@offline_project_db:5432/offline_project_db?sslmode=disable"	down	-all