package readData

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

// " Account smth smth"
type Account struct {
	// id    int    `json:"id"`
	Fname string `json:"fname"`
	Email string `json:"email"`
	// interests []string
	Status string
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

/*ReadFile return byte array with data*/
func ReadFile(filaName string) ([]byte, error) {
	// Open our jsonFile
	// jsonFile, err := os.Open("/home/webserg/data/test_accounts_291218/data/data/accounts_1.json")
	jsonFile, err := os.Open(filaName)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	return byteValue, err
}

/*ReadData read json data to array*/
func ReadData(byteValue []byte) (*Accounts, error) {

	// we initialize our Users array
	var accounts Accounts

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err := json.Unmarshal(byteValue, &accounts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(accounts.Accounts))

	return &accounts, nil
}
