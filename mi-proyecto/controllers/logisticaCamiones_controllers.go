package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tu-usuario/mi-proyecto/common"
	"github.com/tu-usuario/mi-proyecto/models"
)

func GetAllCamiones(writer http.ResponseWriter, request *http.Request) {
	camiones := []models.LogisticaCamione{}
	db := common.GetConnection()
	defer db.Close()

	db.Find(&camiones)
	json, _ := json.Marshal(camiones)
	common.SendResponse(writer, http.StatusOK, json)
}

func GetCamion(write http.ResponseWriter, request *http.Request) {
	camion := models.LogisticaCamione{}
	id := mux.Vars(request)["id"]
	db := common.GetConnection()
	defer db.Close()

	db.Find(&camion, id)

	if &camion.ID != nil {
		json, _ := json.Marshal(camion)
		common.SendResponse(write, http.StatusOK, json)
	} else {
		common.SendError(write, http.StatusNotFound)
	}
}

func SaveCamion(write http.ResponseWriter, request *http.Request) {

	camion := models.LogisticaCamione{}
	db := common.GetConnection()
	defer db.Close()
	error := json.NewDecoder(request.Body).Decode(&camion)

	if error != nil {
		log.Fatal(error)
		common.SendError(write, http.StatusBadRequest)
		return
	}

	error = db.Save(&camion).Error

	if error != nil {
		log.Fatal(error)
		common.SendError(write, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(camion)
	common.SendResponse(write, http.StatusCreated, json)
}

func DeleteCamion(write http.ResponseWriter, request *http.Request) {
	camion := models.LogisticaCamione{}
	db := common.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&camion, id)

	if &camion.ID != nil {
		db.Delete(camion)
		common.SendResponse(write, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(write, http.StatusNotFound)
	}
}
