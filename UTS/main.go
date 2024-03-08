package main

import (
	"UTS/controllers"
	"fmt"
	"log"
	"net/http"

	//GORM

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()

	// Untuk Get Semua Rooms
	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")

	// Menggunakan ID untuk Detail Rooms
	router.HandleFunc("/rooms/{id}", controllers.GetRoomDetail).Methods("GET")

	// Mengguakan ID untuk mendapatkan detail room dan participantsnya
	router.HandleFunc("/rooms/{id}/participants", controllers.GetRoomDetailWithParticipants).Methods("GET")

	// Untuk Insert Participants ke Room
	router.HandleFunc("/rooms", controllers.InsertRoom).Methods("POST")

	// Keluarkan participants dari room
	router.HandleFunc("/rooms/{room_id}/participants/{participant_id}", controllers.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
