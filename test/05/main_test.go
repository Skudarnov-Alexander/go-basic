package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserViewHandler(t *testing.T) {
	type want struct{
		code int
		response User
		contentType string
	}


	tests := []struct{
		name string
		users map[string]User
		want want
		request string
	}{
		{
			name: "statusOK: 200",
			users: map[string]User{
				"id1": {
				ID:        "id1",
				FirstName: "Alex",
				LastName:  "Skudarnov",
				},
				"id2": {
				ID:        "id2",
				FirstName: "Dima",
				LastName:  "Ivanov",
			},
		},
			want:   want{
				code: 200, 
				response: User{
					ID:        "id1",
					FirstName: "Alex",
					LastName:  "Skudarnov",
				}, 
				contentType: "application/json",
			},
			request: "/user?user_id=id1",
		},
		{
			name:    "statusBadRequest: 400",
			users:   map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Alex",
					LastName:  "Skudarnov",
					},
					"id2": {
					ID:        "id2",
					FirstName: "Dima",
					LastName:  "Ivanov",
				},
			},
			want:    want{
				code:        404,
				response:    User{},
				contentType: "text/plain; charset=utf-8",
			},
			request: "/user?user_id=id3",
		},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			request := httptest.NewRequest(http.MethodGet, tt.request, nil)

			w := httptest.NewRecorder()

			h := http.HandlerFunc(UserViewHandler(tt.users))

			h.ServeHTTP(w, request)

			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode)
			fmt.Printf("statusCode:%d\n", res.StatusCode)
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))

			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			
			fmt.Printf("body:%s\n", body)
			assert.NoError(t, err)

			var user User

			err = json.Unmarshal(body, &user)
			fmt.Printf("User:%v\n",user)
			fmt.Println(err)
			assert.NoError(t, err)

			assert.Equal(t, tt.want.response, user)

			

			
		})
	}
	
}