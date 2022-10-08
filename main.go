package main

import (
	"fmt"
	"lcsc/api"
	"log"

	"github.com/TylerBrock/colorjson"
)

func main() {

	category, err := api.FetchCategory()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(category[0])
		fmt.Println("======")
		//category.ToString()
		//j, _ := json.Marshal(category[0])
		fmt.Println(category.ToString())
		format := colorjson.NewFormatter()
		format.Indent = 4
		res, err := format.Marshal(category[0])
		if err != nil {
			log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
		}
		fmt.Println(string(res))
	}
}
