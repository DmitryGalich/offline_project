ROOT_DIR:=$(realpath $(shell dirname $(firstword $(MAKEFILE_LIST))))

gen_kek:
	protoc	-I=$(ROOT_DIR)	--go_out=$(ROOT_DIR)/kek/gen/	--go_opt=paths=source_relative	--go-grpc_out=$(ROOT_DIR)/kek/gen/	--go-grpc_opt=paths=source_relative	$(ROOT_DIR)/kek/kek.proto