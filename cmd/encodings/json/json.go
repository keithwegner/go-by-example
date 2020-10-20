package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Go offers built-in support for JSON encoding and decoding, including to and from built-in custom data types

// We'll use these two structs to demonstrate encoding and decoding of custom types below
type response1 struct {
	Page   int
	Fruits []string
	nuts   []string
}

// Only exported fields witll be encoded/decoded in JSON. Fields must start with capital letters to be exported.
// We can also use the "omitempty" property to ignore a field in our JSON output if its value is empty, otherwise
// we get the type's default value (e.g. int = 0)
type response2 struct {
	Page   int      `json:"page,omitempty"`
	Fruits []string `json:"fruits"`
}

func main() {
	// First we'll look at encoding basic data types to JSON strings. Here are examples for atomic values
	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(2.32)
	fmt.Println(string(floatB))

	stringB, _ := json.Marshal("gopher")
	fmt.Println(string(stringB))

	// And here are some for slices and maps, which encode to JSON arrays and objects as you'd expect
	sliceD := []string{"apple", "peach", "pear"}
	sliceB, _ := json.Marshal(sliceD)
	fmt.Println(string(sliceB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// The JSON package can automatically encode your custom data types. It will only include exported fields in the
	// encoded output and will by default use those names as JSON keys. Notice how `nuts` is not included.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
		nuts:   []string{"almonds"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("res1B=", string(res1B))

	// You can use tags on struct field declarations to customize the encoded JSON key names. Check the definition
	// of reponse2 above to see an example of such tags.
	res2D := &response2{
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("res2B=", string(res2B))

	// Now let's look at decoding JSON data into Go values.
	// Here's an example for a generic data structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON package can put the decoded data. This map[string]interface{}
	// will hold a map of strings to arbitrary data types.
	var dat map[string]interface{}

	// Here is the actual decoding, and a check for associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("dat  =", dat)

	// In order to use the values in the decoded map, we'll need to convert them to their appropriate type.
	// For example, here we convert the value in `num` to the expected `float64` type.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accessing nested data requires a series of conversions
	strings := dat["strs"].([]interface{})
	string1 := strings[0].(string)
	fmt.Println("string1=", string1)

	// We can also decode it into custom data types, such as response2 above.
	// This has the advantages of adding additional type-safety to our program and eliminating the need for type
	// for type assertions when accessing the decoded data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println("res  =", res)
	fmt.Println(res.Fruits[0])

	// In the example above, we always used bytes and strings as intermediates between the data and JSON
	// representation on standard out. We can also stream JSON encodings directory to os.Writers such as
	// os.Stdout or even HTTP response bodies

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// We can use maps to encode unstructured data. The keys must be strings, the values can be any serializable data.
	birds := map[string]interface{}{
		"sounds": map[string]string{
			"raven": "nevermore",
			"eagle": "sqwak",
		},
		"total birds": 2,
	}
	data, _ := json.Marshal(birds)
	fmt.Println("Birds=", string(data))

	// It is preferred to use structs when the data can be modeled.

	// Check out the JSON and Go blog post for more...
	// https://blog.golang.org/json

	// Other great examples:
	// https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/s
}
