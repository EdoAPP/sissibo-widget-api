package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sendgrid/sendgrid-go"
	sgMail "github.com/sendgrid/sendgrid-go/helpers/mail"
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
	EmailAddress     string    `json:"emailAddress"`
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

	_, err := sgMail.ParseEmail(order.EmailAddress)

	if strings.Trim(order.EmailAddress, " ") == "" || err != nil {
		http.Error(w, "Missing or invalid email address", http.StatusBadRequest)
		return
	}

	// sissiboo email
	m := sgMail.NewV3Mail()
	from := sgMail.NewEmail("Drew Tozer", "drewmorgantozer@gmail.com") // Change to your verified sender
	m.SetFrom(from)
	m.SetTemplateID("d-a9133530e5c54f2aac13f7462adedf07")

	p := buildPersonalizedEmail(*order)
	cp := buildPersonalizedEmail(*order)

	p.SetDynamicTemplateData("Subject", fmt.Sprintf("New order from %s", order.CompanyName))
	p.AddTos(
		sgMail.NewEmail("Eduardo Pacheco", "edo@shipyardsoftware.org"),
	)

	cp.SetDynamicTemplateData("Subject", "Sissiboo Coffee has received your wholesale order!")
	cp.AddTos(sgMail.NewEmail(order.CompanyName, order.EmailAddress))

	m.AddPersonalizations(p, cp)

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

