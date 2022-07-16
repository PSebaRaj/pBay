package server

import (
	"context"
	"fmt"
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
// Rate * Price (in Base Currency) --> Price in Destination Currency
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	// TODO: only create one instance of logger on startup, just pass it thru
	logger := log.New(os.Stdout, "GetRate ", 3)
	logger.Printf("Base:%s Dest:%s", rr.GetBase(), rr.GetDestination())
	exchangeRate := EXCHANGE_RATES[fmt.Sprintf("%s", rr.Base)] / EXCHANGE_RATES[fmt.Sprintf("%s", rr.Destination)]

	// need to error handle
	return &protos.RateResponse{Rate: exchangeRate}, nil
}

// TODO: Last Updated: 7/16/22
// Source: https://www.federalreserve.gov/releases/h10/current/
var EXCHANGE_RATES = map[string]float32{
	"USD": 1.0000,    // US Dollar
	"AUD": 0.6858,    // Australian Dollar
	"BRL": 5.2961,    // Brazilian Real
	"CAD": 1.2947,    // Canadian Dollar
	"CNY": 6.6945,    // Chinese Yuan (Renminbi)
	"DKK": 7.3262,    // Danish Krone
	"EUR": 1.0178,    // (EMU) Euro
	"HKD": 7.8483,    // Hong Kong Dollar
	"INR": 79.2500,   // Indian Rupee
	"JPY": 136.1600,  // Japanese Yen
	"MYR": 4.4262,    // Malaysian Ringgit
	"MXN": 20.4440,   // Mexican Peso
	"NZD": 0.6196,    // New Zealand Dollar
	"NOK": 10.0902,   // Norwegian Krone
	"SGD": 1.3985,    // Singapore Dollar
	"ZAR": 16.8550,   // South African Rand
	"KRW": 1299.5100, // (South) Korean Won
	"LKR": 360.0000,  // Sri Lanka Rupee
	"SEK": 10.5523,   // Swedish Krona
	"CHF": 0.9775,    // Swiss Franc
	"TWD": 29.7500,   // New Taiwan Dollar
	"THB": 36.0700,   // Taiwanese Baht
	"GBP": 1.2036,    // (GB) Pound Sterling
	"VES": 5.5607,    // Bolivar Soberano
}
