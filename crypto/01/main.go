package main

import (
    "crypto/rand"
    "encoding/hex"
	"encoding/base64"
    "fmt"
)

func main() {
    // определяем слайс нужной длины
    b := make([]byte, 16)
    _, err := rand.Read(b) // записываем байты в массив b
    if err != nil {
        fmt.Printf("error: %v\n", err)
        return
    }
	fmt.Println(b)
    fmt.Println(hex.EncodeToString(b))

	s, err := RandBytes(32)
	if err != nil {
		fmt.Printf("error: %v\n", err)
        return
	}
	fmt.Printf("base64 encoding: %s\n", s)
} 

func RandBytes (n int) (string, error){
	b := make([]byte, n)
	_, err := rand.Read(b) // записываем байты в массив b
    if err != nil {
        fmt.Printf("error: %v\n", err)
        return "", err
    }
	s := base64.RawStdEncoding.EncodeToString(b)
	return s, nil
}

/*
Напишите функцию, которая будет генерировать массив случайных байт. 
Размер массива передаётся параметром. 
Функция должна возвращать массив в виде строки в кодировке base64.
Используйте пакет encoding/base64.
*/