package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var hraModifiers = map[string]float64{
	"IT":       0.10,
	"Security": 0.14,
}

type Person struct {
	Id          int      `json:"Id"`
	Name        string   `json:"Name"`
	Departments []string `json:"Departments"`
	BasicPay    float64  `json:"BasicPay"`
}

type People []Person

var people = []Person{
	{1, "bob", []string{"IT", "Security"}, 9.99},
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var newPerson Person
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &newPerson)

	newPerson.Id = len(people) + 1
	log.Println(newPerson)
	people = append(people, newPerson)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
		return
	}

	for _, Person := range people {
		if Person.Id == personId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Person)
		}
	}
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
		return
	}

	for index, Person := range people {
		if Person.Id == personId {
			people = append(people[:index], people[index+1:]...)
			fmt.Fprintf(w, "Person with Id %v has been removed successfully", personId)
		}
	}
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId, err := strconv.Atoi(vars["id"])
	var updatePerson Person

	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &updatePerson)

	for index, Person := range people {
		if Person.Id == personId {
			people = append(people[:index], people[index+1:]...)
			updatePerson.Id = personId
			people = append(people, updatePerson)

			fmt.Fprintf(w, "The Person with id %v has been updated successfully", personId)
		}
	}
}

type OutPerson struct {
	Id          int
	Name        string
	BasicPay    float64
	Departments []string
	HRA         float64
	GrossPay    float64
}

func (person Person) calculateGrossPay() OutPerson {
	hraPercentage := 0.0
	for _, dept := range person.Departments {
		if hra, ok := hraModifiers[dept]; ok {
			hraPercentage += hra
		}
	}
	HRA := person.BasicPay * hraPercentage
	GrossPay := person.BasicPay + HRA
	return OutPerson{person.Id, person.Name, person.BasicPay, person.Departments, HRA, GrossPay}
}

func getAllGrossPay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var outPeople []OutPerson
	for _, person := range people {
		outPeople = append(outPeople, person.calculateGrossPay())
	}
	json.NewEncoder(w).Encode(outPeople)
}
func getGrossPay(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
		return
	}

	for _, person := range people {
		if person.Id == personId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(person.calculateGrossPay())
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people", createPerson).Methods("POST")
	router.HandleFunc("/people/{id:[0-9]+}", getPerson).Methods("GET")
	router.HandleFunc("/people/{id:[0-9]+}", deletePerson).Methods("DELETE")
	router.HandleFunc("/people/{id:[0-9]+}", updatePerson).Methods("PUT")
	router.HandleFunc("/people/grosspay", getAllGrossPay).Methods("GET")
	router.HandleFunc("/people/{id:[0-9]+}/grosspay", getGrossPay).Methods("GET")
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
