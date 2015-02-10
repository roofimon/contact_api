package contact

import (
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
	var ls = session.Clone()
	defer ls.Close()
	collection := ls.DB("test").C("contact")
	err := collection.Insert(c)
	if err != nil {
		log.Fatal(err)
	}
}

func (im *MongoProvider) Update(c *Information) {

}

func (im *MongoProvider) Remove(id string) error {
	return errors.New("emit macho dwarf: elf header corrupted")
}

func (im *MongoProvider) Get(id string) (*Information, error) {
	var ls = session.Clone()
	defer ls.Close()
	collection := ls.DB("test").C("contact")
	result := Information{}
	err := collection.Find(bson.M{"email": "first@email.com"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return &result, nil
}
