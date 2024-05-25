package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Handlerfuns() {
	//Obtenemos todos los precios de las criptos ejemplo: http://localhost:8080/cryptoprices
	http.HandleFunc("/cryptoprices",HandleCryptoPrices)

	//Obtenemos solo una moneda ejemplo http://localhost:8080/cryptoprice?symbol=BTCUSDT
	http.HandleFunc("/cryptoprice",HandleSingleCryptoPrice)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	puerto := os.Getenv("PORT")
	if puerto == "" {
		puerto = ":8080"
	}
	log.Printf("Servidor en ejecución en http://localhost%s\n", puerto)
	log.Fatal(http.ListenAndServe(puerto, nil))
}
