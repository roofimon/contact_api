package contact

import (
    "code.google.com/p/go-uuid/uuid"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var session *mgo.Session

type MongoProvider struct {
	store map[string]*Information
}

func NewMongoProvider() *MongoProvider {
	mp := &MongoProvider{}
	session, _ = mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)
	return mp
}

func GetSession() *mgo.Session {
	var ls = session.Clone()
	return ls
}

func Contact(ls *mgo.Session) *mgo.Collection {
	contact := ls.DB("test").C("contact")
	return contact
}

func (im *MongoProvider) All() []Information {
	var c = []Information{}
	s := GetSession()
	defer s.Close()
	Contact(s).Find(nil).All(&c)
	return c
}

func (im *MongoProvider) Add(c *Information) {
	s := GetSession()
	defer s.Close()
    c.Id = uuid.New()
	err := Contact(s).Insert(c)
	if err != nil {
		log.Fatal(err)
	}
}

func (im *MongoProvider) Update(c *Information) {
	target := bson.M{"id": c.Id}
	change := bson.M{"$set": bson.M{"id": c.Id, "email": c.Email, "title": c.Title, "content": c.Content}}
	s := GetSession()
	defer s.Close()
	err := Contact(s).Update(target, change)
	if err != nil {
		panic(err)
	}
}

func (im *MongoProvider) Remove(id string) error {
	target := bson.M{"id": id}
	s := GetSession()
	defer s.Close()
	err := Contact(s).Remove(target)
	if err != nil {
		panic(err)
		return errors.New("emit macho dwarf: elf header corrupted")
	}
	return nil
}

func (im *MongoProvider) Get(id string) (*Information, error) {
	result := Information{}
	s := GetSession()
	defer s.Close()
	err := Contact(s).Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return &result, nil
}
