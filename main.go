package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var PORT int16 = 4321

type roastQty struct {
	OneLb  int `json:"oneLb"`
	TwoLb  int `json:"twoLb"`
	FiveLb int `json:"fiveLb"`
}

type order struct {
	CompanyName string              `json:"companyName"`
	OrderNotes  string              `json:"orderNotes"`
	OrderNumber string              `json:"orderNumber"`
	RoastQty    map[string]roastQty `json:"roastQty"`
}

type response struct {
	message string
}

func uploadImage(w http.ResponseWriter, r *http.Request) {

}

func submitOrder(w http.ResponseWriter, r *http.Request) {
	order := new(order)
	json.NewDecoder(r.Body).Decode(&order)

	if strings.Trim(order.CompanyName, " ") == "" {
		http.Error(w, "Missing company name", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Order created succesfully"))
}

func main() {
	fmt.Println("Staring sever in port", PORT)

	http.HandleFunc("/upload", uploadImage)
	http.HandleFunc("/submit", submitOrder)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}
