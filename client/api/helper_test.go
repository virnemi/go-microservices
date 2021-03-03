package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelperEncoder(t *testing.T) {
	helper := NewHelper()
	rr := httptest.NewRecorder()
	he := &HandlerError{
		Error: "Method not Allowed",
		Code:  http.StatusMethodNotAllowed,
	}

	encodeErr := helper.Encoder(rr, he)
	if encodeErr != nil {
		t.Fatal(encodeErr)
	}

	expected := `{"error":"Method not Allowed","code":405}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %s want %s",
			rr.Body.String(), expected)
	}
}

func TestHelperParseKey(t *testing.T) {
	req, err := http.NewRequest("GET", "/port?id=TEST", nil)
	if err != nil {
		t.Fatal(err)
	}
	helper := NewHelper()

	id, errId := helper.ParseKey(req, "id")
	if errId != nil {
		t.Fatal(errId)
	}
	expected := "TEST"
	if id != expected {
		t.Errorf("Helper returned unexpected key: got %s want %s",
			id, expected)
	}

	value, errId := helper.ParseKey(req, "value")
	expected = "value not found"
	if errId != nil && errId.Error() != expected {
		t.Errorf("Helper returned unexpected error: got '%s' want '%s'",
			errId.Error(), expected)
	}
	if value != "" {
		t.Errorf("Helper returned unexpected value: got '%s' want 'NULL'",
			value)
	}
}

func TestHelperParseIntKeys(t *testing.T) {
	req, err := http.NewRequest("GET", "/ports/list?start=100&size=150", nil)
	if err != nil {
		t.Fatal(err)
	}
	helper := NewHelper()

	start, errId := helper.ParseIntKey(req, "start")
	if errId != nil {
		t.Fatal(errId)
	}
	expected := 100
	if start != expected {
		t.Errorf("Helper returned unexpected key: got %d want %d",
			start, expected)
	}

	size, errId32 := helper.ParseInt32Key(req, "size")
	if errId32 != nil {
		t.Fatal(errId32)
	}
	expected32 := int32(150)
	if size != expected32 {
		t.Errorf("Helper returned unexpected value: got %d want %d",
			size, expected32)
	}
}
