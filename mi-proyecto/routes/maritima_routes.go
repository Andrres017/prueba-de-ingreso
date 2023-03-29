package routes

import (
	"github.com/gorilla/mux"
	"github.com/tu-usuario/mi-proyecto/controllers"
	"github.com/tu-usuario/mi-proyecto/models"
	"github.com/tu-usuario/mi-proyecto/validations"
)

func SetMaritima(router *mux.Router) {
	subRoute := router.PathPrefix("/maritima/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAllMaritimas).Methods("GET")
	subRoute.HandleFunc("/find/{id}", controllers.GetMaritima).Methods("GET")
	subRoute.HandleFunc("/save", validations.ValidateUniversal(&models.LogisticaMaritima{}, "maritima")(controllers.SaveMaritima)).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.DeleteMaritima).Methods("DELETE")
}
