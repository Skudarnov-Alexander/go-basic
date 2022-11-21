package main

import (
    "bytes"
    "crypto/md5"
    "crypto/rand"
    "fmt"
)

func main() {
    var (
        data  []byte         // слайс случайных байт
        hash1 []byte         // хеш с использованием интерфейса hash.Hash
        hash2 [md5.Size]byte // хеш, возвращаемый функцией md5.Sum
    )

	data = make([]byte, 512)

	n, err := rand.Read(data)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("num of bytes: %d\n", n)
	h := md5.New()
	h.Write(data)
	hash1 = h.Sum(nil)

	hash2 = md5.Sum(data)
    // допишите код
    // 1) генерация data длиной 512 байт
    // 2) вычисление hash1 с использованием md5.New
    // 3) вычисление hash2 функцией md5.Sum

    // ...

    // hash2[:] приводит массив байт к слайсу
    if bytes.Equal(hash1, hash2[:]) {
        fmt.Println("Всё правильно! Хеши равны")
    } else {
        fmt.Println("Что-то пошло не так")
    }
}

/*
Допишите программу, которая считает хеш MD5 случайной последовательности 512 байт. 
Один подсчёт сделайте с использованием интерфейса hash.Hash, а другой — функцией md5.Sum([]byte).
*/