NETWORK_NAME	:=	dockers_network

create_network:
	docker	network	create	--subnet=192.168.50.0/24	--gateway=192.168.50.1	${NETWORK_NAME};
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

network_info:
	docker network inspect dockers_network

list_networks:
	docker network ls

docker network connect --ip 192.168.50.10 dockers_network 9e0ce03be0a30089ff81b4b4aaa0574d4675ffdf214369f2693acc9aa0fdb
docker network connect --ip 192.168.50.20 dockers_network 44fa54b0530b405dfa574765817f20cf3b3f2e77f41c5c50696ef150b3816745