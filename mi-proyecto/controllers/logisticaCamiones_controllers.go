package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tu-usuario/mi-proyecto/common"
	"github.com/tu-usuario/mi-proyecto/models"
	"github.com/tu-usuario/mi-proyecto/utils"
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

	if camion.ID != nil {
		json, _ := json.Marshal(camion)
		common.SendResponse(write, http.StatusOK, json)
	} else {
		common.SendError(write, http.StatusNotFound)
	}
}
func GetFilterCamiones(write http.ResponseWriter, request *http.Request) {
	campo := mux.Vars(request)["campo"]
	value := mux.Vars(request)["valor"]
	camion := []models.LogisticaCamione{}
	db := common.GetConnection()

	defer db.Close()
	db.Where(campo+" LIKE ?", "%"+value+"%").Find(&camion)
	json, _ := json.Marshal(camion)
	common.SendResponse(write, http.StatusOK, json)
}
func SaveCamion(write http.ResponseWriter, request *http.Request) {

	camion := new(models.LogisticaCamione)
	db := common.GetConnection()
	defer db.Close()
	camion = request.Context().Value("camion").(*models.LogisticaCamione)
	if camion.ID == nil {
		if camion.CantidadProducto > 10 {
			camion.PrecioEnvio = utils.Descuento(camion.PrecioEnvio, 5)
		}
		camion.NumeroGuia = utils.GenerateCode(10)
	}

	error := db.Save(&camion).Error

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

	if camion.ID != nil {
		db.Delete(camion)
		common.SendResponse(write, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(write, http.StatusNotFound)
	}
}
