package main

import (
	"errors"
	"fmt"
	"testing"

	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Relationship определяет положение в семье.
type Relationship string

// Возможные роли в семье.
const (
	Father      = Relationship("father")
	Mother      = Relationship("mother")
	Child       = Relationship("child")
	GrandMother = Relationship("grandMother")
	GrandFather = Relationship("grandFather")
)

// Family описывает семью.
type Family struct {
	Members map[Relationship]Person
}

// Person описывает конкретного человека в семье.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

var (
	// ErrRelationshipAlreadyExists возвращает ошибку, если роль уже занята.
	// Подробнее об ошибках поговорим в девятой теме: «Errors, log».
	ErrRelationshipAlreadyExists = errors.New("relationship already exists")
)

// AddNew добавляет нового члена семьи.
// Если в семье ещё нет людей, создаётся пустой map.
// Если роль уже занята, метод выдаёт ошибку.
func (f *Family) AddNew(r Relationship, p Person) error {
	if f.Members == nil {
		f.Members = map[Relationship]Person{}
	}
	if _, ok := f.Members[r]; ok {
		return ErrRelationshipAlreadyExists
	}
	f.Members[r] = p
	return nil
}

func main() {
	f := Family{}
	err := f.AddNew(Father, Person{
		FirstName: "Misha",
		LastName:  "Popov",
		Age:       56,
	})
	fmt.Println(f, err)

	err = f.AddNew(Father, Person{
		FirstName: "Drug",
		LastName:  "Mishi",
		Age:       57,
	})
	fmt.Println(f, err)
}

func TestAddNew(t *testing.T) {
	tests := []struct {
		Name  string
		Value Family
		Want  bool
	}{
		{
			Name:  "",
			Value: Family{},
			Want:  false,
		}, {
			Name: "",
			Value: Family{
				Members: map[Relationship]Person{
					"Father": {
						FirstName: "Alex",
						LastName:  "Sku",
						Age:       45,
					},
				},
			},
			Want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := tt.Value.AddNew("Mother", Person{
				FirstName: "Sveta",
				LastName:  "Ivanova",
				Age:       34,
			})
			fmt.Println(tt.Value)
			if !tt.Want {
				require.NoError(t, err)
				return
				
			}
			t.Errorf("Ошибка")
		})
	}

}
