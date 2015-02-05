package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"gojsonrest"
	"gojsonrest/service"
	"gojsonrest/service/memory"
	"log"
	"net/http"
)

var cp service.ContactProvider

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	ch := NewContactHandler()
	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/contact/:id", ch.GetContact},
		&rest.Route{"GET", "/contacts", ch.GetAllContacts},
		&rest.Route{"POST", "/contact", ch.AddEntry},
		&rest.Route{"PUT", "/contact", ch.UpdateEntry},
		&rest.Route{"DELETE", "/contact/:id", ch.RemoveEntry},
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type ContactHandler struct {
	provider service.ContactProvider
}

func NewContactHandler() *ContactHandler {
	return &ContactHandler{provider: memory.NewContactProvider()}
}

func (ch *ContactHandler) RemoveEntry(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	contact, _ := ch.provider.GetEntry(id)
	if contact != nil {
		ch.provider.RemoveEntry(id)
		w.WriteJson("Remove entry successfully")
	}else{
		rest.Error(w, "Contact Not Found", http.StatusInternalServerError)
		return
	}
}

func (ch *ContactHandler) GetContact(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	contact, _ := ch.provider.GetEntry(id)
	if contact != nil {
		w.WriteJson(contact)
	}else{
		rest.Error(w, "Contact Not Found", http.StatusInternalServerError)
		return
	}
}

func (ch *ContactHandler) GetAllContacts(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(ch.provider.GetAllEntries())
}

func (ch *ContactHandler) AddEntry(w rest.ResponseWriter, r *rest.Request) {
	contact := api.Contact{}
	r.DecodeJsonPayload(&contact)
	ch.provider.AddEntry(contact)
	w.WriteJson("Add new entry successfully")
}

func (ch *ContactHandler) UpdateEntry(w rest.ResponseWriter, r *rest.Request) {
	contact := api.Contact{}
	r.DecodeJsonPayload(&contact)
	ch.provider.UpdateEntry(contact)
	w.WriteJson("Update entry successfully")


	id := r.PathParam("id")
	contact, _ := ch.provider.GetEntry(id)
	if contact != nil {
		ch.provider.RemoveEntry(id)
		w.WriteJson("Remove entry successfully")
	}else{
		rest.Error(w, "Contact Not Found", http.StatusInternalServerError)
		return
	}
}
