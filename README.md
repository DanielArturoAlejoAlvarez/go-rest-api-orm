# GO-REAT-API-ORM

## Description

This repository is a REST API with GO, Mux and SQLite.

## Installation

Using GO, SQLite, CompileDaemon,etc preferably.

## DataBase

Using SQLite preferably.

## Apps

Using Postman or RestEasy to feed the api.

## Usage

```html
$ git clone https://github.com/DanielArturoAlejoAlvarez/go-rest-api-orm.git
[NAME APP]
```

Follow the following steps and you're good to go! Important:

![alt text](https://user-images.githubusercontent.com/2971735/71788789-e8731200-3025-11ea-84dd-90298c51d954.gif)

## Coding

### Controllers

```go
...
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{name}", getUser).Methods("GET")
	router.HandleFunc("/users/{name}/{email}", saveUser).Methods("POST")
	router.HandleFunc("/users/{name}/{email}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")
	router.HandleFunc("/", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":5500", router))
}

func main() {
	fmt.Println("Welcome to REST API with ORM")

	InitialMigration()
	handleRequests()
}
...
```

### Models

```go
...
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

//User is esported
type User struct {
	gorm.Model
	Name  string
	Email string
}

//InitialMigration is exported
func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test_db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect DB!")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test_db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect DB!")
	}
	defer db.Close()

	var users []User

	db.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test_db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect DB!")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)

	db.Find(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func saveUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test_db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect DB!")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	fmt.Fprintf(w, "User saved successfully!")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test_db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect DB!")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email
	db.Save(&user)

	fmt.Fprintf(w, "User updated successfully!")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test_db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed connect DB")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User deleted successfully!")
}
...
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/DanielArturoAlejoAlvarez/go-rest-api-orm. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
````
