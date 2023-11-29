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

type RoastsQty struct {
	Ge1LbGr   int `json:"ge1LbGr"`
	Ge1LbWb   int `json:"ge1LbWb"`
	Ge2lbGr   int `json:"ge2LbGr"`
	Ge2lbWb   int `json:"ge2LbWb"`
	Ge5LbGr   int `json:"ge5LbGr"`
	Ge5LbWb   int `json:"ge5LbWb"`
	Tw1LbGr   int `json:"tw1LbGr"`
	Tw1LbWb   int `json:"tw1LbWb"`
	Tw2lbGr   int `json:"tw2LbGr"`
	Tw2lbWb   int `json:"tw2LbWb"`
	Tw5LbGr   int `json:"tw5LbGr"`
	Tw5LbWb   int `json:"tw5LbWb"`
	No1LbGr   int `json:"no1LbGr"`
	No1LbWb   int `json:"no1LbWb"`
	No2lbGr   int `json:"no2LbGr"`
	No2lbWb   int `json:"no2LbWb"`
	No5LbGr   int `json:"no5LbGr"`
	No5LbWb   int `json:"no5LbWb"`
	Fs1LbGr   int `json:"fs1LbGr"`
	Fs1LbWb   int `json:"fs1LbWb"`
	Fs2lbGr   int `json:"fs2LbGr"`
	Fs2lbWb   int `json:"fs2LbWb"`
	Fs5LbGr   int `json:"fs5LbGr"`
	Fs5LbWb   int `json:"fs5LbWb"`
	Fb1LbGr   int `json:"fb1LbGr"`
	Fb1LbWb   int `json:"fb1LbWb"`
	Fb2lbGr   int `json:"fb2LbGr"`
	Fb2lbWb   int `json:"fb2LbWb"`
	Fb5LbGr   int `json:"fb5LbGr"`
	Fb5LbWb   int `json:"fb5LbWb"`
	Fbnd1LbGr int `json:"fbnd1LbGr"`
	Fbnd1LbWb int `json:"fbnd1LbWb"`
	Fbnd2lbGr int `json:"fbnd2LbGr"`
	Fbnd2lbWb int `json:"fbnd2LbWb"`
	Fbnd5LbGr int `json:"fbnd5LbGr"`
	Fbnd5LbWb int `json:"fbnd5LbWb"`
}

