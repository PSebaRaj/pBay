package server

import (
	"context"
	"log"
	"os"

	"github.com/psebaraj/pbay/currency/protos/currency"
	protos "github.com/psebaraj/pbay/currency/protos/currency"
)

type Currency struct {
	currency.UnimplementedCurrencyServer
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	// TODO: only create one instance of logger on startup, just pass it thru
	logger := log.New(os.Stdout, "GetRate ", 3)
	logger.Printf("Base:%s Dest:%s", rr.GetBase(), rr.GetDestination())

	// need to error handle
	return &protos.RateResponse{Rate: 0.5}, nil
}
