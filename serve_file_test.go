package learngolangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/not-found.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/not-found.html
var resourcesNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, request, resourcesOk)
	} else {
		fmt.Fprint(writer, request, resourcesNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
