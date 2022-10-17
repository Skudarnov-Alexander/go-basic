package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func testRequest(t *testing.T, ts *httptest.Server, method, path string) (int, string) {
	req := httptest.NewRequest(method, ts.URL+path, nil)
	//require.NoError(t, err)
	req.RequestURI = ""
	fmt.Printf("request: %+v\n", req)
	

	resp, err := http.DefaultClient.Do(req)
	fmt.Println(err)
	require.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	defer resp.Body.Close()

	return resp.StatusCode, string(respBody)

}

func TestRouter(t *testing.T) {
	r := NewRouter()
	ts := httptest.NewServer(r)

	defer ts.Close()

	statusCode, body := testRequest(t, ts, http.MethodGet, "/cars/audi")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "brand:audi", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/cars/audi/a6")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "brand and model:audi-a6", body)

	statusCode, body = testRequest(t, ts, http.MethodGet, "/car/")
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Equal(t, "404 page not found\n", body)
}


/* 
http.NewRequest									httptest.NewRequest
&{												&{
	Method:GET 										Method:GET
	URL:http://127.0.0.1:51760/cars/audi 			URL:http://127.0.0.1:51933/cars/audi
	Proto:HTTP/1.1 									Proto:HTTP/1.1
	ProtoMajor:1 									ProtoMajor:1
	ProtoMinor:1 									ProtoMinor:1
	Header:map[] 									Header:map[]
	Body:<nil> 										Body:{}
	GetBody:<nil> 									GetBody:<nil>
	ContentLength:0 								ContentLength:0
	TransferEncoding:[] 							TransferEncoding:[]
	Close:false 									Close:false
	Host:127.0.0.1:51760 							Host:127.0.0.1:51933
	Form:map[] 										Form:map[]
	PostForm:map[] 									PostForm:map[]
	MultipartForm:<nil> 							MultipartForm:<nil>
	Trailer:map[] 									Trailer:map[]
	RemoteAddr: 									RemoteAddr:192.0.2.1:1234
	RequestURI: 									RequestURI:http://127.0.0.1:51933/cars/audi
	TLS:<nil> 										TLS:<nil>
	Cancel:<nil> 									Cancel:<nil>
	Response:<nil> 									Response:<nil>
	ctx:0xc0000a6008}								ctx:<nil>}

                   
*/