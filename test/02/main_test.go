package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// User в системе.
type User struct {
    FirstName string
    LastName  string
}

// FullName возвращает фамилию и имя человека.
func (u User) FullName() string {
    return u.FirstName + " " + u.LastName
}

func main() {
    u := User{
        FirstName: "Misha",
        LastName:  "Popov",
    }

    fmt.Println(u.FullName())
} 

func TestFullName(t *testing.T){
	user := User{
		FirstName: "Саша",
		LastName:  "Скударнов",
	}
	name := user.FullName()
	assert.Equal(t, "Саша Скударнов", name)
		
}