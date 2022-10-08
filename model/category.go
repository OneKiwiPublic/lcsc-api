package model

import (
	"fmt"
	"log"

	"github.com/TylerBrock/colorjson"
)

type Category struct {
	CatalogId     int64  `json:"catalogId"`
	CatalogNameEn string `json:"catalogNameEn"`
	ProductNum    int64  `json:"productNum"`
}

type CategoryResponse []struct {
	CatalogId     int64      `json:"catalogId"`
	CatalogNameEn string     `json:"catalogNameEn"`
	ProductNum    int64      `json:"productNum"`
	ChildCatelogs []Category `json:"childCatelogs"`
}

func (c Category) ToString() string {
	res, err := colorjson.Marshal(c)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	return string(res)
}

func (c CategoryResponse) ToString() string {
	format := colorjson.NewFormatter()
	format.Indent = 4
	res, err := format.Marshal(c)
	if err != nil {
		log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	}
	fmt.Println(string(res))
	return string(res)
}
