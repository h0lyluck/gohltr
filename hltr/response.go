package hltr

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// SendStatus sends a status code to the response writer.
func SendStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

// SendError sends an error response with the given status code and HLTRError.
func SendError(w http.ResponseWriter, status int, err HLTRError) error {
	return SendJSON(w, status, err.ToHttpError())
}

// SendErrorString sends an error response with the given status code and message.
func SendErrorString(w http.ResponseWriter, status int, message string) error {
	return SendJSON(w, status, map[string]interface{}{
		"message": message,
	})
}

// SendJSON sends a JSON response with the given status code and data.
func SendJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// SendString sends a plain text response with the given status code and data.
func SendString(w http.ResponseWriter, status int, data string) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	_, err := w.Write([]byte(data))
	return err
}

// GetQueryParam retrieves a query parameter from the request URL.
func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

// GetQueryParamInt retrieves a query parameter as an int from the request URL.
func GetQueryParamInt(r *http.Request, key string) (int, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, nil
	}
	return strconv.Atoi(value)
}

// GetQueryParamInt8 retrieves a query parameter as an int8 from the request URL.
func GetQueryParamInt8(r *http.Request, key string) (int8, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, nil
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return int8(v), nil
}

// GetQueryParamInt16 retrieves a query parameter as an int16 from the request URL.
func GetQueryParamInt16(r *http.Request, key string) (int16, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, nil
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return int16(v), nil
}

// GetQueryParamInt32 retrieves a query parameter as an int32 from the request URL.
func GetQueryParamInt32(r *http.Request, key string) (int32, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, nil
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

// GetQueryParamInt64 retrieves a query parameter as an int64 from the request URL.
func GetQueryParamInt64(r *http.Request, key string) (int64, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, nil
	}
	return strconv.ParseInt(value, 10, 64)
}

// GetQueryParamFloat32 retrieves a query parameter as a float32 from the request URL.
func GetQueryParamFloat32(r *http.Request, key string) (float32, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, nil
	}
	v, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(v), nil
}

// GetQueryParamFloat64 retrieves a query parameter as a float64 from the request URL.
func GetQueryParamFloat64(r *http.Request, key string) (float64, error) {
	value := r.URL.Query().Get(key)
	if value == "" {
		return 0, nil
	}
	return strconv.ParseFloat(value, 64)
}
