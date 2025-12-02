# Video 12 | Response Code

## Response Code
- Dalam HTTP, terdapat yang namanya response code
- Response code merupakan representasi kode response
- Dari response code ini kita bisa melihat apakah sebuah request yang kita kirim itu sukses diproses oleh server atau gagal
- Ada banyak sekali response code yang bisa kita gunakan saat membuat web
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Status 

## Mengubah Response Code
- Secara default, jika kita tidak menyebutkan response code, maka response code nya adalah 200 OK
- Jika kita ingin mengubahnya, kita bisa menggunakan function ResponseWriter.WriteHeader(int)
- Semua data status code juga sudah disediakan di Go-Lang, jadi kita ingin, kita bisa gunakan variable yang sudah disediakan : https://github.com/golang/go/blob/master/src/net/http/status.go 

## Kode Program
```go
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
```