package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Age  int           `bson:"age"`
}

type Admin struct {
	Person      `bson:",inline"`
	Password    string   `bson:"password"`
	Permissions []string `bson:"permissions"`
}

func main() {
	session, _ := mgo.Dial("mongodb://localhost/test")
	defer session.Close()
	db := session.DB("test")

	me := &Admin{
		Person: Person{
			ID:   bson.NewObjectId(),
			Name: "Jhon Donner",
			Age:  17,
		},
		Password:    "foobar",
		Permissions: []string{"read", "write"},
	}
	col := db.C("people")
	err := col.Insert(me)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	query := db.C("people").Find(bson.M{})
	p := new(Person)
	query.One(&p)
	a := new(Admin)
	query.One(&a)

	fmt.Printf("%+v\n", p)
	fmt.Printf("%+v\n", a)
}
