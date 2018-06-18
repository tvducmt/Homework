// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import "encoding/json"
import "fmt"

// We'll use these two structs to demonstrate encoding and
// decoding of custom types below.
type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// In the examples above we always used bytes and
	// strings as intermediates between the data and
	// JSON representation on standard out. We can also
	// stream JSON encodings directly to `os.Writer`s like
	// `os.Stdout` or even HTTP response bodies.
	// enc := json.NewEncoder(os.Stdout)
	// d := map[string]int{"apple": 5, "lettuce": 7}
	// enc.Encode(d)
}
