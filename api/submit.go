package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type order struct {
	CompanyName      string  `json:"companyName"`
	OrderNotes       string  `json:"orderNotes"`
	OrderNumber      string  `json:"orderNumber"`
	DeliveryLocation string  `json:"deliveryLocation,omitempty"`
	OrderTotal       float64 `json:"orderTotal"`
	OrderImage       string  `json:"orderImage"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	order := new(order)
	json.NewDecoder(r.Body).Decode(&order)

	if strings.Trim(order.CompanyName, " ") == "" {
		http.Error(w, "Missing company name", http.StatusBadRequest)
		return
	}

	m := mail.NewV3Mail()

	from := mail.NewEmail("Drew Tozer", "drewmorgantozer@gmail.com") // Change to your verified sender
	m.SetFrom(from)
	m.SetTemplateID("d-a7ce9f0478ab48d3ae6cdfd4a043b3a7")

	tos := []*mail.Email{
		// mail.NewEmail("Sissiboo Wholesale", "wholesale@sissiboocoffee.com"),
		mail.NewEmail("Eduardo Pacheco", "eduardoapp.97@gmail.com"),
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
