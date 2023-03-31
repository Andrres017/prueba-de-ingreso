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
	//writer.Header().Set("Access-Control-Allow-Origin", "*")
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
func GetFilterPersona(write http.ResponseWriter, request *http.Request) {
	campo := mux.Vars(request)["campo"]
	value := mux.Vars(request)["valor"]
	personas := []models.Persona{}
	db := common.GetConnection()

	defer db.Close()
	db.Where(campo+" LIKE ?", "%"+value+"%").Find(&personas)
	json, _ := json.Marshal(personas)
	common.SendResponse(write, http.StatusOK, json)
}
func Save(write http.ResponseWriter, request *http.Request) {
	persona := new(models.Persona)
	db := common.GetConnection()
	defer db.Close()
	persona = request.Context().Value("persona").(*models.Persona)
	if persona.ID == nil {
		sha1 := utils.Sha1Hex(persona.Clave)
		persona.Clave = sha1
	}

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
	//write.Header().Set("Access-Control-Allow-Origin", "*")
	persona := models.Persona{}
	db := common.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&persona, id)

	if persona.ID != nil {
		db.Delete(persona)
		common.SendResponse(write, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(write, http.StatusBadRequest)
	}

}

func Getlogin(write http.ResponseWriter, request *http.Request) {

	personas := models.Persona{}
	db := common.GetConnection()
	err := json.NewDecoder(request.Body).Decode(&personas)
	if err != nil {
		return
	}
	defer db.Close()
	claveOLD := personas.Clave
	db.Where("correo = ?", personas.Correo).First(&personas)
	fmt.Printf("%v \n", personas)

	var response []byte
	if personas.Clave == utils.Sha1Hex(claveOLD) {
		token, err := utils.GenerateToken(personas)
		if err != nil {
			fmt.Println("GG")
		}
		data := map[string]interface{}{
			"token":       token,
			"status_code": http.StatusOK,
		}
		response, _ = json.Marshal(data)
		common.SendResponse(write, http.StatusOK, response)
		fmt.Print(token)
	} else {
		fmt.Println("FF")
	}
}
