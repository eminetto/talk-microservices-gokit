package user

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestHTTP(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8081", "caller", log.DefaultCaller)
	s := NewService()
	r := NewHttpServer(s, logger)
	srv := httptest.NewServer(r)

	for _, testcase := range []struct {
		method, url, body string
		want              int
	}{
		{"POST", srv.URL + "/v1/auth", `{"email": "eminetto@gmail.com", "password":"1234567"}`, http.StatusOK},
		{"GET", srv.URL + "/v1/auth", `{"email": "eminetto@gmail.com", "password":"1234567"}`, http.StatusMethodNotAllowed},
		{"POST", srv.URL + "/v1/auth", `{"email": "eminetto@gmail.com", "password":"invalid"}`, http.StatusNotFound},
		{"POST", srv.URL + "/v1/validate-token", `{"token": "invalid"}`, http.StatusUnauthorized},
	} {
		req, _ := http.NewRequest(testcase.method, testcase.url, strings.NewReader(testcase.body))
		resp, _ := http.DefaultClient.Do(req)
		if testcase.want != resp.StatusCode {
			t.Errorf("%s %s %s: want %d have %d", testcase.method, testcase.url, testcase.body, testcase.want, resp.StatusCode)
		}

	}
}
