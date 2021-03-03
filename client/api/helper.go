package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	ApplicationJSON = "application/json"
)

// Helper class to read and write JSON content into ResponseWriter.
type Helper struct {
}

// Helper constructor
func NewHelper() *Helper {
	helper := &Helper{}
	return helper
}

// Converts the provided interface into JSON content in ResponseWriter
func (helper *Helper) Encoder(w http.ResponseWriter, v interface{}) error {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}
	w.Header().Set("Content-Type", ApplicationJSON)

	return nil
}

// Reads JSON from the Body of ResponseWriter and converts this in the interface
func (helper *Helper) Decoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// Reads the value of the provided key from the Request. Returns error if this
// value doesn't exists in request
func (helper *Helper) ParseKey(r *http.Request, key string) (string, error) {
	txt := getFromRequest(r, key)
	if txt == "" {
		return "", fmt.Errorf("%s not found", key)
	}
	return txt, nil
}

// Reads the value of the provided key from the Request and converts this value
// to int. Returns error if this value doesn't exists in request
func (helper *Helper) ParseIntKey(r *http.Request, key string) (int, error) {
	txt := getFromRequest(r, key)
	if txt == "" {
		return 0, fmt.Errorf("%s not found", key)
	}
	return strconv.Atoi(txt)
}

// Reads the value of the provided key from the Request and converts this value
// to int32. Returns error if this value doesn't exists in request
func (helper *Helper) ParseInt32Key(r *http.Request, key string) (int32, error) {
	txt := getFromRequest(r, key)
	if txt == "" {
		return 0, fmt.Errorf("%s not found", key)
	}
	i64, err := strconv.ParseInt(txt, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i64), nil
}

// Reads the value of the provided key from the Request. It tries to get this from
// the QueryString or the Form, depending on the Request's method.
// Returns an empty string if an error occours when parsing the form.
func getFromRequest(r *http.Request, key string) string {
	if r.Method == http.MethodGet {
		return r.URL.Query().Get(key)
	} else {
		err := r.ParseForm()
		if err != nil {
			return ""
		}
		return r.Form.Get(key)
	}
}
