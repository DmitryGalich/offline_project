ROOT_DIR:=$(realpath $(shell dirname $(firstword $(MAKEFILE_LIST))))

gen_example:
	protoc	-I=$(ROOT_DIR)	--go_out=$(ROOT_DIR)/example/gen/	$(ROOT_DIR)/example/example.proto