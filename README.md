<p align="center">
    <img
        src="./assets/logo.png"
        width="196" height="239" border="0" alt="IJSON"
    >
    <br>
    <a href="https://pkg.go.dev/badge/github.com/akshaybharambe14/ijson">
        <img src="https://pkg.go.dev/badge/github.com/akshaybharambe14/ijson" alt="PkgGoDev">
    </a>
    <a href="https://github.com/akshaybharambe14/ijson/actions?query=workflow%3A%22Build+and+test%22">
        <img src="https://github.com/akshaybharambe14/ijson/workflows/Build%20and%20test/badge.svg" alt="Build and Test Status">
    </a>
    <a href="https://goreportcard.com/report/github.com/akshaybharambe14/ijson">
        <img src="https://goreportcard.com/badge/github.com/akshaybharambe14/ijson" alt="Go report">
    </a>
    <br>
    Query <b><i>I</i></b>nterface <b><i>JSON</i></b> and set or delete values easily
</p

<!--

[![PkgGoDev](https://pkg.go.dev/badge/github.com/akshaybharambe14/ijson)](https://pkg.go.dev/github.com/akshaybharambe14/ijson)
[![Build and Test Status](https://github.com/akshaybharambe14/ijson/workflows/Build%20and%20test/badge.svg)](https://github.com/akshaybharambe14/ijson/actions?query=workflow%3A%22Build+and+test%22)
[![PkgGoDev](https://goreportcard.com/badge/github.com/akshaybharambe14/ijson)](https://goreportcard.com/report/github.com/akshaybharambe14/ijson)

-->

**IJSON** is a small but effective utility to deal with **dynamic** or **unknown JSON structures** in [Go](https://golang.org). It's a helpful wrapper for navigating hierarchies of `map[string]interface{}` OR `[]interface{}`. It is the best solution for one time data access and manipulation.

Other libraries parse the whole json structure in their own format and again to interface if required, not suitable if you have `interface{}` as input and want the output in same format.

> **Note** - This is not a json parser. It just plays with raw interface data.

## Features

- Very fast in accessing and manipulating top level values.
- Avoids parsing whole JSON structure to intermediate format. Saves allocations.
- Easy API to perform **query**, **set** or **delete** operations on raw interface data.
- **One line syntax** to chain multiple operations together.

## Known limitations

- Not suitable if you want to perform multiple operations on same data.

## Getting started

### Installation

```sh
go get -u github.com/akshaybharambe14/ijson
```

### Usage and Example

This package provides two types of functions. The functions suffixed with `<action>P` accept a path separated by `"."`.

Ex. "#0.friends.#~name"

```go
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

func main() {
	r := ijson.ParseByes(dataBytes).
		GetP("#0.friends.#~name"). // list the friend names for 0th record -
		// []interface {}{"Justine Bird", "Justine Bird", "Marianne Rutledge"}

		Del("#0"). // delete 0th record
		// []interface {}{"Marianne Rutledge", "Justine Bird"}

		Set("tom", "#") // append "tom" in the list
		// []interface {}{"Marianne Rutledge", "Justine Bird", "tom"}

	fmt.Printf("%#v\n", r.Value())

	// returns error if the data type differs than the type expected by query
	fmt.Println(r.Set(1, "name").Error())
}

```

### Path syntax

IJSON follows a specific path syntax to access the data. The implementation sticks to the analogy that, user knows the path. So if caller wants to access an index, the underlying data must be an array otherwise, an error will be returned.

Use functions and methods suffixed by `P` to provide a `"."` separated path.

#### Get

```json
{
	"index": 0,
	"name": { "first": "Tom", "last": "Anderson" },
	"friends": [
		{ "id": 1, "name": "Justine Bird" },
		{ "id": 2, "name": "Justine Rutledge" },
		{ "id": 3, "name": "Marianne Rutledge" }
	]
}
```

Summary of get operations on above data.

```text
"name.last"    >> "Anderson"                          // GET "last" field from "name" object
"friends.#"    >> 3                                   // GET length of "friends" array
"friends.#~id" >> [ 1, 2, 3 ]                         // GET all values of "id" field from "friends" array
"friends.#0"   >> { "id": 1, "name": "Justine Bird" } // GET "0th" element from "friends" array
```

#### Set

Set overwrites the existing data. An error will be returned if the data does not match the query. If the data is `nil`, it will create the structure.

There is an alternative for datatype mismatch. Use `SetF` instead of `Set` function. It will **forcefully** replace the existing with provided.

Following path syntax sets "Anderson" as a value in empty structure.

```text
"name.last"    >> { "name": { "last": "Anderson" } }  // Create an object and SET value of "last" field in "name" object
"#2"           >> ["", "", "Anderson"]                // Create an array and SET value at "2nd" index
"friends.#"    >> { "friends": [ "Anderson" ] }       // Create an object and APPEND to "friends" array
```

#### Delete

While deleting at an index, you have two options. By default, deletes does not preserve order. This helps to save unnecessary allocations as it just replaces the data at given index with last element. Refer following syntax for details.

```json
{
	"index": 0,
	"friends": ["Justine Bird", "Justine Rutledge", "Marianne Rutledge"]
}
```

Summary of delete operations on above data.

```text
"index"        >> { "friends": [ "Justine Bird", "Justine Rutledge", "Marianne Rutledge" ] } // DELETE "index" field
"friends.#"    >> { "index": 0, "friends": [ "Justine Bird", "Justine Rutledge" ] }          // DELETE last element from "friends" array
"friends.#0"   >> { "index": 0, "friends": [ "Marianne Rutledge", "Justine Rutledge" ] }     // DELETE "0th" element from "friends" array WITHOUT preserving order
"friends.#~0"  >> { "index": 0, "friends": [ "Justine Rutledge", "Marianne Rutledge" ] }     // DELETE "0th" element from "friends" array WITH preserving order
```

### Operations chaining

You can chain multiple operations and check if it succeeds or fails.

```go
    r := ijson.New(data).Get("#0", "friends", "#~name").Del("#0").Set(value, "#")
    if r.Error() != nil {
        ...
    }

    // access value
    _ = r.Value()
```

### Parsing the json

This package uses standard library [encoding/json](https://golang.org/pkg/encoding/json/) as a json parser. We already have a very wide range of json parsers. I would recommend [GJSON](https://https://github.com/tidwall/gjson). It is probably the fastest, as far as I know.

See `ijson.ParseBytes()` and `ijson.Parse()` functions.

Please check following awesome projects, you might find a better match for you.

1. [GJSON](https://github.com/tidwall/gjson), [SJSON](https://github.com/tidwall/sjson)
2. [FASTJSON](https://github.com/valyala/fastjson)
3. [GABS](https://github.com/Jeffail/gabs)

## Contact

Akshay Bharambe [@akshaybharambe1](http://twitter.com/akshaybharambe1)

## License

IJSON source code is available under the MIT [License](/LICENSE).
