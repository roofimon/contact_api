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
	ch := NewContactHandler()
	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/contact/:id", ch.Get},
		&rest.Route{"GET", "/contacts", ch.All},
		&rest.Route{"POST", "/contact", ch.Add},
		&rest.Route{"PUT", "/contact", ch.Update},
		&rest.Route{"DELETE", "/contact/:id", ch.Remove},
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type ContactHandler struct {
	provider contact.Provider
}

func NewContactHandler() *ContactHandler {
	//return &ContactHandler{provider: contact.NewMemoryProvider()}
	return &ContactHandler{provider: contact.NewMongoProvider()}
}

func (ch *ContactHandler) Remove(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	err := ch.provider.Remove(id)
	if err == nil {
		w.WriteJson("Remove successful")
	} else {
		rest.Error(w, "Contact Not Found", http.StatusInternalServerError)
		return
	}
}

func (ch *ContactHandler) Get(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	contact, err := ch.provider.Get(id)
	if err == nil {
		w.WriteJson(contact)
	} else {
		rest.Error(w, "Contact Not Found", http.StatusInternalServerError)
		return
	}
}

func (ch *ContactHandler) All(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(ch.provider.All())
}

func (ch *ContactHandler) Add(w rest.ResponseWriter, r *rest.Request) {
	var information contact.Information
	r.DecodeJsonPayload(&information)
	ch.provider.Add(&information)
	w.WriteJson("Add new entry successfully")
}

func (ch *ContactHandler) Update(w rest.ResponseWriter, r *rest.Request) {
	var information contact.Information
	r.DecodeJsonPayload(&information)
	ch.provider.Update(&information)
	w.WriteJson("Update entry successfully")
}
