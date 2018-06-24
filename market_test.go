package main

import (
	"testing"

	"io/ioutil"
	"net/http"
	"strconv"
)

func TestGetName(t *testing.T) {
	r, err := http.Get("http://127.0.0.1:8080/api/products/getName?code=AP1")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	name, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	if string(name) != "Apples" {
		t.Errorf("getName is broken, got: %s, want: %s.", name, "Apples")
	}
}

func TestGetPrice(t *testing.T) {
	r, err := http.Get("http://127.0.0.1:8080/api/products/getPrice?code=AP1")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	price, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	val, _ := strconv.Atoi(string(price))
	if val != 600 {
		t.Errorf("getPrice is broken, got: %d, want: %d.", price, 600)
	}
}
