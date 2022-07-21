.PHONY: protos protoCurrency

protos: protoProduct

protoProduct:
	protoc --go_out=./proto-gen --go-grpc_out=./proto-gen proto/product.proto

install_swag:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: install_swag
	swagger generate spec -o ./swagger-gen/swagger.yaml --scan-models
