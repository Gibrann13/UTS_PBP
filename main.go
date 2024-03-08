package main

import (
	"Uts_Gibran/controllers"
	"fmt"
	"log"
	"net/http"

	//"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")

	router.HandleFunc("/rooms/{id}", controllers.GetRoomDetail).Methods("GET")

	router.HandleFunc("/rooms", controllers.InsertRoom).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 888")
	log.Fatal(http.ListenAndServe(":8888", router))

}
