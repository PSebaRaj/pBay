.PHONY: protos protoUser protoCurrency protoProduct

protos: protoUser protoCurrency protoProduct

protoUser:
	protoc --go_out=protos/user --go-grpc_out=protos/user protos/user.proto

protoCurrency:
	protoc --go_out=protos/currency --go-grpc_out=protos/currency protos/currency.proto

protoProduct:
	protoc --go_out=protos/product --go-grpc_out=protos/product protos/product.proto
