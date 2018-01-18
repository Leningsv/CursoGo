package models

import "github.com/zebresel-com/mongodm"

type User struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`

	FirstName string      `json:"firstname" bson:"firstname"`
	LastName  string      `json:"lastname"	 bson:"lastname"`
	UserName  string      `json:"username"	 bson:"username"`
	Messages  interface{} `json:"messages"	 bson:"messages" 	model:"Message" relation:"1n" autosave:"true"`
}
type Message struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`

	Sender   string `json:"sender" 	 bson:"sender"`
	Receiver string `json:"receiver"	 bson:"receiver"`
	Text     string `json:"text"	 bson:"text"`
}
type Customer struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`

	FirstName string   `json:"firstname" bson:"firstname"`
	LastName  string   `json:"lastname"	 bson:"lastname"`
	Address   *Address `json:"address"	 bson:"address"`
}

type Address struct {
	City    string `json:"city" 	 bson:"city"`
	Street  string `json:"street"	 bson:"street"`
	ZipCode int16  `json:"zip"	 bson:"zip"`
}
