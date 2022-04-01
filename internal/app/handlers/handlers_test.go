package handlers_test

import (
	"bytes"
	"github.com/SamoylenkoVadim/golang-practicum/internal/app/handlers"
	"github.com/SamoylenkoVadim/golang-practicum/internal/app/routers"
	"github.com/SamoylenkoVadim/golang-practicum/internal/app/storage"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, client *http.Client, method, path string, body string) (*http.Response, string) {

	req, err := http.NewRequest(method, ts.URL+path, bytes.NewReader([]byte(body)))
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)

	respBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	defer resp.Body.Close()

	return resp, string(respBody)
}

func TestRouter(t *testing.T) {

	s := storage.New()
	h, _ := handlers.New(s)
	router, _ := routers.NewRouter(h)
	ts := httptest.NewServer(router)
	defer ts.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	initStorageLink := "http://yandex.ru"
	s.Save("id1", initStorageLink)

	tests := []struct {
		name                string
		path                string
		method              string
		body                string
		statusWant          int
		checkBodyResp       bool
		checkLocationHeader bool
	}{
		{
			name:          "simple test #1",
			path:          "/",
			method:        "POST",
			body:          "http://google.com",
			statusWant:    201,
			checkBodyResp: true,
		},
		{
			name:       "simple test #2",
			path:       "/",
			method:     "POST",
			body:       "",
			statusWant: 400,
		},
		{
			name:       "simple test #3",
			path:       "/",
			method:     "POST",
			body:       "433443443433434",
			statusWant: 400,
		},
		{
			name:                "simple test #4",
			path:                "/id1",
			method:              "GET",
			body:                "",
			statusWant:          307,
			checkLocationHeader: true,
		},
		{
			name:       "simple test #5",
			path:       "/id1/123454323456",
			method:     "GET",
			body:       "",
			statusWant: 400,
		},
		{
			name:       "simple test #6",
			path:       "/",
			method:     "GET",
			body:       "",
			statusWant: 400,
		},
		{
			name:       "simple test #7",
			path:       "/unknown",
			method:     "GET",
			body:       "",
			statusWant: 400,
		},
		{
			name:       "simple test #8",
			path:       "/id1",
			method:     "PUT",
			body:       "",
			statusWant: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, body := testRequest(t, ts, client, tt.method, tt.path, tt.body)
			require.Equal(t, tt.statusWant, resp.StatusCode)

			if tt.checkBodyResp {
				_, err := url.ParseRequestURI(body)
				require.NoError(t, err)
			}

			if tt.checkLocationHeader {
				require.Equal(t, initStorageLink, resp.Header.Get("Location"))
			}
			defer resp.Body.Close()
		})
	}

}
