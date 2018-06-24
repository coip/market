package vendor

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetName(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Body != nil {
			itemCode, _ := ioutil.ReadAll(r.Body)
			fmt.Fprint(w, nameLookup(string(itemCode)))
		} else {
			fmt.Fprintf(w, "pass item code in body for Name")
		}
	case "GET":
		if r.FormValue("code") != "" {
			fmt.Fprint(w, nameLookup(r.FormValue("code")))
		} else {
			fmt.Fprintf(w, "specify `code` in querystring")
		}
	}
}

func GetPrice(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Body != nil {
			itemCode, _ := ioutil.ReadAll(r.Body)
			fmt.Fprint(w, priceLookup(string(itemCode)))
		} else {
			fmt.Fprintf(w, "pass item code in body for Name")
		}
	case "GET":
		if r.FormValue("code") != "" {
			fmt.Fprint(w, priceLookup(r.FormValue("code")))
		} else {
			fmt.Fprintf(w, "specify `code` in querystring")
		}
	}
}

func ListVendorInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v", store)
}
func ListProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%-v", store.Items)
}
func ListSpecials(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%-v", store.Deals)
}
