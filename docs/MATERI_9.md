# Video 9 | Query Parameter

## Query Parameter
- Query parameter adalah salah satu fitur yang biasa kita gunakan ketika membuat web
- Query parameter biasanya digunakan untuk mengirim data dari client ke server
- Query parameter ditempatkan pada URL
- Untuk menambahkan query parameter, kita bisa menggunakan ?nama=value pada URL nya

## url.URL
- Dalam parameter Request, terdapat attribute URL yang berisi data url.URL
- Dari data URL ini, kita bisa mengambil data query parameter yang dikirim dari client dengan menggunakan method Query() yang akan mengembalikan map

## Kode test query parameter
```go
func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Thomas", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
```

## Multiple Query Parameter
- Dalam spesifikasi URL, kita bisa menambahkan lebih dari satu query parameter
- Ini cocok sekali jika kita memang ingin mengirim banyak data ke server, cukup tambahkan query parameter lainnya
- Untuk menambahkan query parameter, kita bisa gunakan tanda & lalu diikuti dengan query parameter berikutnya

## Kode multiple query parameter
```go
func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?first_name=Thomas&last_name=Alberto", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
```

## Multiple value query parameter
- Sebenarnya URL melakukan parsing query parameter dan menyimpannya dalam map[string][]string
- Artinya, dalam satu key query parameter, kita bisa memasukkan beberapa value
- Caranya kita bisa menambahkan query parameter dengan nama yang sama, namun value berbeda, misal :
- name=Eko&name=Kurniawan

## kode multiple value query parameter
```go
func MultipleParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprintf(w, "%s", strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Thomas", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
```