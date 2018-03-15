package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

type Person struct {
	gorm.Model

	Firstname string
	Lastname  string
}

type People []Person

var db *gorm.DB
var err error

func main() {
	// Load Env File
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Open DB Connection
	db, err = gorm.Open(
		"mysql",
		os.Getenv("USERNAME")+":"+os.Getenv("PASSWORD")+
			"@tcp("+os.Getenv("HOST")+":3306)/"+os.Getenv("DBNAME")+
			"?charset=utf8&parseTime=True&loc=Local",
	)

	if err != nil {
		panic("Failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Person{})

	// Add Router
	router := mux.NewRouter()

	// Handle Routes
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people", CreatePeople).Methods("POST")
	router.HandleFunc("/people/list", InsertPeople).Methods("POST")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", router))
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person

	if err = db.First(&person, params["id"]).Error; err != nil {
		fmt.Fprintf(w, "ID %s Not Found", params["id"])
		return
	}

	db.Delete(&person)
	json.NewEncoder(w).Encode(&person)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person

	db.First(&person, params["id"])
	json.NewEncoder(w).Encode(person)

}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people People
	db.Find(&people)

	json.NewEncoder(w).Encode(people)
}

func CreatePeople(w http.ResponseWriter, r *http.Request) {
	var person Person

	json.NewDecoder(r.Body).Decode(&person)
	db.Create(&person)

	json.NewEncoder(w).Encode(person)
}

func InsertPeople(w http.ResponseWriter, r *http.Request) {
	var person Person
	records := 1000

	json.NewDecoder(r.Body).Decode(&person)
	now := time.Now()

	for i := 0; i < records; i++ {
		db.Exec(
			"insert into people (firstname, lastname, created_at, updated_at) values(?, ?, ?, ?)",
			person.Firstname,
			person.Lastname,
			now,
			now,
		)
	}
}
