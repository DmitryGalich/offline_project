NETWORK_NAME := dockers_network

connect_containers:
    @docker network create $(NETWORK_NAME) || true
    @docker network connect $(NETWORK_NAME) $(shell docker ps -q)

disconnect_containers:
    @docker network disconnect $(NETWORK_NAME) $(shell docker ps -q)
	@docker network rm $(NETWORK_NAME)