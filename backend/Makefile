NETWORK_NAME	:=	dockers_network
SUBNET	:=	192.168.50.0/24
GATEWAY	:=	192.168.50.1

create_network:
	docker	network	create	--subnet=${SUBNET}	--gateway=${GATEWAY}	${NETWORK_NAME};
remove_network:
	docker	network	rm	${NETWORK_NAME};

SERVICE_GATEWAY_NAME	:=	service_gateway
SERVICE_GATEWAY_IP	:=	192.168.50.10

connect_gateway:
	docker	network	connect	--ip	${SERVICE_GATEWAY_IP}	${NETWORK_NAME}	${SERVICE_GATEWAY_NAME};
diconnect_gateway:
	docker	network	disconnect	${NETWORK_NAME}	${SERVICE_GATEWAY_NAME};

SERVICE_DB_HANDLER_NAME	:=	service_db_handler
SERVICE_DB_HANDLER_IP	:=	192.168.50.20

connect_db_handler:
	docker	network	connect	--ip	${SERVICE_DB_HANDLER_IP}	${NETWORK_NAME}	${SERVICE_DB_HANDLER_NAME};
disconnect_db_handler:
	docker	network	disconnect	${NETWORK_NAME}	${SERVICE_DB_HANDLER_NAME};

connect_containers:	connect_gateway	connect_db_handler
disconnect_containers:	disconnect_gateway	disconnect_db_handler

stop_network:	disconnect_containers	remove_network
start_network:	create_network	connect_containers

#	INFO	==========================================

list_connected_containers:
	docker	network	inspect	\
	-f	'{{	range	$$key,	$$value	:=	.Containers	}}{{	printf	"%s\n"	$$key}}{{	end	}}'	\
	${NETWORK_NAME}
list_containers_ip:
	docker ps -q | xargs -n 1 docker inspect --format '{{ .Name }} {{range .NetworkSettings.Networks}} {{.IPAddress}}{{end}}' | sed 's#^/##';

network_info:
	docker network inspect dockers_network

list_networks:
	docker network ls