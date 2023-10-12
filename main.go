package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type RoastName string

type order struct {
	CompanyName      string  `json:"companyName"`
	OrderNotes       string  `json:"orderNotes"`
	OrderNumber      string  `json:"orderNumber"`
	DeliveryLocation string  `json:"deliveryLocation,omitempty"`
	OrderTotal       float64 `json:"orderTotal"`
	OrderImage       string  `json:"orderImage"`
}

func submitOrder(w http.ResponseWriter, r *http.Request) {
	order := new(order)
	json.NewDecoder(r.Body).Decode(&order)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if strings.Trim(order.CompanyName, " ") == "" {
		http.Error(w, "Missing company name", http.StatusBadRequest)
		return
	}

	m := mail.NewV3Mail()

	from := mail.NewEmail("Eduardo", "eduardoapp.97@gmail.com") // Change to your verified sender
	m.SetFrom(from)
	m.SetTemplateID("d-a7ce9f0478ab48d3ae6cdfd4a043b3a7")

	tos := []*mail.Email{
		mail.NewEmail("Sissiboo Wholesale", "wholesale@sissiboocoffee.com"),
	}

	p := mail.NewPersonalization()
	p.AddTos(tos...)
	p.SetDynamicTemplateData("Company_Name", order.CompanyName)
	p.SetDynamicTemplateData("Delivery_Location", order.DeliveryLocation)
	p.SetDynamicTemplateData("Purchase_Order", order.OrderNumber)
	p.SetDynamicTemplateData("Order_Notes", order.OrderNotes)
	p.SetDynamicTemplateData("Order_Total", order.OrderTotal)

	m.AddPersonalizations(p)

	m.AddAttachment(&mail.Attachment{
		Content:     strings.Split(order.OrderImage, "base64,")[1],
		Filename:    fmt.Sprintf("order-%v", order.CompanyName),
		Type:        "image/png",
		Disposition: "attachment",
	})

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(m)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while sending email", http.StatusInternalServerError)
		return
	} else {

		w.WriteHeader(response.StatusCode)
		json.NewEncoder(w).Encode(response.Body)
	}

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4321"
	}

	fmt.Println("Staring sever in port", port)
	http.HandleFunc("/submit", corsHandler(submitOrder))

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func corsHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
