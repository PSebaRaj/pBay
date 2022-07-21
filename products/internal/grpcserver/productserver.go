package grpcserver

import (
	"context"
	"fmt"
	"log"
	"os"

	protos "github.com/psebaraj/pbay/products/proto-gen"
)

type Product struct {
	protos.UnimplementedProductServer
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
// Rate * Price (in Base Currency) --> Price in Destination Currency
func (p *Product) GetRate(ctx context.Context, sr *protos.ShippingRequest) (*protos.ShippingResponse, error) {
	// TODO: only create one instance of logger on startup, just pass it thru
	logger := log.New(os.Stdout, "GetRate ", 3)
	logger.Printf("Urgency:%s Weight:%f", sr.GetUrgency(), sr.GetWeight())
	cos := COST_OF_SHIPPING_PER_POUND[fmt.Sprintf("%s", sr.Urgency)] / sr.Weight

	// need to error handle
	return &protos.ShippingResponse{CostOfShipping: cos}, nil
}

var COST_OF_SHIPPING_PER_POUND = map[string]float32{
	"OVERNIGHT": 30.00,
	"NEXTDAY": 15.99,
	"TWODAY": 10.25,
	"STANDARD": 4.99,
	"INTERNATIONAL": 15.00,
}
