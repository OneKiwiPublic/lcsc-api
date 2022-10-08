package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lcsc/model"
	"log"
	"net/http"
)

func GetCategoryList() {
	resp, err := http.Get(CATEGORY)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func FetchCategory() (model.CategoryResponse, error) {
	//We make HTTP request using the Get function
	resp, err := http.Get(CATEGORY)
	if err != nil {
		fmt.Println("ooopsss an error occurred, please try again")
		log.Fatal("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()
	//Create a variable of the same type as our model
	var response model.CategoryResponse
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	//Invoke the text output function & return it with nil as the error value
	return response, nil
}
