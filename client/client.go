package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	res, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(body[:10])
	

	fmt.Println(res.Header.Get("Date"))
	
	client := http.Client{}

	res, err = client.Get("https://golang.org")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	JSONbody := []byte(`{"message": "Hello"}`)
	URL := "https://research.swtch.com/interfaces"

	request, err := http.NewRequest(http.MethodGet, URL , bytes.NewBuffer(JSONbody))
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("Accept", "application/json")

	

	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(requestDump))

	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileBody := &bytes.Buffer{}

	writer := multipart.NewWriter(fileBody)

	part, err := writer.CreateFormFile("upload file", "text.txt")
	if err != nil {
		fmt.Println(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
	}

	writer.Close()

	myRequest, err := http.NewRequest(http.MethodPost, URL, fileBody)
	if err != nil {
		fmt.Println(err)
	}
	// добавляем заголовок запроса
	myRequest.Header.Add("Content-Type", writer.FormDataContentType()) 


	requestDump, err = httputil.DumpRequest(myRequest, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(requestDump))












}