package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")

	if query == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Name is empty")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello %s", query)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080?name=Thomas", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
