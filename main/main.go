package main

import (
	"contact"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	handler := contact.NewHandler()
	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/contact/:id", handler.Get},
		&rest.Route{"GET", "/contacts", handler.All},
		&rest.Route{"POST", "/contact", handler.Add},
		&rest.Route{"PUT", "/contact", handler.Update},
		&rest.Route{"DELETE", "/contact/:id", handler.Remove},
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