type order struct {
	CompanyName      string    `json:"companyName"`
	OrderNotes       string    `json:"orderNotes"`
	OrderNumber      string    `json:"orderNumber"`
	DeliveryLocation string    `json:"deliveryLocation,omitempty"`
	OrderTotal       float64   `json:"orderTotal"`
	RoastsQty        RoastsQty `json:"roastsQty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	order := new(order)
	json.NewDecoder(r.Body).Decode(&order)

	if r.Method == http.MethodOptions {
		return
	}

	if strings.Trim(order.CompanyName, " ") == "" {
		http.Error(w, "Missing company name", http.StatusBadRequest)
		return
	}

	m := mail.NewV3Mail()

	from := mail.NewEmail("Drew Tozer", "drewmorgantozer@gmail.com") // Change to your verified sender
	m.SetFrom(from)
	m.SetTemplateID("d-a9133530e5c54f2aac13f7462adedf07")

	tos := []*mail.Email{
		mail.NewEmail("Sissiboo Wholesale", "edo@shipyardsoftware.org"),
	}

	p := mail.NewPersonalization()
	p.AddTos(tos...)

	p.SetDynamicTemplateData("Company_Name", order.CompanyName)
	p.SetDynamicTemplateData("Delivery_Location", order.DeliveryLocation)
	p.SetDynamicTemplateData("Purchase_Order", order.OrderNumber)
	p.SetDynamicTemplateData("Order_Notes", order.OrderNotes)
	p.SetDynamicTemplateData("Order_Total", order.OrderTotal)

	p.SetDynamicTemplateData("ge_1lb_gr", order.RoastsQty.Ge1LbGr)
	p.SetDynamicTemplateData("ge_1lb_wb", order.RoastsQty.Ge1LbWb)
	p.SetDynamicTemplateData("ge_2lb_gr", order.RoastsQty.Ge2lbGr)
	p.SetDynamicTemplateData("ge_2lb_wb", order.RoastsQty.Ge2lbWb)
	p.SetDynamicTemplateData("ge_5lb_gr", order.RoastsQty.Ge5LbGr)
	p.SetDynamicTemplateData("ge_5lb_wb", order.RoastsQty.Ge5LbWb)

	p.SetDynamicTemplateData("tw_1lb_gr", order.RoastsQty.Tw1LbGr)
	p.SetDynamicTemplateData("tw_1lb_wb", order.RoastsQty.Tw1LbWb)
	p.SetDynamicTemplateData("tw_2lb_gr", order.RoastsQty.Tw2lbGr)
	p.SetDynamicTemplateData("tw_2lb_wb", order.RoastsQty.Tw2lbWb)
	p.SetDynamicTemplateData("tw_5lb_gr", order.RoastsQty.Tw5LbGr)
	p.SetDynamicTemplateData("tw_5lb_wb", order.RoastsQty.Tw5LbWb)

	p.SetDynamicTemplateData("no_1lb_gr", order.RoastsQty.No1LbGr)
	p.SetDynamicTemplateData("no_1lb_wb", order.RoastsQty.No1LbWb)
	p.SetDynamicTemplateData("no_2lb_gr", order.RoastsQty.No2lbGr)
	p.SetDynamicTemplateData("no_2lb_wb", order.RoastsQty.No2lbWb)
	p.SetDynamicTemplateData("no_5lb_gr", order.RoastsQty.No5LbGr)
	p.SetDynamicTemplateData("no_5lb_wb", order.RoastsQty.No5LbWb)

	p.SetDynamicTemplateData("fs_1lb_gr", order.RoastsQty.Fs1LbGr)
	p.SetDynamicTemplateData("fs_1lb_wb", order.RoastsQty.Fs1LbWb)
	p.SetDynamicTemplateData("fs_2lb_gr", order.RoastsQty.Fs2lbGr)
	p.SetDynamicTemplateData("fs_2lb_wb", order.RoastsQty.Fs2lbWb)
	p.SetDynamicTemplateData("fs_5lb_gr", order.RoastsQty.Fs5LbGr)
	p.SetDynamicTemplateData("fs_5lb_wb", order.RoastsQty.Fs5LbWb)

	p.SetDynamicTemplateData("fb_1lb_gr", order.RoastsQty.Fb1LbGr)
	p.SetDynamicTemplateData("fb_1lb_wb", order.RoastsQty.Fb1LbWb)
	p.SetDynamicTemplateData("fb_2lb_gr", order.RoastsQty.Fb2lbGr)
	p.SetDynamicTemplateData("fb_2lb_wb", order.RoastsQty.Fb2lbWb)
	p.SetDynamicTemplateData("fb_5lb_gr", order.RoastsQty.Fb5LbGr)
	p.SetDynamicTemplateData("fb_5lb_wb", order.RoastsQty.Fb5LbWb)

	p.SetDynamicTemplateData("fbnd_1lb_gr", order.RoastsQty.Fbnd1LbGr)
	p.SetDynamicTemplateData("fbnd_1lb_wb", order.RoastsQty.Fbnd1LbWb)
	p.SetDynamicTemplateData("fbnd_2lb_gr", order.RoastsQty.Fbnd2lbGr)
	p.SetDynamicTemplateData("fbnd_2lb_wb", order.RoastsQty.Fbnd2lbWb)
	p.SetDynamicTemplateData("fbnd_5lb_gr", order.RoastsQty.Fbnd5LbGr)
	p.SetDynamicTemplateData("fbnd_5lb_wb", order.RoastsQty.Fbnd5LbWb)

	total1lb := order.RoastsQty.Ge1LbGr + order.RoastsQty.Ge1LbWb + order.RoastsQty.Tw1LbGr + order.RoastsQty.Tw1LbWb + order.RoastsQty.No1LbGr + order.RoastsQty.No1LbWb + order.RoastsQty.Fs1LbGr + order.RoastsQty.Fs1LbWb + order.RoastsQty.Fb1LbGr + order.RoastsQty.Fb1LbWb + order.RoastsQty.Fbnd1LbGr + order.RoastsQty.Fbnd1LbWb
	total2lb := order.RoastsQty.Ge2lbGr + order.RoastsQty.Ge2lbWb + order.RoastsQty.Tw2lbGr + order.RoastsQty.Tw2lbWb + order.RoastsQty.No2lbGr + order.RoastsQty.No2lbWb + order.RoastsQty.Fs2lbGr + order.RoastsQty.Fs2lbWb + order.RoastsQty.Fb2lbGr + order.RoastsQty.Fb2lbWb + order.RoastsQty.Fbnd2lbGr + order.RoastsQty.Fbnd2lbWb
	total5lb := order.RoastsQty.Ge5LbGr + order.RoastsQty.Ge5LbWb + order.RoastsQty.Tw5LbGr + order.RoastsQty.Tw5LbWb + order.RoastsQty.No5LbGr + order.RoastsQty.No5LbWb + order.RoastsQty.Fs5LbGr + order.RoastsQty.Fs5LbWb + order.RoastsQty.Fb5LbGr + order.RoastsQty.Fb5LbWb + order.RoastsQty.Fbnd5LbGr + order.RoastsQty.Fbnd5LbWb

	p.SetDynamicTemplateData("1lb_total", total1lb)
	p.SetDynamicTemplateData("2lb_total", total2lb)
	p.SetDynamicTemplateData("5lb_total", total5lb)

	m.AddPersonalizations(p)

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
