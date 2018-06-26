package vendor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"market/template"
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

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	var customerList = store.getCustomers()

	// Had better use buffer pool. Hero exports `GetBuffer` and `PutBuffer` for this.
	//
	// For convenience, hero also supports `io.Writer`. For example, you can also define
	// the function to `func UserList(userList []string, w io.Writer) (int, error)`,
	// and then:
	//
	//   template.UserList(userList, w)
	//
	template.CustomerList(customerList, w)

	//fmt.Fprintf(w, "%v", store.getCustomers())
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v", store.Baskets[r.FormValue("id")])
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

//POST w { "basket": "$BASKET_ID", "code": "$ITEM_ID"}
func AddItem(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		var requestedOrder Order
		jsonstr, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal([]byte(jsonstr), &requestedOrder)

		var b Basket
		if validatedBasket, ok := store.Baskets[requestedOrder.BasketName]; ok {
			//log.Println("return customer:", requestedOrder.BasketName)
			b = validatedBasket
		} else {
			log.Printf("###\tNEW CUSTOMER:%s\n", requestedOrder.BasketName)
			store.Baskets[requestedOrder.BasketName] = make(map[string]int32)
			b = store.Baskets[requestedOrder.BasketName]
		}

		if validatedProduct, ok := products[requestedOrder.ItemCode]; ok {
			b[validatedProduct.ID]++
			log.Printf("%s\tadded [%s] to their basket.\n", requestedOrder.BasketName, validatedProduct.ID)
		} else {
			w.WriteHeader(404)
			fmt.Fprintf(w, "product code[%s] not found", requestedOrder.ItemCode)
		}

	}
	fmt.Fprintf(w, "hit addItem w qs:[%s]", r.URL)
}

func ListBaskets(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n\n\n\nthere are", len(store.Baskets), "patrons; inspecting baskets")
	for key, _ := range store.Baskets {
		//log.Printf("Key:%v value:%v", key, _)
		log.Printf("Looking at %s's basket:\n", key)
		pennies := peek(store.Baskets[key])
		log.Printf("basket[customer=%s]:TTL:$%d.%.2d\n\n", key, pennies/100, pennies%100)
	}
	fmt.Fprintf(w, "done")
}
