package main

import (
	"net/http"
	"testing"

	//"net/http/httptest"
	"fmt"
	"io/ioutil"
)

func TestGetPeople(t *testing.T) {
	res, _ := http.Get("http://localhost:8080/people")
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Print("body")
	fmt.Printf("%s", b)
}
