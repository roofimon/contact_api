package service

import "gojsonrest"

type ContactProvider interface {
    AddEntry(contact api.Contact)
    GetAllEntries() map[string]*api.Contact
    GetEntry(id string) (*api.Contact, error)
    RemoveEntry(id string) 
    UpdateEntry(contact api.Contact)
}
