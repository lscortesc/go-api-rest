package main

import (
	"encoding/json"
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

	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people", CreatePeople).Methods("POST")
	router.HandleFunc("/people/list", InsertPeople).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", router))
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

	json.NewDecoder(r.Body).Decode(&person)
	now := time.Now()

	for i := 0; i < 1000; i++ {
		db.Exec(
			"insert into people (firstname, lastname, created_at, updated_at) values(?, ?, ?, ?)",
			person.Firstname,
			person.Lastname,
			now,
			now,
		)
	}
}
