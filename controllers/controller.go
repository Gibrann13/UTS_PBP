package controllers

import (
	"Uts_Gibran/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}
	var rooms []m.Room
	db := connectgorm()

	result := db.Find(&rooms)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.RoomResponse{}
	response.Status = http.StatusOK
	response.Message = "Success Get the users"
	json.NewEncoder(w).Encode(rooms)
}

func GetRoomDetail(w http.ResponseWriter, r *http.Request) {
	db := GetDB()
	if db == nil {
		return nil, ErrDBConnection
	}

	query := "SELECT id, room_name FROM rooms WHERE id = ?"
	row := db.QueryRow(query, roomID)

	var roomDetail RoomDetail
	if err := row.Scan(&roomDetail.ID, &roomDetail.RoomName); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrRoomNotFound
		}
		log.Println(err)
		return nil, err
	}

	// Jika includeParticipants true, Anda dapat menambahkan logika pengambilan participants di sini

	return &roomDetail, nil
}
