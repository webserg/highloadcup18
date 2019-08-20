package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/webserg/highloadcup18/readData"
)

var accounts *readData.Accounts

func init(){
	fmt.Println("init")
	var err error
	accounts, err = readData.ReadData()
	check(err)
	fmt.Println(accounts)
}

// our main function
func main() {
	// accounts, err := readData.ReadData()
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/accounts/filter", filter)
	http.ListenAndServe(":8080", nil)
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func filter(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s \n", req.URL.Path)
	if req.URL.Path != "/accounts/filter" {
		http.NotFound(res, req)
		return
	}
	params := [6]string{"sex", "email", "status", "fname", "sname", "phone"}
	query := req.URL.Query()
	for _, k := range params {
		v, exists := query[k]
		if exists {
			fmt.Printf("%s -> %s\n", k, v[0])
			switch k {
			case "sex":
				fmt.Println(v[0])
				fmt.Println(accounts)
			}
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
