package controllers

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/orutrax/go-crud-postgres-api/commons"
	"github.com/orutrax/go-crud-postgres-api/models"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	var movements []models.Movement

	db := commons.GetConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.Find(&movements)
	data, _ := json.Marshal(movements)
	commons.SendResponse(w, http.StatusOK, data)
}

func Get(w http.ResponseWriter, r *http.Request) {
	movement := models.Movement{}
	vars := mux.Vars(r)
	MovementUUID, err := uuid.Parse(vars["uuid"])

	if err != nil {
		commons.SendError(w, http.StatusBadRequest)
	}

	db := commons.GetConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	err = db.First(&movement, MovementUUID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		commons.SendError(w, http.StatusNotFound)
	} else {
		data, _ := json.Marshal(movement)
		commons.SendResponse(w, http.StatusOK, data)
	}
}

func Save(w http.ResponseWriter, r *http.Request) {
	movement := models.Movement{}

	db := commons.GetConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	err := json.NewDecoder(r.Body).Decode(&movement)

	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusBadRequest)
	}

	movement.UUID = uuid.New()
	err = db.Create(&movement).Error

	if err != nil {
		log.Fatal(err)
		commons.SendError(w, http.StatusInternalServerError)
	}

	data, _ := json.Marshal(movement)

	commons.SendResponse(w, http.StatusCreated, data)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	movement := models.Movement{}

	db := commons.GetConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	vars := mux.Vars(r)
	MovementUUID, err := uuid.Parse(vars["uuid"])

	if err != nil {
		commons.SendError(w, http.StatusBadRequest)
	}

	err = db.Delete(&movement, MovementUUID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		commons.SendError(w, http.StatusNotFound)
	} else {
		commons.SendResponse(w, http.StatusOK, []byte(`{}`))
	}
}
