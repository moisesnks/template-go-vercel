package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
)

// obtener el precio de una criptomoneda específica desde Binance
func getSingleCryptoPrice(symbol string) (float64, error) {
	err := godotenv.Load()
	if err != nil {
		return 0, fmt.Errorf("error al cargar el archivo .env: %v", err)
	}

	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || secretKey == "" {
		return 0, fmt.Errorf("API key and/or secret key are missing")
	}

	client := binance.NewClient(apiKey, secretKey)

	price, err := client.NewListPricesService().Symbol(symbol).Do(context.Background())

	if err != nil {
		return 0, fmt.Errorf("failed to get crypto price: %v", err)
	}

	if len(price) == 0 {
		return 0, fmt.Errorf("no price found for symbol: %s", symbol)
	}

	priceFloat, err := strconv.ParseFloat(price[0].Price, 64)

	if err != nil {
		return 0, fmt.Errorf("error al convertir el precio: %v", err)
	}

	return priceFloat, nil

}


// función serverless que maneja las solicitudes HTTP a la ruta /cryptoprice

func Cryptoprice (w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")

	if symbol == "" {
		http.Error(w, "Missing symbol parameter", http.StatusBadRequest)
		return
	}

	price, err := getSingleCryptoPrice(symbol)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]float64{symbol: price}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

