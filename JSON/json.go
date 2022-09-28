package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)
}
