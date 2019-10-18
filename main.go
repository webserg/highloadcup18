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
	byteValue, err := d.ReadFile("C:/Users/webse/go/src/github.com/webserg/highloadcup18/readData/datatest.1.json")
	check(err)
	accounts, err = d.ReadData(byteValue)
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
	params := []string{"sex_eq", "email_domain", "email_lt", "email_gt", "status_eq", "status_neq", "fname_eq", "fname_any", "fname_null", "sname_eq"}
	query := req.URL.Query()
	log.Println(query)
	log.Println(len(query))
	filtersArray := make([]func(d.Account) bool, 0)
	for _, k := range params {
		v, exists := query[k]
		if exists {
			filterValue := v[0]
			log.Println(k + "=" + filterValue)
			switch k {

			case "sex_eq":
				filtersArray = append(filtersArray, func(s d.Account) bool { return s.Sex == filterValue })

			case "email_domain":
				filtersArray = append(filtersArray, func(s d.Account) bool { return strings.HasSuffix(s.Email, filterValue) })

			case "email_lt":
				filtersArray = append(filtersArray, func(s d.Account) bool { return s.Email < filterValue })

			case "email_gt":
				filtersArray = append(filtersArray, func(s d.Account) bool { return s.Email > filterValue })

			case "status_eq":
				filtersArray = append(filtersArray, func(s d.Account) bool { return s.Status == filterValue })

			case "status_neq":
				filtersArray = append(filtersArray, func(s d.Account) bool { return s.Status != filterValue })

			case "fname_eq":
				filtersArray = append(filtersArray, func(s d.Account) bool { return s.Fname == filterValue })
			
			case "sname_eq":
				filtersArray = append(filtersArray, func(s d.Account) bool { return s.Sname == filterValue })


			case "fname_any":
				names := strings.Split(filterValue,",")
				filtersArray = append(filtersArray, func(s d.Account) bool { return index(names, s.Fname)!=-1 })

			case "fname_null":
				if filterValue != "0" && filterValue != "1" {
					res.WriteHeader(http.StatusBadRequest)
					log.Panic("fname_null(0,1) doesn't have parameter " + filterValue)
					return
				}
				filtersArray = append(filtersArray, func(s d.Account) bool {
					if filterValue == "1" {
						return s.Fname == ""
					}
					return len(s.Fname) > 0

				})
			}
		}
	}
	finalArray := filter2(filtersArray)
	fmt.Println(finalArray)
	if len(finalArray) == 0 {
		res.WriteHeader(http.StatusNotFound)
		log.Println("nothing is found")
		return
	}
	json.NewEncoder(res).Encode(d.Accounts{Accounts: finalArray})
}

func filter2(filtersArray []func(d.Account) bool) (ret []d.Account) {
	if len(filtersArray) > 0 {
		log.Println(filtersArray)
		for _, s := range accounts.Accounts {
			log.Println(s)
			res := true
			for _, filter := range filtersArray {
				log.Println(s)
				if !filter(s) {
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

func index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

func check(e error) {
	if e != nil {
		log.Panic(e)
		panic(e)
	}
}

func checkHTTP(res http.ResponseWriter, req *http.Request, e error) {
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
