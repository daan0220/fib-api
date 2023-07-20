package fib_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daan0220/fib-api/context"
	"github.com/daan0220/fib-api/presentation/fib"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name           string
		param          string
		wantStatusCode int
		wantResponse   fib.Response
		wantErr        bool
	}{
		{
			name:           "ValidInput",
			param:          "n=10",
			wantStatusCode: http.StatusOK,
			wantResponse: fib.Response{
				Result: "55",
			},
		},
		{
			name:           "MissingQueryParam",
			param:          "",
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   fib.Response{},
		},
		{
			name:           "InvalidQueryParam",
			param:          "n=abc",
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   fib.Response{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/?"+test.param, nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			echoCtx := context.Context{Context: ctx}

			if err := fib.Get(echoCtx); (err != nil) != test.wantErr {
				t.Errorf("Unexpected error occurred: %s", err)
				return
			}

			if rec.Code != test.wantStatusCode {
				t.Errorf("Controller.Get() statusCode = %v, wantStatusCode = %v", rec.Code, test.wantStatusCode)
				return
			}

			if test.wantStatusCode == http.StatusOK {
				var gotResponse fib.Response
				if err := json.Unmarshal(rec.Body.Bytes(), &gotResponse); err != nil {
					t.Error(err)
					return
				}
				if diff := cmp.Diff(gotResponse, test.wantResponse); diff != "" {
					t.Errorf("Controller.Get() diff(-gotResponse +wantResponse)\n%s", diff)
				}
			}
		})
	}
}
