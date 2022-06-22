.PHONY: protos protoUser protoCurrency protoAddress

protos: protoUser protoCurrency protoAddress

protoUser:
	protoc --go_out=protos/user --go-grpc_out=protos/user protos/user.proto

protoCurrency:
	protoc --go_out=protos/currency --go-grpc_out=protos/currency protos/currency.proto

protoAddress:
	protoc --go_out=protos/address --go-grpc_out=protos/address protos/address.proto
