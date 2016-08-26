package main

import (
    "io"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func to_roman(n int)  string { 
    result := ""
    for n >= 1000 {
	n = n-1000
	result += "M"
    }
    for n >= 900 {
	n = n-900
	result += "CM"
    }
    for n >= 500 {
	n = n-500
	result += "D"
    }
    for n >= 400 {
	n = n-400
	result += "CD"
    }
    for n >= 100 {
	n = n-100
	result += "C"
    }
    for n >= 90 {
	n = n-90
	result += "XC"
    }
    for n >= 50 {
	n = n-50
	result += "L"
    }
    for n >= 40 {
	n = n-40
	result += "XL"
    }
    for n >= 10 {
	n = n-10
	result += "X"
    }
    for n >= 9 {
	n = n-9
	result += "IX"
    }
    for n >= 5 {
	n = n-5
	result += "V"
    }
    for n >= 4 {
	n = n-4
	result += "IV"
    }
    for n >= 1 {
	n = n-1
	result += "I"
    }   
    return result
}

type romanGenerator int
func (n romanGenerator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ascii_num := r.URL.Path[7:]
    i, err := strconv.Atoi(ascii_num)
    if err != nil {
        log.Print(err)
    }
    fmt.Fprintf(w, "Here's your number: %s\n", to_roman(i))
}



func main() {
    h := http.NewServeMux()

    h.Handle("/roman/", romanGenerator(1))
    h.HandleFunc("/", hello)

    err := http.ListenAndServe(":8000", h)
    log.Fatal(err)
}
