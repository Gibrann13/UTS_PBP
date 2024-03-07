package main

import (
	"Modul_2/controllers"
	"fmt"
	"log"
	"net/http"

	//"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
)

func main() {
	//loadEnv()
	router := mux.NewRouter()

	//Get all users
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	//Get all product
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	//Get all transactions
	router.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")

	//Insert new users
	router.HandleFunc("/insertuser", func(w http.ResponseWriter, r *http.Request) {
		controllers.InsertUser(w, r, "Munsic", 19, "Mayann")
	}).Methods("POST")
	//Insert new product
	router.HandleFunc("/insertproduct", func(w http.ResponseWriter, r *http.Request) {
		controllers.InsertProduct(w, r, "Lem", 44)
	}).Methods("POST")
	//insert new transaction
	router.HandleFunc("/inserttransaction", func(w http.ResponseWriter, r *http.Request) {
		controllers.InsertTransaction(w, r, 1, 1, 5)
	}).Methods("POST")

	//delete user
	router.HandleFunc("/deleteuser", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUser(w, r, "Kinan", 5)
	}).Methods("DELETE")
	//delete product
	router.HandleFunc("/deleteproduct", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteProduct(w, r, 1)
	}).Methods("DELETE")
	//delete product
	router.HandleFunc("/deletetransaction", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteTransaction(w, r, 1, 1)
	}).Methods("DELETE")

	//update user
	router.HandleFunc("/updateuser", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateUser(w, r, "nanki", "Sedanggg")
	}).Methods("PUT")
	//update transaction
	router.HandleFunc("/updatetransaction", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateTransaction(w, r, 1, 1, 2, 10)
	}).Methods("PUT")
	//update product
	router.HandleFunc("/updateproduct", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateProduct(w, r, 2, "LilinApi", 3)
	}).Methods("PUT")

	router.HandleFunc("/insertuser/v2", controllers.InsertUserGORM).Methods("POST")
	//updateuser2
	router.HandleFunc("updateuser/v2", controllers.UpdateUserGORM).Methods("PUT")
	//deleteuser2
	router.HandleFunc("deleteuser/v2", controllers.DeleteUserGORM).Methods("DELETE")
	//selectuser2
	router.HandleFunc("selectuser2/v2", controllers.SelectUserGORM).Methods("GET")
	//raw/complex query
	router.HandleFunc("updateraw/v2", controllers.UpdateRawGORM).Methods("PUT")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
