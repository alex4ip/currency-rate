module exchangerate

go 1.22

//replace github.com/ecommerce-challenge/currency-rate/go/exchangerate => ./go/exchangerate

require (
	github.com/gorilla/mux v1.8.1
	github.com/mattn/go-sqlite3 v1.14.22
)
