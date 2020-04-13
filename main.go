package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Asset struct {
	Id string `json: Id`
	Name string `json:"Name"`
	Cnpj int `json:"Cnpj"`
}

var Assets []Asset


func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Diego Parra")
	fmt.Println("end hit: homepage")
}

func createNewAsset(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var asset Asset

	json.Unmarshal(reqBody, &asset)

	Assets = append(Assets, asset)

	json.NewEncoder(w).Encode(asset)
}

func returnAllAssets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("end hit: returnAllAssets")
	json.NewEncoder(w).Encode(Assets)
}

func returnSingleAsset( w http.ResponseWriter, r *http.Request) {
	fmt.Println("end git: returnSingleAsset")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, asset := range Assets {
		if asset.Id == key {
			json.NewEncoder(w).Encode(asset)
		}
	}
}

func deleteAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for index, asset := range Assets {
		if asset.Id == key {
			Assets = append(Assets[:index], Assets[index+1:]...)
		}
	}
}

func handleRequests()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/asset", createNewAsset).Methods("POST")
	router.HandleFunc("/asset/{id}", deleteAsset).Methods("DELETE")
	router.HandleFunc("/assets", returnAllAssets)
	router.HandleFunc("/asset/{id}", returnSingleAsset)
	log.Fatal(http.ListenAndServe(":8080", router))
}


func main()  {
	Assets = []Asset{
		Asset{Id: "1", Name: "ITSA4", Cnpj: 123456},
		Asset{Id: "2", Name: "UNIP6", Cnpj: 123789},
	}
	handleRequests()
}