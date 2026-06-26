package hltr

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		return errors.New("Expected Content-Type application/json got " + ct)
	}
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return err
	}
	return nil
}

func GetBodyAsString(w http.ResponseWriter, r *http.Request) (string, error) {
	if r.Body == nil || r.Body == http.NoBody {
		return "", nil
	}
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
