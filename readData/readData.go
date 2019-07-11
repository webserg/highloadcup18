package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

// "Account smth smth"
type Account struct {
	// id    int    `json:"id"`
	Fname string `json:"fname"`
	Email string `json:"email"`
	// interests []string
	// status    string
	// premium   *Premium
	Sex string `json:"sex"`
	// phone     string
	// likes     []*Likes
	// birth     int64
	// city      string
	// country string
	// joined    int64
}

// type Likes struct {
// 	ts int64
// 	id int64
// }

// type Premium struct {
// 	start  string
// 	finish string
// }

// our main function
func main() {
	// Open our jsonFile
	// jsonFile, err := os.Open("/home/webserg/data/test_accounts_291218/data/data/accounts_1.json")
	jsonFile, err := os.Open("./datatest.1.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var accounts Accounts

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &accounts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(accounts.Accounts))
	fmt.Println(accounts.Accounts[0])
	//fmt.Println(accounts.Accounts)
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < 1; i++ {
		fmt.Println("fname: " + accounts.Accounts[i].Fname)
		fmt.Println("sex: " + accounts.Accounts[i].Sex)
		// fmt.Println(len(accounts.Accounts[i].interests))
		// fmt.Println("User id: " + strconv.Itoa(accounts.Accounts[i].id))
		//fmt.Println("premiun start: " + accounts.Accounts[i].premium.start)
	}

}
