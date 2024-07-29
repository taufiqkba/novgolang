package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Using HTTP Request
var baseURL = "http://localhost:8080"

type studentAPI_Client struct {
	ID    string
	Name  string
	Grade int
}

func fetchUsers() ([]studentAPI_Client, error) {
	var err error
	var data []studentAPI_Client
	client := &http.Client{}

	req, err := http.NewRequest("GET", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// fetch User by ID
func fetchUserByID(ID string) (studentAPI_Client, error) {
	var err error
	var client = &http.Client{}
	var data studentAPI_Client

	var param = url.Values{}
	param.Set("id", ID)
	payload := bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	fmt.Println(request)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	//execute fetchUsers()
	users, err := fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}

	// execute fetchUserByID
	var user1, err1 = fetchUserByID("E001")
	if err != nil {
		fmt.Println("Error!", err1.Error())
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}
