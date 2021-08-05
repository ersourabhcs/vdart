package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

var currencyTicker Ticker

func main() {

	// origin := "http://localhost/"
	// url := "wss://api.hitbtc.com/api/3/ws/public"
	// ws, err := websocket.Dial(url, "", origin)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if _, err := ws.Write([]byte(`{"type":"subscribe"}`)); err != nil {
	// 	log.Fatal(err)
	// }
	// var msg = make([]byte, 512)
	// var n int
	// for 1 > 0 {
	// 	if n, err = ws.Read(msg); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("Received: %s.\n", msg[:n])
	// }

	port := ":9000"

	router := httprouter.New()
	router.RedirectTrailingSlash = true
	AddRouteHandlers(router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "Authorization"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		//Debug: true,
	})

	log.Fatal(http.ListenAndServe(port, c.Handler(router)))
}

func AddRouteHandlers(router *httprouter.Router) {

	router.GET("/api/Currency/:symbol", GetCurrencyPrice)
}
