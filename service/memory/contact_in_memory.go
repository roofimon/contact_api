package memory

import (
	"code.google.com/p/go-uuid/uuid"
	"errors"
	"gojsonrest"
)

func NewContactProvider() *InMemoryContactProvider {
	im := &InMemoryContactProvider{}
	im.store = map[string]*api.Contact{}
	initialContactStore(im)
	return im
}

func initialContactStore(im *InMemoryContactProvider) {
	firstContact := &api.Contact{Id: uuid.New(), Email: "first@email.com", Title: "First", Content: "First content"}
	secondContact := &api.Contact{Id: uuid.New(), Email: "second@email.com", Title: "Second", Content: "Second content"}

	im.store[firstContact.Id] = firstContact
	im.store[secondContact.Id] = secondContact
}

type InMemoryContactProvider struct {
	store map[string]*api.Contact
}

func (im *InMemoryContactProvider) GetAllEntries() map[string]*api.Contact {
	return im.store
}

func (im *InMemoryContactProvider) AddEntry(contact api.Contact) {
	contact.Id = uuid.New()
	im.store[contact.Id] = &contact
}

func (im *InMemoryContactProvider) UpdateEntry(contact api.Contact) {
	im.store[contact.Id] = &contact
}

func (im *InMemoryContactProvider) RemoveEntry(id string) {
	delete(im.store, id)
}

func (im *InMemoryContactProvider) GetEntry(id string) (*api.Contact, error) {
	return im.store[id], errors.New("emit macho dwarf: elf header corrupted")
}
