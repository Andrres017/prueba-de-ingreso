package routes

import (
	"github.com/gorilla/mux"
	"github.com/tu-usuario/mi-proyecto/controllers"
	"github.com/tu-usuario/mi-proyecto/models"
	"github.com/tu-usuario/mi-proyecto/validations"
)

func SetCamion(router *mux.Router) {
	subRoute := router.PathPrefix("/camion/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAllCamiones).Methods("GET")
	subRoute.HandleFunc("/find/{id}", controllers.GetAllCamiones).Methods("GET")
	subRoute.HandleFunc("/save", validations.ValidateUniversal(&models.LogisticaCamione{}, "camion")(controllers.SaveCamion)).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteCamion).Methods("DELETE")
}
