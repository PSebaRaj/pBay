# PBay
[![Latest Release](https://img.shields.io/github/v/release/psebaraj/pbay?include_prereleases&style=for-the-badge)](https://github.com/psebaraj/pbay/releases)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/psebaraj/pbay/Go?style=for-the-badge)](https://github.com/PSebaRaj/pbay/actions/workflows/go.yml)
[![Go ReportCard](https://goreportcard.com/badge/github.com/psebaraj/pbay?style=for-the-badge)](https://goreportcard.com/report/psebaraj/pbay)
[![Lines of Code](https://img.shields.io/tokei/lines/github/psebaraj/pbay?style=for-the-badge)](https://github.com/psebaraj/pbay/actions)

NOTE: This project is currently being refactored to somewhat follow this Go microservices [structure](https://betterprogramming.pub/how-are-you-structuring-your-go-microservices-a355d6293932).

An eBay clone, built with Go and React. Frontend can be found [here](https://github.com/PSebaRaj/pBay-Frontend).

Goal is to implement microservices, rather than the monolithic approach that was taken for my other projects (TradingPlatform, GoGetItDone, etc.), as well as implement RPCs (gRPC).

## Notes for myself
- [YAML Configs](https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp)


## Dependencies
- Go
- (gRPC) Protocol buffer compiler (protoc)
- (gRPC) Plug-ins:
	- `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
	- `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
- Each microservice has it's own list of dependencies, which can be found in it's respective `README.md`

## Diagram
![pBay](https://github.com/PSebaRaj/pBay/blob/main/PBayDiagram.jpg)

## To-Do:
- [ ] Add a deployment/usage section to the parent `README.md` file

## Contributor:
- [Patrick SebaRaj](https://github.com/PSebaRaj)
