package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const baseUrl string = "http://52.52.253.24:8888/orders"

type order struct {
	Id 			string `xml:"id,omitempty" json:"id"`
	Data		string `xml:"data,omitempty" json:"data"`
	CreatedAt	string `xml:"createdAt,omitempty" json:"created_at"`
	UpdatedAt	string `xml:"updatedAt,omitempty" json:"updated_at"`
}

func main() {
	router := mux.NewRouter()
	router.Use(globalSettings)
	router.HandleFunc("/orders/{id}", GetOrder).Methods("GET")
	router.HandleFunc("/orders", CreateOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func globalSettings(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Content-Type", "application/json")
			handler.ServeHTTP(writer, request)
		})
}

func GetOrder(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	url := baseUrl + "/" + params["id"]
	xmlResponse, xmlError := http.Get(url)

	if xmlResponse != nil {
		writer.WriteHeader(xmlResponse.StatusCode)
		if http.StatusOK == xmlResponse.StatusCode {
			responseOrder := xmlToJson(xmlResponse)
			_ = json.NewEncoder(writer).Encode(responseOrder)
		}
	} else if xmlError != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func CreateOrder(writer http.ResponseWriter, r *http.Request) {

	var order order
	_ = json.NewDecoder(r.Body).Decode(&order)
	xmlByteData, _ := xml.Marshal(order)

	xmlResponse, error := http.Post(baseUrl, "application/xml", bytes.NewReader(xmlByteData))

	if xmlResponse != nil {
		writer.WriteHeader(xmlResponse.StatusCode)
		if http.StatusOK == xmlResponse.StatusCode {
			responseOrder := xmlToJson(xmlResponse)
			_ = json.NewEncoder(writer).Encode(responseOrder)
		}
	} else if error != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func xmlToJson(response *http.Response) interface{} {
	defer response.Body.Close()
	xmlData, _ := ioutil.ReadAll(response.Body)
	var order order
	_ = xml.Unmarshal(xmlData, &order)
	return order
}

