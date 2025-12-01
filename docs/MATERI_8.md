# Video 8 | HTTP Test

## HTTP Test
- Go-Lang sudah menyediakan package khusus untuk membuat unit test terhadap fitur Web yang kita buat
- Semuanya ada di dalam package net/http/httptest https://golang.org/pkg/net/http/httptest/ 
- Dengan menggunakan package ini, kita bisa melakukan testing handler web di Go-Lang tanpa harus menjalankan aplikasi web nya
- Kita bisa langsung fokus terhadap handler function yang ingin kita test

## httptest.NewRequest()
- NewRequest(method, url, body) merupakan function yang digunakan untuk membuat http.Request
- Kita bisa menentukan method, url dan body yang akan kita kirim sebagai simulasi unit test
- Selain itu, kita juga bisa menambahkan informasi tambahan lainnya pada request yang ingin kita kirim, seperti header, cookie, dan lain-lain

## httptest.NewRecorder()
- httptest.NewRecorder() merupakan function yang digunakan untuk membuat ResponseRecorder
- ResponseRecorder merupakan struct bantuan untuk merekam HTTP response dari hasil testing yang kita lakukan

## Contoh Kode
```go
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	fmt.Println(response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)

}
```