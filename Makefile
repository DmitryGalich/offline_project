init_db:
	docker-compose	up	offline_project_db	-d

migrate_up:
	docker-compose up	offline_project_db_migrate_up

migrate_down:
	docker-compose up	offline_project_db_migrate_down

check_migration_version:
	docker-compose up	offline_project_db_check_version