# Protos - PBay
First use of gRPC / any RPCs.

Use gRPC as a lightweight, fast method to transfer data across services. Each service should only know the user email.

## Protos:

### User Proto:
- Send user name when given email (RPC called by Messages service)
- Send user payment info when given email (RPC called by Payments service)

### Currency Proto:
- Enable change of currency denomination for the Products, Shipping, and Payments services
	- System stores everything as USD
	- Changes denomination based on client location

### Product Proto:
- Send product weight when given product name (RPC called by Shipping service)

## Usage
- To make the protocall buffers for Go:
	- First update protoc plugin PATH: `export PATH="$PATH:$(go env GOPATH)/bin"`
	- Make all: `make protos`
	- Make specific XXX protobuf: `make protoXXX`

## To-Do:
- [x] user proto
- [x] currency proto
- [x] ~~address proto~~
- [x] product proto

