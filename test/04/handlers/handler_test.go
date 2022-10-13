package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Skudarnov-Alexander/go_basic/test/04/handlers"
)

func TestStatusHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "first test",
			want: want{
				code:        200,
				response:    `{"status":"ok"}`,
				contentType: "application/json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/status", nil)

			w := httptest.NewRecorder()

			h := http.HandlerFunc(handlers.StatusHandler)

			h.ServeHTTP(w, request)
			res := w.Result()

			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code - %d, got - %d", tt.want.code, res.StatusCode)
			}

			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			if string(resBody) != tt.want.response {
				t.Errorf("Expected response - %s, got - %s", tt.want.response, w.Body.String())
			}

			if res.Header.Get("Content-Type") != tt.want.contentType {
				t.Errorf("Expected Content-Type - %s, got - %s", tt.want.contentType, res.Header.Get("Content-Type"))
			}

		})
	}
}
