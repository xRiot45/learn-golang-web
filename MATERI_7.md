# Video 7 | Request

## Request
- Request adalah struct yang merepresentasikan HTTP Request yang dikirim oleh Web Browser
- Semua informasi request yang dikirim bisa kita dapatkan di Request
- Seperti, URL, http method, http header, http body, dan lain-lain

## Contoh Kode
```go
func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
		fmt.Fprintln(w, r.URL.Path)
		fmt.Fprintln(w, r.URL.RawPath)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```