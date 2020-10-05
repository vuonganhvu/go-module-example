package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Users struct {
	Users []User `json:"users"`
}
type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age int `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter string `json:"twitter"`
}

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		panic(err.Error())
		return
	}
	defer jsonFile.Close()
	byteValues, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValues, &users)
	for _, user := range users.Users {
		fmt.Println("User Type: " + user.Type)
		fmt.Println("User Age: " + strconv.Itoa(user.Age))
		fmt.Println("User Name: " + user.Name)
		fmt.Println("Facebook Url: " + user.Social.Facebook)
	}
	for {
		rand.Seed(time.Now().Unix())

		//Generate a random character between lowercase a to z
		randomChar := 'A' + rune(rand.Intn(26))
		print(randomChar)
	}
}
