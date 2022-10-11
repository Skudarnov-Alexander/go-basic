package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)

	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://google.com", nil)

	client := &http.Client{}

	resp, _ := client.Do(req)

	fmt.Println(resp)
}