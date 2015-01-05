/* DISCLAIMER
** Under Development
** Do Not Use - Won't Work
*/

package jsonar

import (
	"encoding/json"
	"fmt"
)

type RawJSON []byte

func Decode(b RawJSON) {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
    check(err)

	return f.(map[string]interface{})

}

func (m *map[string]interface{}) Structify(s *struct) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	

}

func (m *map[string]interface{}) Stringify() string {

	return string(m)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
