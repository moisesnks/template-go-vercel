package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/adshao/go-binance/v2"
)

// obtener los precios de las criptomonedas desde Binance
func getCryptoPrices() map[string]float64 {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || secretKey == "" {
		log.Fatal("API key and/or secret key are missing")
	}

	client := binance.NewClient(apiKey, secretKey)

	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		log.Fatalf("Failed to get crypto prices: %v", err)
	}

	cryptoPrices := make(map[string]float64)
	for _, p := range prices {
		priceFloat, err := strconv.ParseFloat(p.Price, 64)
		if err != nil {
			log.Printf("Error al convertir el precio: %v", err)
			continue
		}
		cryptoPrices[p.Symbol] = priceFloat
	}

	return cryptoPrices
}

// función serverless que maneja las solicitudes HTTP a la ruta /cryptoprices
func Cryptoprices(w http.ResponseWriter, r *http.Request) {
	prices := getCryptoPrices()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prices)
}
