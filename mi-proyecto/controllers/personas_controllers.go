package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tu-usuario/mi-proyecto/common"
	"github.com/tu-usuario/mi-proyecto/models"
	"github.com/tu-usuario/mi-proyecto/utils"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := common.GetConnection()

	defer db.Close()

	db.Find(&personas)
	json, _ := json.Marshal(personas)
	common.SendResponse(writer, http.StatusOK, json)
}

func Get(write http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	id := mux.Vars(request)["id"]
	db := common.GetConnection()
	defer db.Close()

	db.Find(&persona, id)

	if persona.ID != nil {
		json, _ := json.Marshal(persona)
		common.SendResponse(write, http.StatusOK, json)
	} else {
		common.SendError(write, http.StatusNotFound)
	}
}

func Save(write http.ResponseWriter, request *http.Request) {
	persona := new(models.Persona)
	db := common.GetConnection()
	defer db.Close()
	persona = request.Context().Value("persona").(*models.Persona)
	sha1 := utils.Sha1Hex(persona.Clave)
	fmt.Print(sha1)
	persona.Clave = sha1
	err := db.Save(persona).Error

	if err != nil {
		log.Fatal(err)
		common.SendError(write, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)
	common.SendResponse(write, http.StatusCreated, json)
}

func Delete(write http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	db := common.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&persona, id)

	if persona.ID != nil {
		db.Delete(persona)
		common.SendResponse(write, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(write, http.StatusNotFound)
	}

}
