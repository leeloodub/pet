package storage

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionParsedImg = "ParsedImages"
	CollectionRawImg    = "RawImages"
)

type Storage struct {
	address string
	DBName  string
	DBConn  *mgo.Database
}

func NewStorage(address, db string) *Storage {
	s, err := mgo.Dial(address)
	if err != nil {
		log.Fatal(err)
	}

	return &Storage{
		address: address,
		DBName:  db,
		DBConn:  s.DB(db),
	}
}

func (s *Storage) InsertParsed(img string) (string, error) {
	return s.insert(img, CollectionParsedImg)
}

func (s *Storage) InsertRaw(img string) (string, error) {
	return s.insert(img, CollectionRawImg)
}

func (s *Storage) insert(img string, collection string) (string, error) {
	i := bson.NewObjectId()
	return i.String(), s.DBConn.C(collection).Insert(bson.M{"_id": i, "record": img})
}
