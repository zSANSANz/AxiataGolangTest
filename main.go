package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Product holds your product attribute
type Product struct {
	KodeProduk string `json:"kodeProduk"`
	Kuantitas  int    `json:"kuantitas"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat datang di home page")
}

func allProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Products)
}

func singleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	kodeProduk := params["id"]

	for _, product := range Products {
		if product.KodeProduk == kodeProduk {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func createProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Products = append(Products, product)
	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, p := range Products {
		if p.KodeProduk == id {
			Products[i].KodeProduk = product.KodeProduk
			Products[i].Kuantitas = product.Kuantitas
			json.NewEncoder(w).Encode(Products[i])
			return
		}
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	kodeProduk := params["id"]

	for i, p := range Products {
		if p.KodeProduk == kodeProduk {
			Products = append(Products[:i], Products[i+1:]...)
			json.NewEncoder(w).Encode(p)
			return
		}
	}
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", home)
	r.HandleFunc("/products", allProducts).Methods("GET")
	r.HandleFunc("/products/{id}", singleProduct).Methods("GET")
	r.HandleFunc("/products", createProducts).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	// pesan kalau aplikasi berjalan
	fmt.Println("Application running")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	Products = []Product{
		Product{
			KodeProduk: "apel",
			Kuantitas:  2,
		},
		Product{
			KodeProduk: "melon",
			Kuantitas:  5,
		},
		Product{
			KodeProduk: "semangka",
			Kuantitas:  7,
		},
	}
	handleRequest()
}

/*
Products is a global variable to hold collection of products
atau bisa diartikan variable global product berfungsi sebagai pengganti database
*/

/*
		[4:40 PM] Pandu Dwi Putra Nugroho
	    Data Structure
		Buat sebuah Software Library yang harus memilik fungsi:
		void tambahProduk(string kodeProduk, int kuantitas)
		void hapusProduk(string kodeProduk, int kuantitas)
		void hapusSeluruhProduk()
		void tampilkanSeluruhProduk()
		void tampilkanProduk(string kodeProduk)

Edited
*/
var Products []Product
