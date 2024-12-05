package traefik_simpleredirect_test

import (
	"context"
	"github.com/daniels0056/traefik-simpleredirect"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimpleRedirect(t *testing.T) {
	argTo := "https://google.com"
	argCode := 302
	argRequestUrl := "https://bing.com"

	cfg := traefik_simpleredirect.CreateConfig()
	cfg.RedirectTo = argTo
	cfg.RedirectCode = argCode

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefik_simpleredirect.New(ctx, next, cfg, "simpleRedirect")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, argRequestUrl, nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)
	result := recorder.Result()

	if result.StatusCode != argCode {
		t.Fatal("Response did not return the correct status code")
	}

	if result.Header.Get("Location") != argTo {
		t.Fatal("Response did not contain a Location header or had an unexpected value")
	}
}
