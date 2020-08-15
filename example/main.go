package main

import (
	"fmt"

	"github.com/akshaybharambe14/ijson"
)

var dataBytes = []byte(`
[
	{
	  "index": 0,
	  "friends": [
		{
		  "id": 0,
		  "name": "Justine Bird"
		},
		{
		  "id": 0,
		  "name": "Justine Bird"
		},
		{
		  "id": 1,
		  "name": "Marianne Rutledge"
		}
	  ]
	}
]
`)

// var data = []interface{}{
// 	map[string]interface{}{
// 		"index": 0,
// 		"friends": []interface{}{
// 			map[string]interface{}{
// 				"id":   0,
// 				"name": "Justine Bird",
// 			},
// 			map[string]interface{}{
// 				"id":   0,
// 				"name": "Justine Bird",
// 			},
// 			map[string]interface{}{
// 				"id":   1,
// 				"name": "Marianne Rutledge",
// 			},
// 		},
// 	},
// }

func main() {

	r := ijson.ParseBytes(dataBytes).
		GetP("#0.friends.#~name"). // list the friend names for 0th record -
		// []interface {}{"Justine Bird", "Justine Bird", "Marianne Rutledge"}

		Del("#0"). // delete 0th record
		// []interface {}{"Marianne Rutledge", "Justine Bird"}

		Set("tom", "#") // append "tom" in the list
		// // []interface {}{"Marianne Rutledge", "Justine Bird", "tom"}

	fmt.Printf("%#v\n", r.Value())
	// output: []interface {}{"Marianne Rutledge", "Justine Bird", "tom"}

	// returns error if the data type differs than the type expected by query
	fmt.Println(r.Set(1, "name").Error())
}
