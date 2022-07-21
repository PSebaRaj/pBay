# PBay 
[![Latest Release](https://img.shields.io/github/v/release/psebaraj/pbay?include_prereleases&style=for-the-badge)](https://github.com/psebaraj/pbay/releases)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/psebaraj/pbay/Go?style=for-the-badge)](https://github.com/PSebaRaj/pbay/actions/workflows/go.yml)
[![Go ReportCard](https://goreportcard.com/badge/github.com/psebaraj/pbay?style=for-the-badge)](https://goreportcard.com/report/psebaraj/pbay)
[![Lines of Code](https://img.shields.io/tokei/lines/github/psebaraj/pBay?style=for-the-badge)](https://github.com/psebaraj/pbay/actions)

NOTE: This project is currently being refactored. Check out the `refactoring` branch for more info.

An eBay clone, built with Go and React. Frontend can be found [here](https://github.com/PSebaRaj/pBay-Frontend)

Goal is to implement microservices, rather than the monolithic approach that was taken for my other projects (TradingPlatform, etc.), as well as implement RPCs (gRPC)

## Dependencies
- Go
- (gRPC) Protocol buffer compiler (protoc)
- (gRPC) Plug-ins:
	- `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
	- `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

## Usage
- To run the backend of the application, first clone the repository:
	- `git clone https://github.com/psebaraj/pbay.git`

- Navigate to the pBay directory

- (Re)make the protocall buffers for Go:
	- First update protoc plugin PATH: `export PATH="$PATH:$(go env GOPATH)/bin"`
	- Make all: `make protos`
	- Make specific XXX protobuf: `make protoXXX`

- Build and run the application:
	- `./start-pbay-backend.sh`


## Diagram
![pBay](./PBayDiagram.jpg)

Microservices should only know user email or given product name, and use RPC to retrieve the rest of the user's
	data

## To-Do:
- [ ] Add Swagger
	- [x] products
	- [ ] users
	- [ ] messages
- [x] shell script for starting all services
- [ ] add /products?=....., deleteProduct, modifyProduct routes to products
- [ ] gRPC
	- [x] create protobufs
	- [ ] implement user
	- [ ] implement address
	- [ ] implement currency
- [ ] create remaining microservices
	- [ ] shipping
	- [ ] payments
	- [ ] advertisements

## Contributor:
- [Patrick SebaRaj](https://github.com/PSebaRaj)
