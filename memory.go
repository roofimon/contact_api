package contact

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"errors"
)

func NewMemoryProvider() *InMemoryContactProvider {
	im := &InMemoryContactProvider{}
	im.store = map[string]*Information{}
	initialContactStore(im)
	return im
}

func initialContactStore(im *InMemoryContactProvider) {
	firstContact := &Information{Id: uuid.New(), Email: "first@email.com", Title: "First", Content: "First content"}
	secondContact := &Information{Id: uuid.New(), Email: "second@email.com", Title: "Second", Content: "Second content"}

	im.store[firstContact.Id] = firstContact
	im.store[secondContact.Id] = secondContact
}

type InMemoryContactProvider struct {
	store map[string]*Information
}

func (im *InMemoryContactProvider) All() map[string]*Information {
	return im.store
}

func (im *InMemoryContactProvider) Add(c *Information) {
	c.Id = uuid.New()
	im.store[c.Id] = c
}

func (im *InMemoryContactProvider) Update(c *Information) {
	im.store[c.Id] = c
}

func (im *InMemoryContactProvider) Remove(id string) error {
	delete(im.store, id)
	i := im.store[id]
	if i != nil {
		return errors.New("emit macho dwarf: elf header corrupted")
	}else{
		return nil
	}
}

func (im *InMemoryContactProvider) Get(id string) (*Information, error) {
	i := im.store[id]
	if i != nil {
		fmt.Println(i)
		return i, nil
	}else{
		return nil, errors.New("emit macho dwarf: elf header corrupted")
	}
}
