package main

import (
	"fmt"
	"log"
	"market/vendor"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	log.Println("Status Request Received")
	w.WriteHeader(200)
	fmt.Fprint(w, "OK\n")
}

func main() {

	var offeringslink string = "https://gist.githubusercontent.com/coip/ecf3fac70da64d32973fcb7347413891/raw"
	var specialslink string = "https://gist.githubusercontent.com/coip/411332578f63e9aee5e65a01328d9769/raw"

	vendor.Init(offeringslink, specialslink)

	http.HandleFunc("/", Status)

	//single-product lookup, mapped ItemCode -> Attribute. Supports [GET,POST]: QueryString `code` & posting a simple {-d "AP1"}
	http.HandleFunc("/api/products/getName", vendor.GetName)
	http.HandleFunc("/api/products/getPrice", vendor.GetPrice)

	//vendor-context list requests.
	http.HandleFunc("/api/vendors/getProducts", vendor.ListProducts)
	http.HandleFunc("/api/vendors/getSpecials", vendor.ListSpecials)
	http.HandleFunc("/api/vendors/getVendorInfo", vendor.ListVendorInfo)

	log.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
