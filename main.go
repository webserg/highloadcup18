package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	d "github.com/webserg/highloadcup18/readData"
)

var accounts *d.Accounts

func init() {
	fmt.Println("init")
	var err error
	accounts, err = d.ReadData()
	check(err)
	fmt.Println(accounts)
}

// our main function
func main() {
	// accounts, err := readData.ReadData()
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/accounts/filter", filterService)
	http.ListenAndServe(":8080", nil)
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func filterService(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s \n", req.URL.Path)
	if req.URL.Path != "/accounts/filter" {
		http.NotFound(res, req)
		log.Panic("path not found")
		return
	}
	params := [6]string{"sex_eq", "email_domain", "status", "fname", "sname", "phone"}
	query := req.URL.Query()
	log.Println(query)
	log.Println(len(query))
	testArray := make([]func(d.Account) bool, 0)
	for _, k := range params {
		v, exists := query[k]
		if exists {
			filterValue := v[0]
			switch k {

			case "sex_eq":
				log.Println(k + "=" + filterValue)
				testArray = append(testArray, func(s d.Account) bool { return s.Sex == filterValue })

			case "email_domain":
				log.Println(k + "=" + filterValue)
				testArray = append(testArray, func(s d.Account) bool { return strings.HasSuffix(s.Email, filterValue) })

			}
		}
	}
	S2 := filter2(testArray)
	fmt.Println(S2)
	json.NewEncoder(res).Encode(S2)
}

func filter2(testArray []func(d.Account) bool) (ret []d.Account) {
	if len(testArray) > 0 {
		log.Println(testArray)
		for _, s := range accounts.Accounts {
			log.Println(s)
			res := true
			for _, test := range testArray {
				log.Println(s)
				if !test(s) {
					res = false
					break
				}
			}
			if res {
				ret = append(ret, s)
			}
		}
	}
	return
}

func check(e error) {
	if e != nil {
		log.Panic("error")
		panic(e)
	}
}

func checkHttp(res http.ResponseWriter, req *http.Request, e error) {
	if e != nil {
		log.Panic("error")
		http.Error(res, "error", -1)
	}
}

func splitField(f string) (field string, predicate string, e error) {
	fieldPred := strings.Split(f, "_")
	if len(fieldPred) != 2 {
		return "", "", errors.New("field is wrong " + f)
	}
	field = fieldPred[0]
	predicate = fieldPred[1]
	return field, predicate, nil
}
