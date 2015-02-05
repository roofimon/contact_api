package contact

type Provider interface {
	Add(information *Information)
	All() map[string]*Information
	Get(id string) (*Information, error)
	Remove(id string) error
	Update(information *Information)
}
