package main

import (
	"flag"
	"fmt"
	"net/http"

	"gopherzz/curapi/internal/routes"
)

func main() {
	PORT := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	http.HandleFunc("/convert", routes.ConvertHandler)

	fmt.Println("Listening on port:", *PORT)
	http.ListenAndServe(":"+*PORT, nil)
}
