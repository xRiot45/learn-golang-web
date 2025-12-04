# Video 15 | ServeFile

## ServeFile
- Kadang ada kasus misal kita hanya ingin menggunakan static file sesuai dengan yang kita inginkan
- Hal ini bisa dilakukan menggunakan function http.ServeFile()
- Dengan menggunakan function ini, kita bisa menentukan file mana yang ingin kita tulis ke http response

## Contoh kode program ServeFile
```go
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
```

## Go-Lang Embed
- Parameter function http.ServeFile hanya berisi string file name, sehingga tidak bisa menggunakan Go-Lang Embed
- Namun bukan berarti kita tidak bisa menggunakan Go-Lang embed, karena jika untuk melakukan load file, kita hanya butuh menggunakan package fmt dan ResponseWriter saja


## Contoh kode program
```go
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
```