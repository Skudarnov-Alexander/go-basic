package main

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/go-resty/resty/v2"
)

type (
	User struct {
		ID       int     `json:"id"`
		Name     string  `json:"name"`
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Address  Address `json:"address"`
		Phone    string  `json:"phone"`
		Website  string  `json:"website"`
		Company  Company `json:"company"`
	}
	Address struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
	}

	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	}
)

func main() {
	client := resty.New().
		SetBaseURL("https://jsonplaceholder.typicode.com/").
		SetRetryCount(5).
		SetRetryWaitTime(10 * time.Second).
		SetRetryMaxWaitTime(20 * time.Second)

	users, err := getUsers(client)

	if err != nil {
		fmt.Println(err)
		return
	}

	sortUsers(users)

	for _, u := range users {
		fmt.Println(u.Name)
	}

}

func getUsers(client *resty.Client) ([]User, error) {
	var users []User

	resp, err := client.R().
		SetResult(&users).
		Get("/users")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("невозможно получить информацию о пользователях. статус код не равен 200")
	}

	return users, nil

}

func sortUsers(users []User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})

}

/*

Выберите любую библиотеку из списка и сделайте запросы к API: https://jsonplaceholder.typicode.com/.
Приложение должно уметь:
Выводить всех пользователей /users/ в виде JSON: GET https://jsonplaceholder.typicode.com/users.
Выводить пользователей в отсортированном виде по полю name, а не по id.
Для этого сделайте сортировку внутри приложения.

В случае ошибки от сервера повторять запрос 5 раз с интервалом 10 секунд.
После пятого неуспешного запроса выводить в консоль сообщение о невозможности сделать запрос.

*/
