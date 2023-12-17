ROOT_DIR:=$(realpath $(shell dirname $(firstword $(MAKEFILE_LIST))))

gen_helloworld:
	protoc	-I=$(ROOT_DIR)	--go_out=$(ROOT_DIR)/helloworld/gen/	--go_opt=paths=source_relative	--go-grpc_out=$(ROOT_DIR)/helloworld/gen/	--go-grpc_opt=paths=source_relative	$(ROOT_DIR)/helloworld/helloworld.proto