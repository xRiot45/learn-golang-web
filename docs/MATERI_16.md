# Video 16 | Middleware

## Middleware
- Dalam pembuatan web, ada konsep yang bernama middleware atau filter atau interceptor
- Middleware adalah sebuah fitur dimana kita bisa menambahkan kode sebelum dan setelah sebuah handler di eksekusi

## Middleware di Go-Lang web
- Sayangnya, di Go-Lang web tidak ada middleware
- Namun karena struktur handler yang baik menggunakan interface, kita bisa membuat middleware sendiri menggunakan handler

## Error Handler
- Kadang middleware juga biasa digunakan untuk melakukan error handler
- Hal ini sehingga jika terjadi panic di Handler, kita bisa melakukan recover di middleware, dan mengubah panic tersebut menjadi error response
- Dengan ini, kita bisa menjaga aplikasi kita tidak berhenti berjalan

## Contoh kode program
```go
package learngolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After execute handler")
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)
		}
	}()
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```