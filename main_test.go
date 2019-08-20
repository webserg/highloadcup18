package main

import (
	"net/http"
	"testing"
	"log"
	// "encoding/json"
	//"net/http/httptest"
	"fmt"
	"io/ioutil"
)

func TestGetPeople(t *testing.T) {
	res, _ := http.Get("http://localhost:8080/accounts/filter?sex=f")
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	log.Printf("body")
	fmt.Printf("%s", b)
}

func TestGetPersone(t *testing.T) {
	// res, _ := http.Get("http://localhost:8080/people/2")
	// var person Person
	// _ = json.NewDecoder(res.Body).Decode(&person)
	// res.Body.Close()
	// if person.ID != "2"  {
        // t.Errorf("got pesone.ID = %s; want persone.ID=2", person.ID)
	// }
	// log.Printf("%s", person.ID)
}
