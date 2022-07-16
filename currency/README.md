# Currency gRPC Server

https://grpc.io/docs/languages/go/quickstart/

## Progress
- Logging works (using pkg log)
- Need to add actual currencies

## gRPCURL
### Example usage
```
grpcurl --plaintext -d '{"Base":"GBP", "Destination":"USD"}' localhost:9092 currency.Currency.GetRate
```

### Example response
```json
{
  "rate": 0.5
}
```

