package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    src := []byte("Здесь могло быть написано, чем Go лучше Rust. " +
        "Но после хеширования уже не прочитаешь.")

    h := sha256.New()
    h.Write(src)
    dst := h.Sum(nil)

    fmt.Printf("%x\n", dst)
	fmt.Printf("%d\n", h.Size())
	fmt.Printf("%d\n", h.BlockSize())

	hash := sha256.Sum256(src)
	fmt.Printf("%x\n", hash)
}