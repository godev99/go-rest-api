package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Name struct {
	Type       string `json:"Type"`
	Kind       string `json:"Kind"`
	Expression string `json:"Expression"`
}

type Names []Name

var allNames = Names{
	{
		Type:       "KeyVault",
		Kind:       "PublicPaas",
		Expression: "regexatrouver",
	},
	{
		Type:       "WebApp",
		Kind:       "PrivatePaas",
		Expression: "regexatrouver",
	},
	{
		Type:       "VirtualMachine",
		Kind:       "Iaas",
		Expression: "regexatrouver",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createName(w http.ResponseWriter, r *http.Request) {
	var newName Name
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the resource type and its formula")
	}

	json.Unmarshal(reqBody, &newName)

	allNames = append(allNames, newName)

	// Return the 201 created status code
	w.WriteHeader(http.StatusCreated)
	// Return the newly created event
	json.NewEncoder(w).Encode(newName)
}

func getName(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	resourceType := mux.Vars(r)["Type"]

	// Get the details from an existing event
	// Use the blank identifier to avoid creating a value that will not be used
	for _, singleResource := range allNames {
		if singleResource.Type == resourceType {
			json.NewEncoder(w).Encode(singleResource)
		}
	}
}

func getNames(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(allNames)
}

func updateName(w http.ResponseWriter, r *http.Request) {
	resourceType := mux.Vars(r)["Type"]
	var updatedName Name

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the type and formula in order to update")
	}

	json.Unmarshal(reqBody, &updatedName)

	for i, singleResource := range allNames {
		if singleResource.Type == resourceType {
			singleResource.Expression = updatedName.Expression
			singleResource.Kind = updatedName.Kind
			allNames[i] = singleResource
			json.NewEncoder(w).Encode(singleResource)
		}
	}
}

func deleteName(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	resourceType := mux.Vars(r)["Type"]

	for i, singleResource := range allNames {
		if singleResource.Type == resourceType {
			allNames = append(allNames[:i], allNames[i+1:]...)
			fmt.Fprintf(w, "The name with Kind %v has been deleted successfully", resourceType)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/name", createName).Methods("POST")
	router.HandleFunc("/name/{Type}", getName).Methods("GET")
	router.HandleFunc("/names", getNames).Methods("GET")
	router.HandleFunc("/names/{Type}", updateName).Methods("PATCH")
	router.HandleFunc("/names/{Type}", deleteName).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