func buildPersonalizedEmail(o order) *sgMail.Personalization {
	p := sgMail.NewPersonalization()

	p.SetDynamicTemplateData("Company_Name", o.CompanyName)
	p.SetDynamicTemplateData("Delivery_Location", o.DeliveryLocation)
	p.SetDynamicTemplateData("Order_Number", o.OrderNumber)
	p.SetDynamicTemplateData("Order_Notes", o.OrderNotes)
	p.SetDynamicTemplateData("Order_Total", o.OrderTotal)
	p.SetDynamicTemplateData("Invoice_Date", time.Now().Format("2006-01-02"))

	p.SetDynamicTemplateData("ge_1lb_gr", o.RoastsQty.Ge1LbGr)
	p.SetDynamicTemplateData("ge_1lb_wb", o.RoastsQty.Ge1LbWb)
	p.SetDynamicTemplateData("ge_2lb_gr", o.RoastsQty.Ge2lbGr)
	p.SetDynamicTemplateData("ge_2lb_wb", o.RoastsQty.Ge2lbWb)
	p.SetDynamicTemplateData("ge_5lb_gr", o.RoastsQty.Ge5LbGr)
	p.SetDynamicTemplateData("ge_5lb_wb", o.RoastsQty.Ge5LbWb)

	p.SetDynamicTemplateData("tw_1lb_gr", o.RoastsQty.Tw1LbGr)
	p.SetDynamicTemplateData("tw_1lb_wb", o.RoastsQty.Tw1LbWb)
	p.SetDynamicTemplateData("tw_2lb_gr", o.RoastsQty.Tw2lbGr)
	p.SetDynamicTemplateData("tw_2lb_wb", o.RoastsQty.Tw2lbWb)
	p.SetDynamicTemplateData("tw_5lb_gr", o.RoastsQty.Tw5LbGr)
	p.SetDynamicTemplateData("tw_5lb_wb", o.RoastsQty.Tw5LbWb)

	p.SetDynamicTemplateData("no_1lb_gr", o.RoastsQty.No1LbGr)
	p.SetDynamicTemplateData("no_1lb_wb", o.RoastsQty.No1LbWb)
	p.SetDynamicTemplateData("no_2lb_gr", o.RoastsQty.No2lbGr)
	p.SetDynamicTemplateData("no_2lb_wb", o.RoastsQty.No2lbWb)
	p.SetDynamicTemplateData("no_5lb_gr", o.RoastsQty.No5LbGr)
	p.SetDynamicTemplateData("no_5lb_wb", o.RoastsQty.No5LbWb)

	p.SetDynamicTemplateData("fs_1lb_gr", o.RoastsQty.Fs1LbGr)
	p.SetDynamicTemplateData("fs_1lb_wb", o.RoastsQty.Fs1LbWb)
	p.SetDynamicTemplateData("fs_2lb_gr", o.RoastsQty.Fs2lbGr)
	p.SetDynamicTemplateData("fs_2lb_wb", o.RoastsQty.Fs2lbWb)
	p.SetDynamicTemplateData("fs_5lb_gr", o.RoastsQty.Fs5LbGr)
	p.SetDynamicTemplateData("fs_5lb_wb", o.RoastsQty.Fs5LbWb)

	p.SetDynamicTemplateData("fb_1lb_gr", o.RoastsQty.Fb1LbGr)
	p.SetDynamicTemplateData("fb_1lb_wb", o.RoastsQty.Fb1LbWb)
	p.SetDynamicTemplateData("fb_2lb_gr", o.RoastsQty.Fb2lbGr)
	p.SetDynamicTemplateData("fb_2lb_wb", o.RoastsQty.Fb2lbWb)
	p.SetDynamicTemplateData("fb_5lb_gr", o.RoastsQty.Fb5LbGr)
	p.SetDynamicTemplateData("fb_5lb_wb", o.RoastsQty.Fb5LbWb)

	p.SetDynamicTemplateData("fbnd_1lb_gr", o.RoastsQty.Fbnd1LbGr)
	p.SetDynamicTemplateData("fbnd_1lb_wb", o.RoastsQty.Fbnd1LbWb)
	p.SetDynamicTemplateData("fbnd_2lb_gr", o.RoastsQty.Fbnd2lbGr)
	p.SetDynamicTemplateData("fbnd_2lb_wb", o.RoastsQty.Fbnd2lbWb)
	p.SetDynamicTemplateData("fbnd_5lb_gr", o.RoastsQty.Fbnd5LbGr)
	p.SetDynamicTemplateData("fbnd_5lb_wb", o.RoastsQty.Fbnd5LbWb)

	total1lb := o.RoastsQty.Ge1LbGr + o.RoastsQty.Ge1LbWb + o.RoastsQty.Tw1LbGr + o.RoastsQty.Tw1LbWb + o.RoastsQty.No1LbGr + o.RoastsQty.No1LbWb + o.RoastsQty.Fs1LbGr + o.RoastsQty.Fs1LbWb + o.RoastsQty.Fb1LbGr + o.RoastsQty.Fb1LbWb + o.RoastsQty.Fbnd1LbGr + o.RoastsQty.Fbnd1LbWb
	total2lb := o.RoastsQty.Ge2lbGr + o.RoastsQty.Ge2lbWb + o.RoastsQty.Tw2lbGr + o.RoastsQty.Tw2lbWb + o.RoastsQty.No2lbGr + o.RoastsQty.No2lbWb + o.RoastsQty.Fs2lbGr + o.RoastsQty.Fs2lbWb + o.RoastsQty.Fb2lbGr + o.RoastsQty.Fb2lbWb + o.RoastsQty.Fbnd2lbGr + o.RoastsQty.Fbnd2lbWb
	total5lb := o.RoastsQty.Ge5LbGr + o.RoastsQty.Ge5LbWb + o.RoastsQty.Tw5LbGr + o.RoastsQty.Tw5LbWb + o.RoastsQty.No5LbGr + o.RoastsQty.No5LbWb + o.RoastsQty.Fs5LbGr + o.RoastsQty.Fs5LbWb + o.RoastsQty.Fb5LbGr + o.RoastsQty.Fb5LbWb + o.RoastsQty.Fbnd5LbGr + o.RoastsQty.Fbnd5LbWb

	p.SetDynamicTemplateData("total_1lb", total1lb)
	p.SetDynamicTemplateData("total_2lb", total2lb)
	p.SetDynamicTemplateData("total_5lb", total5lb)

	return p

}
