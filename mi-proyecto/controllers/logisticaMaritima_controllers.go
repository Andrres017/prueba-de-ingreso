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

func GetAllMaritimas(writer http.ResponseWriter, request *http.Request) {
	maritimas := []models.LogisticaMaritima{}
	db := common.GetConnection()
	defer db.Close()

	db.Find(&maritimas)
	json, _ := json.Marshal(maritimas)
	common.SendResponse(writer, http.StatusOK, json)
}

func GetMaritima(write http.ResponseWriter, request *http.Request) {
	maritima := models.LogisticaMaritima{}
	id := mux.Vars(request)["id"]
	db := common.GetConnection()
	defer db.Close()

	db.Find(&maritima, id)

	if maritima.ID != nil {
		json, _ := json.Marshal(maritima)
		common.SendResponse(write, http.StatusOK, json)
	} else {
		common.SendError(write, http.StatusNotFound)
	}
}
func GetFilterMaritima(write http.ResponseWriter, request *http.Request) {
	campo := mux.Vars(request)["campo"]
	value := mux.Vars(request)["valor"]
	maritima := []models.LogisticaMaritima{}
	db := common.GetConnection()

	defer db.Close()
	db.Where(campo+" LIKE ?", "%"+value+"%").Find(&maritima)
	json, _ := json.Marshal(maritima)
	common.SendResponse(write, http.StatusOK, json)
}
func SaveMaritima(write http.ResponseWriter, request *http.Request) {

	maritima := new(models.LogisticaMaritima)
	db := common.GetConnection()

	defer db.Close()

	maritima = request.Context().Value("maritima").(*models.LogisticaMaritima)
	if maritima.ID == nil {
		if maritima.CantidadProducto > 10 {
			maritima.PrecioEnvio = utils.Descuento(maritima.PrecioEnvio, 3)
		}
		maritima.NumeroGuia = utils.GenerateCode(10)
	}

	err := db.Save(maritima).Error

	if err != nil {
		log.Fatal(err)
		common.SendError(write, http.StatusInternalServerError) //500
		return
	}

	json, _ := json.Marshal(maritima)
	common.SendResponse(write, http.StatusCreated, json)
}

func DeleteMaritima(write http.ResponseWriter, request *http.Request) {
	maritima := models.LogisticaMaritima{}
	db := common.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&maritima, id)

	if maritima.ID != nil {
		db.Delete(maritima)
		common.SendResponse(write, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(write, http.StatusNotFound)
	}

}
