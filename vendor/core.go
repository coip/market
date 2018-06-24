package vendor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var store Store

//Input: ItemCode; Maps to corresponding value based on map name
var names map[string]string
var prices map[string]int32
var products map[string]Product

var deals map[string]Special

func Init(offeringslink, specialslink string) {

	names = make(map[string]string)
	prices = make(map[string]int32)

	products = make(map[string]Product)
	deals = make(map[string]Special)

	store.Items = getProducts(offeringslink)
	store.Deals = getSpecials(specialslink)

}

func getProducts(link string) []Product {

	log.Printf("getting gist [%s] for product inventory offerings data", link)
	res, _ := http.Get(link)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	jsonstr := string(body)
	log.Println(jsonstr)
	var inventory []Product
	json.Unmarshal(body, &inventory)

	for _, element := range inventory {
		products[element.ID] = element
		names[element.ID] = element.Name
		price, _ := strconv.Atoi(stripCurrencyChars(element.Price))
		log.Println("price:", price, " element.Price:", element.Price)
		prices[element.ID] = int32(price)
	}

	return inventory
}

func getSpecials(link string) (specials []Special) {

	log.Printf("getting gist [%s] for product inventory specials data", link)
	res, _ := http.Get(link)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	jsonstr := string(body)
	log.Println(jsonstr)
	json.Unmarshal(body, &specials)

	for id, element := range specials {
		log.Printf("special[%v]=%-v\n", id, element)
		deals[element.BuyProduct] = element
	}

	return
}

func stripCurrencyChars(in string) string {
	tmp := strings.Replace(in, "$", "", -1)
	return strings.Replace(tmp, ".", "", -1)
}

func nameLookup(productId string) string {
	return names[productId]
}

func priceLookup(productId string) int32 {
	return prices[productId]
}