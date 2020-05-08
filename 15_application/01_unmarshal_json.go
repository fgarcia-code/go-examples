package main

import "fmt"
import "encoding/json"

// This is the people type on which the JSON will be "Unmarshaled", using it's json tags
type person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:lastname`
	Age       int    `json:age`
}

func main() {
	// This is the JSON string to be "Unmarshaled"
	json_string := `[{"firstname":"John","lastname":"Snow","age":22},{"firstname":"Daenerys","lastname":"Targaryen","age":21}]`

	// This is the json string converted to `[]byte`
	json_byte := []byte(json_string)

	// This is the variable wher the json will be "Unmarshaled"
	people := []person{}

	// The json is "Unmarshaled"
	err := json.Unmarshal(json_byte, &people)
	// Verifying if any errors
	if err != nil {
		fmt.Println(err)
	}

	// Printing the `[]person` "Unmarshaled" and it's type
	fmt.Println("\n This is the people 'Unmarshaled' in the `[]person")
	fmt.Println(people)
	fmt.Printf("\nThis is the type of people:\n%T\n", people)
}
