package controllers

import (
	"UTS/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := models.GetAllRooms()
	if err != nil {
		http.Error(w, "Gagal mendapatkan data rooms", http.StatusInternalServerError)
		return
	}

	response := models.RoomsResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    rooms,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetRoomDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID Room yang dimasukkan salah", http.StatusBadRequest)
		return
	}

	includeParticipants := r.URL.Query().Get("include_participants") == "true"
	room, err := models.GetRoomDetail(roomID, includeParticipants)
	if err != nil {
		http.Error(w, "Gagal mendapatkan data detail room", http.StatusInternalServerError)
		return
	}

	response := models.RoomDetailResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    room,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, "Gagal", http.StatusBadRequest)
		return
	}

	numParticipants, err := models.CountParticipantsByRoomID(room.GameID)
	if err != nil {
		http.Error(w, "Gagal menghitung participants dalam room", http.StatusInternalServerError)
		return
	}

	maxPlayers, err := models.GetMaxPlayersByGameID(room.GameID)
	if err != nil {
		http.Error(w, "Gagal mendapatkan info max pemain", http.StatusInternalServerError)
		return
	}

	if numParticipants >= maxPlayers {
		http.Error(w, "Max pemain sudah tercapai", http.StatusBadRequest)
		return
	}

	err = models.InsertRoom(&room)
	if err != nil {
		http.Error(w, "Gagal untuk masuk ke room", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := models.RoomResponse{
		Status:  http.StatusCreated,
		Message: "Room berhasil dimasukkan",
		Data:    room,
	}
	json.NewEncoder(w).Encode(response)
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	participantID, err := strconv.Atoi(params["participant_id"])
	if err != nil {
		http.Error(w, "Participants ID Salah", http.StatusBadRequest)
		return
	}
	roomID, err := strconv.Atoi(params["room_id"])
	if err != nil {
		http.Error(w, "Room ID Salah", http.StatusBadRequest)
		return
	}

	err = models.LeaveRoom(participantID, roomID)
	if err != nil {
		http.Error(w, "Gagal Keluar Ruangan", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := models.ErrorResponse{
		Status:  http.StatusOK,
		Message: "Berhasil Meninggalkan Ruangan",
	}
	json.NewEncoder(w).Encode(response)
}
