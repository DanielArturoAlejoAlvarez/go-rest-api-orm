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
