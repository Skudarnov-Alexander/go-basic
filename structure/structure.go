package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	type Person struct {
		Name        string    `json:"Имя"`
		Email       string    `json:"Почта"`
		DateOfBirth time.Time `json:"-"`
	}

	user := Person{
		Name:  "Алекс",
		Email: "alex@yandex.ru",
	}

	userJSON, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", userJSON)

	rawdata := `{
		"header": {
			"code": 0,
			"message": ""
		},
		"data": [{
			"type": "user",
			"id": 100,
			"attributes": {
				"email": "bob@yandex.ru",
				"article_ids": [10, 11, 12]
			}
		}]
	}`

	
	type AttrType struct {
		Email       string	`json:"email"`
		Article_ids [3]int	`json:"article_ids"`
	}
	

	type ResponseJSON struct {
		Header struct {
			Code    int		`json:"code"`
			Message string	`json:"message,omitempty"`
		}	`json:"header"`
		Data []struct {
			TypeData   string 		`json:"type"`
			Id         int			`json:"id"`
			Attributes AttrType		`json:"attributes"`
		} `json:"data"`
	}

	var JSONdata ResponseJSON

	err = json.Unmarshal([]byte(rawdata), &JSONdata)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(JSONdata)

}
