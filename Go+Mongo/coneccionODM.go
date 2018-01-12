package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"./models"
	"github.com/zebresel-com/mongodm"
)

func main() {
	file, err := ioutil.ReadFile("locals.json")

	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var localMap map[string]map[string]string
	json.Unmarshal(file, &localMap)

	dbConfig := &mongodm.Config{
		DatabaseHosts: []string{"localhost:27017"},
		DatabaseName:  "mongodm_sample",
		// DatabaseUser:     "admin",
		// DatabasePassword: "admin",
		Locals: localMap}["en-US"],
	}

	connection, err := mongodm.Connect(dbConfig)

	if err != nil {
		fmt.Println("Database connection error: %v", err)
	}

	connection.Register(&models.User{}, "users")
	connection.Register(&models.Message{}, "messages")
	connection.Register(&models.Customer{}, "customers")

	User := connection.Model("User")

	User.Find()

	user := &models.User{}

	User.New(user) //this sets the connection/collection for this type and is strongly necessary(!) (otherwise panic)

	user.FirstName = "Max"
	user.LastName = "Mustermann"

	err = user.Save()

	fmt.Println("Fin")
}
