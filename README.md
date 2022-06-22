# PBay
An eBay clone, built with Go and React.

Goal is to implement microservices, rather than the monolithic approach that was taken for my other projects (TradingPlatform, etc.), as well as implement RPCs (gRPC)

## Usage
To run the backend of the application, first clone the repository:
	`git clone https://github.com/psebaraj/pbay.git`

Open a terminal and navigate to the pBay directory

Build and run the application:
	`./start-pbay-backend.sh`


## Diagram
![pBay](./PBayDiagram.jpg)

## To-Do: (in-order)
- Add Swagger to Users microservice
- add /products?=....., deleteProduct, modifyProduct routes to products
- shipping microservice
- payments microservice
- advertisements microservice

## Contributor:
- [Patrick SebaRaj](https://github.com/PSebaRaj)

