package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Transaction описывает запись блокчейна.
type Transaction struct {
    SequenceId  int `json:"-"`
    BlockNumber int `json:"block_number,omitempty"`
    Hash        string
    Parent      *Transaction `json:"parent,omitempty"`
    ExtData     []byte       `json:"ext_data,omitempty"`
    CreatedAt   time.Time    `json:"created_at"`
    receivedAt  time.Time
}

type TransactionNew Transaction


func (t Transaction) MarshalJSON() ([]byte, error) {
    // чтобы избежать рекурсии при json.Marshal, объявляем новый тип
    

    aliasValue := struct {
        TransactionAlias
        // переопределяем поле внутри анонимной структуры
        CreatedAt int64 `json:"created_at"`
    }{
        // встраиваем значение всех полей изначального объекта (embedding)
        TransactionAlias: TransactionAlias(t),
        // задаём значение для переопределённого поля
        CreatedAt: t.CreatedAt.Unix(),
    }

    return json.Marshal(aliasValue) // вызываем стандартный Marshal
}

type TransactionAlias Transaction

func main() {
    now := time.Now().UTC()

	parentTest := Transaction{
		SequenceId:  2,
		BlockNumber: 100,
		Hash:        "parent",
		CreatedAt:   now,
		receivedAt:  now.Add(10 * time.Millisecond),
	}

	txTest := Transaction{
		SequenceId:  4,
		BlockNumber: 20,
		Hash:        "test",
		Parent:      &parentTest,
		ExtData:     []byte{1, 2, 3},
		CreatedAt:   now.Add(1 * time.Second),
        receivedAt:  now.Add(1*time.Second + 10*time.Millisecond),
	}
	tNew := struct {
		TransactionNew
		CreatedAt string `json:"created_at"`
	}{
		TransactionNew: TransactionNew(txTest),
		CreatedAt:     "testId",
	}

	fmt.Printf("%+v\n\n", tNew)

    // создаём первую запись
    parentTx := Transaction{
        SequenceId: 1,
        Hash:       "0102AABA",
        CreatedAt:  now,
        receivedAt: now.Add(10 * time.Millisecond),
    }
    // у второй записи в качестве родителя указываем parentTx
    tx := Transaction{
        SequenceId:  2,
        BlockNumber: 1,
        Hash:        "0102AABB",
        Parent:      &parentTx,
        ExtData:     []byte{1, 2, 3},
        CreatedAt:   now.Add(1 * time.Second),
        receivedAt:  now.Add(1*time.Second + 10*time.Millisecond),
    }
    // преобразуем tx в JSON-формат
    txBz, err := json.Marshal(tx)
    if err != nil {
        panic(err)
    }
    // txBz — это []byte, поэтому приводим его к типу string для печати
    fmt.Println(string(txBz))


	testJSON, err := json.Marshal(tNew)
    if err != nil {
        panic(err)
    }
    // txBz — это []byte, поэтому приводим его к типу string для печати
    fmt.Println(string(testJSON))
}