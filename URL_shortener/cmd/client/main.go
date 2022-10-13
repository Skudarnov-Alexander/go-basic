package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "strconv"
    "strings"
)

func main() {
	endpoint := "https://google.com"

	data := url.Values{}

	fmt.Println("Введите длинный URL")

	reader := bufio.NewReader(os.Stdin)

	long, err := reader.ReadString('\n')

	if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

	long = strings.TrimSuffix(long, "\n")

	data.Set("url", long)

	client := &http.Client{}

	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBufferString(data.Encode()))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	response, err := client.Do(request)

	if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    // печатаем код ответа
    fmt.Println("Статус-код ", response.Status)
    defer response.Body.Close()
    // читаем поток из тела ответа
    body, err := io.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    // и печатаем его
    fmt.Println(string(body))
}