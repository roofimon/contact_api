package contact

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

type ContactHandler struct {
	provider Provider
}

func NewHandler() *ContactHandler {
	return &ContactHandler{provider: NewMongoProvider()}
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
	var information Information
	r.DecodeJsonPayload(&information)
	ch.provider.Add(&information)
	w.WriteJson("Add new entry successfully")
}

func (ch *ContactHandler) Update(w rest.ResponseWriter, r *rest.Request) {
	var information Information
	r.DecodeJsonPayload(&information)
	ch.provider.Update(&information)
	w.WriteJson("Update entry successfully")
}
