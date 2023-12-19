NETWORK_NAME	:=	dockers_network

create_network:
	docker	network	create	${NETWORK_NAME};
remove_network:
	docker	network	rm	${NETWORK_NAME};

connect_containers:
	for	container_id	in	$$(docker	ps	-aq);	\
	do	\
		docker	network	connect	${NETWORK_NAME}	$$container_id;	\
	done;
disconnect_containers:
	for	container_id	in	$$(docker	ps	-aq);	\
	do	\
		docker	network	disconnect	${NETWORK_NAME}	$$container_id;	\
	done;

set_down_network:	disconnect_containers	remove_network
set_up_network:	create_network	connect_containers

list_connected_containers:
	docker	network	inspect	\
	-f	'{{	range	$$key,	$$value	:=	.Containers	}}{{	printf	"%s\n"	$$key}}{{	end	}}'	\
	${NETWORK_NAME}

list_containers_ip:
	docker ps -q | xargs -n 1 docker inspect --format '{{ .Name }} {{range .NetworkSettings.Networks}} {{.IPAddress}}{{end}}' | sed 's#^/##';
