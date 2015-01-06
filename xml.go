package jsonar

/*
** Initially based on response here:
** stackoverflow.com/questions/25879570/how-to-convert-byte-xml-to-json-output-in-golang
**
*/

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type DataFormat struct {
	ProductList []struct {
		Sku      string `xml:"sku" json:"sku"`
		Quantity int    `xml:"quantity" json:"quantity"`
	} `xml:"Product" json:"products"`
}

func main() {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
        <ProductList>
            <Product>
                <sku>ABC123</sku>
                <quantity>2</quantity>
            </Product>
            <Product>
                <sku>ABC123</sku>
                <quantity>2</quantity>
            </Product>
        </ProductList>`)

	data := &DataFormat{}
	err := xml.Unmarshal(xmlData, data)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return
	}

	result, err := json.Marshal(data)
	if nil != err {
		fmt.Println("Error marshalling to JSON", err)
		return
	}

	fmt.Printf("%s\n", result)
}
