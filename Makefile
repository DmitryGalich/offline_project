init_db:
	docker-compose	up	offline_project_db	-d

migrate_up:
	docker-compose up	offline_project_db_migrate_up

migrate_down:
	docker-compose up	offline_project_db_migrate_down