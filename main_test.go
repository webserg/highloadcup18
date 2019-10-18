package main

import (
	"log"
	"net/http"
	"testing"

	// "encoding/json"
	//"net/http/httptest"
	"fmt"
	"io/ioutil"

	d "github.com/webserg/highloadcup18/readData"
)

func TestGetPeople(t *testing.T) {
	res, _ := http.Get("http://localhost:8080/accounts/filter?sex=f")
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	log.Printf("body")
	fmt.Printf("%s", b)
}

func TestGetFnameEq(t *testing.T) {
	res, _ := http.Get("http://localhost:8080/accounts/filter?fname_eq=maria")
	byteValue, _ := ioutil.ReadAll(res.Body)
	accounts, err := d.ReadData(byteValue)
	check(err)
	t.Log(accounts)
	if len(accounts.Accounts) != 1 {
		t.Errorf("Len was incorrect, got: %d, want: %d.", len(accounts.Accounts), 1)
	}
	res.Body.Close()
	t.Log("body")
	t.Logf("%s", accounts.Accounts)
}

func TestGetFnameAny(t *testing.T) {
	res, _ := http.Get("http://localhost:8080/accounts/filter?fname_any=maria,vlad,olav")
	byteValue, _ := ioutil.ReadAll(res.Body)
	accounts, err := d.ReadData(byteValue)
	check(err)
	t.Log(accounts)
	if len(accounts.Accounts) != 1 {
		t.Errorf("Len was incorrect, got: %d, want: %d.", len(accounts.Accounts), 1)
	}
	res.Body.Close()
	t.Log("body")
	t.Logf("%s", accounts.Accounts)
}

func TestGetFnameAnyFalse(t *testing.T) {
	res, _ := http.Get("http://localhost:8080/accounts/filter?fname_any=vsevolod,vlad,olav")
	byteValue, _ := ioutil.ReadAll(res.Body)
	accounts, err := d.ReadData(byteValue)
	check(err)
	t.Log(accounts)
	if len(accounts.Accounts) != 0 {
		t.Errorf("Len was incorrect, got: %d, want: %d.", len(accounts.Accounts), 1)
	}
	res.Body.Close()
	t.Log("body")
	t.Logf("%s", accounts.Accounts)
}

func TestGetSexEqM(t *testing.T) {
	res, _ := http.Get("http://localhost:8080/accounts/filter?sex_eq=m")
	byteValue, _ := ioutil.ReadAll(res.Body)
	accounts, err := d.ReadData(byteValue)
	check(err)
	if len(accounts.Accounts) != 2 {
		t.Errorf("Len was incorrect, got: %d, want: %d.", len(accounts.Accounts), 2)
	}
	res.Body.Close()
	t.Logf("%s", accounts.Accounts)
}

func TestGetSnameEq(t *testing.T) {
	res, _ := http.Get("http://localhost:8080/accounts/filter?sname_eq=Стаметаный")
	byteValue, _ := ioutil.ReadAll(res.Body)
	accounts, err := d.ReadData(byteValue)
	check(err)
	if len(accounts.Accounts) != 2 {
		t.Errorf("Len was incorrect, got: %d, want: %d.", len(accounts.Accounts), 2)
	}
	res.Body.Close()
	t.Logf("%s", accounts.Accounts)
}

func TestGetSnameEq2(t *testing.T) {
	res, _ := http.Get("http://localhost:8080/accounts/filter?sname_eq=jackson")
	byteValue, _ := ioutil.ReadAll(res.Body)
	accounts, err := d.ReadData(byteValue)
	check(err)
	if len(accounts.Accounts) != 1 {
		t.Errorf("Len was incorrect, got: %d, want: %d.", len(accounts.Accounts), 1)
	}
	res.Body.Close()
	t.Logf("%s", accounts.Accounts)
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
