package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/planetmint/faucet/config"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Show index.html
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, config.GetConfig())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
}

func main() {
	// Load our configuration file
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("fatal error loading config file: %s", err)
	}
	serviceBind := config.GetString("service-bind")
	servicePort := config.GetInt("service-port")
	// Start our service
	log.Printf("Listening on '%s:%d' ...", serviceBind, servicePort)
	http.HandleFunc("/", indexHandler)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", serviceBind, servicePort), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
