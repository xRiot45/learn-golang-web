# Video 3 | Golang Web

## Golang Web
- Go-Lang saat ini populer dijadikan salah satu pilihan bahasa pemrograman untuk membuat Web, terutama Web API (Backend)
- Selain itu, di Go-Lang juga sudah disediakan package untuk membuat Web, bahkan di sertakan pula package untuk implementasi unit testing untuk Web
- Hal ini menjadikan pembuatan Web menggunakan Go-Lang lebih mudah, karena tidak butuh menggunakan library atau framework

## Cara Kerja Golang Web
- Web Browser akan melakukan HTTP Request ke Web Server
- Golang menerima HTTP Request, lalu mengeksekusi request tersebut sesuai dengan yang diminta.
- Setelah melakukan eksekusi request, Golang akan mengembalikan data dan di render sesuai dengan kebutuhannya, misal HTML, CSS, JavaScript dan lain-lain
- Golang akan mengembalikan content hasil render tersebut tersebut sebagai HTTP Response ke Web Browser
- Web Browser menerima content dari Web Server, lalu me-render content tersebut sesuai dengan tipe content nya 

## Package net/http
- Pada beberapa bahasa pemrograman lain, seperti Java misalnya, untuk membuat web biasanya dibutuhkan tambahan library atau framework
- Sedangkan di Go-Lang sudah disediakan package untuk membuat web bernama package net/http
- Sehingga untuk membuat web menggunakan Go-Lang, kita tidak butuh lagi library tambahan, kita bisa menggunakan package yang sudah tersedia
- Walaupun memang saat kita akan membuat web dalam skala besar, direkomendasikan menggunakan framework karena beberapa hal sudah dipermudah oleh web framework
- Namun pada course ini, kita akan fokus menggunakan package net/http untuk membuat web nya, karena semua framework web Go-Lang akan menggunakan net/http sebagai basis dasar framework nya

