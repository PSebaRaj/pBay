.PHONY: protoUser

protoUser:
	protoc -I protos/ protos/user.proto --go_out=plugins=grpc:protos/user



