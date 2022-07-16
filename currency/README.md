# Currency gRPC Server

RUNNING ON PORT 9092

gRPC server used by other services that involve serving list prices, estimated shipping prices, etc. to provide the user with the price in their own
	regional currency. Specifically, the Products, Shipping, and Payments services are RPC clients.

## Current Rates
Sourced from the Federal Reserve's [Foreign Exchange Rates](https://www.federalreserve.gov/releases/h10/current/) (weekly updated).

Locally stored exchange rates, which can be found in the `EXCHANGE_RATES` map in `server/currency.go`, were last updated on: 7/16/22

## gRPCurl
- Docs can be found [here](https://github.com/fullstorydev/grpcurl).

### Example usage
If a service requests the exchange rate from US Dollars to Australian Dollars, they can request the following through
	gRPC (simulated through gRPCurl) below:
```
grpcurl --plaintext -d '{"Base":"USD", "Destination":"AUD"}' localhost:9092 Currency.GetRate
```
And the current exchange rate is sent as a response to the service:
```
{
	"rate": 1.458151
}
```

### Description
Running:
```
grpcurl --plaintext localhost:9092 describe ___
```
Describing Currency:
```
Currency is a service:
service Currency {
  rpc GetRate ( .RateRequest ) returns ( .RateResponse );
}
```
Describing Currency.GetRate:
```
Currency.GetRate is a method:
rpc GetRate ( .RateRequest ) returns ( .RateResponse );
```
Describing .RateRequest:
```
RateRequest is a message:
message RateRequest {
  .Currencies Base = 1;
  .Currencies Destination = 2;
}
```
Describing .RateResponse:
```
RateResponse is a message:
message RateResponse {
  float rate = 1;
}
```
Describing .Currencies:
```
Currencies is an enum:
enum Currencies {
  USD = 0;
  AUD = 1;
  BRL = 2;
  CAD = 3;
  CNY = 4;
  DKK = 5;
  EUR = 6;
  HKD = 7;
  INR = 8;
  JPY = 9;
  MYR = 10;
  MXN = 11;
  NZD = 12;
  NOK = 13;
  SGD = 14;
  ZAR = 15;
  KRW = 16;
  LKR = 17;
  SEK = 18;
  CHF = 19;
  TWD = 20;
  THB = 21;
  GBP = 22;
  VES = 23;
}
```

