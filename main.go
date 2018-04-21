package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Halaman Beranda")
}

func kontak(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Halaman Kontak")
}

func main() {
	http.HandleFunc("/",home)
	http.HandleFunc("/",kontak)
	http.ListenAndServe(":8282", nil)
}
