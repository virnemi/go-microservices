package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../../models"
)

func TestPortRouteGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/port?id=TEST", nil)
	if err != nil {
		t.Fatal(err)
	}

	params, er := models.NewEnvParams(false)
	if er != nil {
		t.Fatal(er)
	}
	router := NewPortRoute(params)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(router.Get)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"Port TEST not found.","code":404}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %s want %s",
			rr.Body.String(), expected)
	}
}

func TestPortRouteList(t *testing.T) {
	req, err := http.NewRequest("GET", "/ports/list?start=0&size=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	params, er := models.NewEnvParams(false)
	if er != nil {
		t.Fatal(er)
	}
	router := NewPortRoute(params)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(router.List)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"rpc error: code = DeadlineExceeded desc = context deadline exceeded","code":500}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %s want %s",
			rr.Body.String(), expected)
	}
}

func TestPortRouteListStartNotInformed(t *testing.T) {
	req, err := http.NewRequest("GET", "/ports/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	params, er := models.NewEnvParams(false)
	if er != nil {
		t.Fatal(er)
	}
	router := NewPortRoute(params)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(router.List)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"start not found","code":400}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %s want %s",
			rr.Body.String(), expected)
	}
}

func TestPortRouteListSizeNotInformed(t *testing.T) {
	req, err := http.NewRequest("GET", "/ports/list?start=0", nil)
	if err != nil {
		t.Fatal(err)
	}

	params, er := models.NewEnvParams(false)
	if er != nil {
		t.Fatal(er)
	}
	router := NewPortRoute(params)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(router.List)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"size not found","code":400}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %s want %s",
			rr.Body.String(), expected)
	}
}
