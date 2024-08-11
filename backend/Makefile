start_all:
	docker compose up -d

stop_all:
	docker compose down

start_dummy:
	cd dummy; \
	docker build -o dummy .;\
	docker run dummy